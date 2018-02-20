package pub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	// ErrObjectRequired means the activity needs its object property set.
	ErrObjectRequired = errors.New("object property required")
	// ErrTypeRequired means the activity needs its type property set.
	ErrTypeRequired = errors.New("type property required")
)

// TODO: Helper http Handler for serving ActivityStream objects
// TODO: Helper http Handler for serving Tombstone objects
// TODO: Helper http Handler for serving deleted objects

// TODO: Helper http Handler for serving actor's likes
// TODO: Helper http Handler for serving actor's followers
// TODO: Helper http Handler for serving actor's following

// TODO: Authorization client-to-server.
// TODO: Authenticate server-to-server deliveries.

type federator struct {
	// Clock determines the time of this federator.
	Clock Clock
	// App is the client application that is ActivityPub aware.
	//
	// It is always required.
	App Application
	// FederateApp provides utility when handling incoming messages received
	// via the Federated Protocol, or server-to-server communications.
	//
	// It is only required if EnableServer is true.
	FederateApp FederateApp
	// SocialApp provides utility when handling incoming messages
	// received via the Social API, or client-to-server communications.
	//
	// It is only required if EnableClient is true.
	SocialApp SocialApp

	ClientCallbacker Callbacker
	ServerCallbacker Callbacker

	// Client is used to federate with other ActivityPub servers.
	//
	// It is only required if EnableServer is true.
	Client *http.Client
	// Agent is the User-Agent string to use in HTTP headers when
	// federating with another server. It will automatically be appended
	// with '(go-fed ActivityPub)'.
	//
	// It is only required if EnableServer is true.
	Agent string
	// MaxDeliveryDepth is how deep collections of recipients will be
	// expanded for delivery. It must be at least 1 to be compliant with the
	// ActivityPub spec.
	//
	// It is only required if EnableServer is true.
	MaxDeliveryDepth int
	// EnableClient enables or disables the Social API, the client to
	// server part of ActivityPub. Useful if permitting remote clients to
	// act on behalf of the users of the client application.
	EnableClient bool
	// EnableServer enables or disables the Federated Protocol, or the
	// server to server part of ActivityPub. Useful to permit integrating
	// with the rest of the federative web.
	EnableServer bool
	// pool properly handles rate-limiting and retrying deliveries to other
	// federated servers.
	//
	// It is only required if EnableServer is true.
	pool *delivererPool
}

func (f *federator) Stop() {
	if f.pool != nil {
		f.pool.Stop()
	}
}

func (f *federator) Errors() <-chan error {
	if f.pool != nil {
		return f.pool.Errors()
	}
	return nil
}

func (f *federator) PostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubPost(r) {
		return false, nil
	}
	if !f.EnableServer {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return true, nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return true, err
	}
	var m map[string]interface{}
	if err = json.Unmarshal(b, &m); err != nil {
		return true, err
	}
	ao, err := getActorObject(m)
	if err != nil {
		return true, err
	}
	var iris []url.URL
	for i := 0; i < ao.ActorLen(); i++ {
		if ao.IsActorObject(i) {
			obj := ao.GetActorObject(i)
			if obj.HasId() {
				iris = append(iris, obj.GetId())
			}
		} else if ao.IsActorLink(i) {
			l := ao.GetActorLink(i)
			if l.HasHref() {
				iris = append(iris, l.GetHref())
			}
		} else if ao.IsActorIRI(i) {
			iris = append(iris, ao.GetActorIRI(i))
		}
	}
	if err = f.FederateApp.Unblocked(c, iris); err != nil {
		return true, err
	}
	if err = f.getPostInboxResolver(c).Deserialize(m); err != nil {
		return true, err
	}
	// TODO: 7.1.2 Inbox forwarding
	w.WriteHeader(http.StatusOK)
	return true, nil
}

func (f *federator) GetInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubGet(r) {
		return false, nil
	}
	oc, err := f.App.GetInbox(c, r)
	if err != nil {
		return true, err
	}
	oc, err = f.dedupeOrderedItems(oc)
	if err != nil {
		return true, err
	}
	m, err := oc.Serialize()
	if err != nil {
		return true, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return true, err
	}
	w.Header().Set(contentTypeHeader, responseContentTypeHeader)
	w.WriteHeader(http.StatusOK)
	n, err := w.Write(b)
	if err != nil {
		return true, err
	} else if n != len(b) {
		return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(b))
	}
	return true, nil
}

// PostOutbox provides a HTTP handler for ActivityPub requests for the given id
// token. The client ID token is passed forwards to other interfaces for
// application specific behavior. The handler will return true if it handled
// the request as an ActivityPub request. If it returns an error, it is up to
// the client to determine how to respond via HTTP.
//
// Note that the error could be ErrObjectRequired or ErrTypeRequired.
func (f *federator) PostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubPost(r) {
		return false, nil
	}
	if !f.EnableClient {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return true, nil
	}
	ok, err := f.App.PostOutboxAuthorized(c, r)
	if err != nil {
		return true, err
	}
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return true, nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return true, err
	}
	var m map[string]interface{}
	if err = json.Unmarshal(b, &m); err != nil {
		return true, err
	}
	typer, err := toTypeIder(m)
	if err != nil {
		return true, err
	}
	if !isActivityType(typer) {
		actorIri, err := f.App.ActorIRI(c, r)
		if err != nil {
			return true, err
		}
		obj, ok := typer.(vocab.ObjectType)
		if !ok {
			return true, fmt.Errorf("wrap in create: cannot convert to vocab.ObjectType: %T", typer)
		}
		typer = f.wrapInCreate(obj, actorIri)
	}
	newId := f.App.NewId(c, typer)
	typer.SetId(newId)
	if m, err = typer.Serialize(); err != nil {
		return true, err
	}
	if err := f.addToOutbox(c, r, m); err != nil {
		return true, err
	}
	deliverable := false
	if err = f.getPostOutboxResolver(c, &deliverable).Deserialize(m); err != nil {
		return true, err
	}
	if f.EnableServer && deliverable {
		obj, err := toAnyActivity(m)
		if err != nil {
			return true, err
		}
		if err := f.deliver(obj); err != nil {
			return true, err
		}
	}
	w.Header().Set("Location", newId.String())
	w.WriteHeader(http.StatusCreated)
	return true, nil
}

func (f *federator) GetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubGet(r) {
		return false, nil
	}
	oc, err := f.App.GetOutbox(c, r)
	if err != nil {
		return true, err
	}
	m, err := oc.Serialize()
	if err != nil {
		return true, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return true, err
	}
	w.Header().Set(contentTypeHeader, responseContentTypeHeader)
	w.WriteHeader(http.StatusOK)
	n, err := w.Write(b)
	if err != nil {
		return true, err
	} else if n != len(b) {
		return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(b))
	}
	return true, nil
}

func (f *federator) addToOutbox(c context.Context, r *http.Request, m map[string]interface{}) error {
	outbox, err := f.App.GetOutbox(c, r)
	if err != nil {
		return err
	}
	activity, err := toAnyActivity(m)
	if err != nil {
		return err
	}
	outbox.AddOrderedItemsObject(activity)
	return f.App.Set(c, outbox)
}

func (f *federator) getPostOutboxResolver(c context.Context, deliverable *bool) *streams.Resolver {
	return &streams.Resolver{
		CreateCallback: f.handleClientCreate(c, deliverable),
		UpdateCallback: f.handleClientUpdate(c, deliverable),
		DeleteCallback: f.handleClientDelete(c, deliverable),
		FollowCallback: f.handleClientFollow(c, deliverable),
		AcceptCallback: f.handleClientAccept(c, deliverable),
		RejectCallback: f.handleClientReject(c, deliverable),
		AddCallback:    f.handleClientAdd(c, deliverable),
		RemoveCallback: f.handleClientRemove(c, deliverable),
		LikeCallback:   f.handleClientLike(c, deliverable),
		UndoCallback:   f.handleClientUndo(c, deliverable),
		BlockCallback:  f.handleClientBlock(c, deliverable),
		// TODO: Extended activity types, such as Announce, Arrive, etc.
	}
}

func (f *federator) handleClientCreate(ctx context.Context, deliverable *bool) func(s *streams.Create) error {
	return func(s *streams.Create) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		c := s.Raw()
		// When a Create activity is posted, the actor of the activity
		// SHOULD be copied onto the object's attributedTo field.
		// Presumably only if it doesn't already exist, to prevent
		// duplicate deliveries.
		createActorIds := make(map[string]interface{})
		for i := 0; i < c.ActorLen(); i++ {
			if c.IsActorObject(i) {
				obj := c.GetActorObject(i)
				id := obj.GetId()
				createActorIds[(&id).String()] = obj
			} else if c.IsActorLink(i) {
				l := c.GetActorLink(i)
				href := l.GetHref()
				createActorIds[(&href).String()] = l
			} else if c.IsActorIRI(i) {
				iri := c.GetActorIRI(i)
				createActorIds[(&iri).String()] = iri
			}
		}
		var obj []vocab.ObjectType
		for i := 0; i < c.ObjectLen(); i++ {
			if c.IsObject(i) {
				obj = append(obj, c.GetObject(i))
			} else if c.IsObjectIRI(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("unsupported: Create Activity with 'object' that is only an IRI")
			}
		}
		objectAttributedToIds := make([]map[string]interface{}, len(obj))
		for i := range objectAttributedToIds {
			objectAttributedToIds[i] = make(map[string]interface{})
		}
		for k, o := range obj {
			for i := 0; i < o.AttributedToLen(); i++ {
				if o.IsAttributedToObject(i) {
					at := o.GetAttributedToObject(i)
					id := o.GetId()
					objectAttributedToIds[k][(&id).String()] = at
				} else if o.IsAttributedToLink(i) {
					at := o.GetAttributedToLink(i)
					href := at.GetHref()
					objectAttributedToIds[k][(&href).String()] = at
				} else if o.IsAttributedToIRI(i) {
					iri := o.GetAttributedToIRI(i)
					objectAttributedToIds[k][(&iri).String()] = iri
				}
			}
		}
		for k, v := range createActorIds {
			for i, attributedToMap := range objectAttributedToIds {
				if _, ok := attributedToMap[k]; !ok {
					if vObj, ok := v.(vocab.ObjectType); ok {
						obj[i].AddAttributedToObject(vObj)
					} else if vLink, ok := v.(vocab.LinkType); ok {
						obj[i].AddAttributedToLink(vLink)
					} else if vIRI, ok := v.(url.URL); ok {
						obj[i].AddAttributedToIRI(vIRI)
					}
				}
			}
		}
		// As such, a server SHOULD copy any recipients of the Create activity to its
		// object upon initial distribution, and likewise with copying recipients from
		// the object to the wrapping Create activity.
		// Again, presumably if it does not already exist.
		for _, attributedToMap := range objectAttributedToIds {
			for k, v := range attributedToMap {
				if _, ok := createActorIds[k]; !ok {
					if vObj, ok := v.(vocab.ObjectType); ok {
						c.AddActorObject(vObj)
					} else if vLink, ok := v.(vocab.LinkType); ok {
						c.AddActorLink(vLink)
					} else if vIRI, ok := v.(url.URL); ok {
						c.AddActorIRI(vIRI)
					}
				}
			}
		}
		// Create requires the client application to persist the 'object' that
		// was created.
		for _, o := range obj {
			if err := f.App.Set(ctx, o); err != nil {
				return err
			}
		}
		return f.ClientCallbacker.Create(ctx, s)
	}
}

func (f *federator) handleClientUpdate(c context.Context, deliverable *bool) func(s *streams.Update) error {
	return func(s *streams.Update) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		// TODO: Support redactions via the '"json": null' value.
		// Update should partially replace the 'object' with only the
		// changed top-level fields.
		ids, err := getObjectIds(s.Raw())
		if err != nil {
			return err
		} else if len(ids) == 0 {
			return fmt.Errorf("update has no id: %v", s)
		}
		for idx, id := range ids {
			pObj, err := f.App.Get(c, id)
			if err != nil {
				return err
			}
			obj, ok := pObj.(vocab.Serializer)
			if !ok {
				return fmt.Errorf("PubObject is not vocab.Serializer: %T", pObj)
			}
			m, err := obj.Serialize()
			if err != nil {
				return err
			}
			if !s.Raw().IsObject(idx) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("update requires object to be wholly provided at index %d", idx)
			}
			updated, err := s.Raw().GetObject(idx).Serialize()
			if err != nil {
				return err
			}
			for k, v := range updated {
				m[k] = v
			}
			p, err := ToPubObject(m)
			if err != nil {
				return err
			}
			for _, elem := range p {
				if err := f.App.Set(c, elem); err != nil {
					return err
				}
			}
		}
		return f.ClientCallbacker.Update(c, s)
	}
}

func (f *federator) handleClientDelete(c context.Context, deliverable *bool) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		ids, err := getObjectIds(s.Raw())
		if err != nil {
			return err
		} else if len(ids) == 0 {
			return fmt.Errorf("delete has no id: %v", s)
		}
		for _, id := range ids {
			pObj, err := f.App.Get(c, id)
			if err != nil {
				return err
			}
			obj, ok := pObj.(vocab.ObjectType)
			if !ok {
				return fmt.Errorf("cannot delete non-ObjectType: %T", pObj)
			}
			tomb := toTombstone(obj, id, f.Clock.Now())
			if err := f.App.Set(c, tomb); err != nil {
				return err
			}
		}
		return f.ClientCallbacker.Delete(c, s)
	}
}

func (f *federator) handleClientFollow(c context.Context, deliverable *bool) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		// Nothing extra to do.
		return f.ClientCallbacker.Follow(c, s)
	}
}

func (f *federator) handleClientAccept(c context.Context, deliverable *bool) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		*deliverable = true
		return f.ClientCallbacker.Accept(c, s)
	}
}

func (f *federator) handleClientReject(c context.Context, deliverable *bool) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		*deliverable = true
		return f.ClientCallbacker.Reject(c, s)
	}
}

func (f *federator) handleClientAdd(c context.Context, deliverable *bool) func(s *streams.Add) error {
	return func(s *streams.Add) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		} else if s.LenType() == 0 {
			return ErrTypeRequired
		}
		raw := s.Raw()
		ids, err := getTargetIds(raw)
		if err != nil {
			return err
		} else if len(ids) == 0 {
			return fmt.Errorf("add target has no ids: %v", s)
		}
		objIds, err := getObjectIds(s.Raw())
		if err != nil {
			return err
		} else if len(objIds) == 0 {
			return fmt.Errorf("add object has no ids: %v", s)
		}
		var targets []vocab.ObjectType
		for _, id := range ids {
			if !f.SocialApp.Owns(c, id) {
				continue
			}
			target, err := f.App.Get(c, id)
			if err != nil {
				return err
			}
			ct, okCollection := target.(vocab.CollectionType)
			oct, okOrdered := target.(vocab.OrderedCollectionType)
			if !okCollection && !okOrdered {
				return fmt.Errorf("cannot add to type that is not Collection and not OrderedCollection: %v", target)
			} else if okCollection {
				targets = append(targets, ct)
			} else {
				targets = append(targets, oct)
			}
		}
		for i := 0; i < raw.ObjectLen(); i++ {
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("add object must be object type: %v", raw)
			}
			obj := raw.GetObject(i)
			if !f.SocialApp.Owns(c, obj.GetId()) {
				continue
			}
			for _, target := range targets {
				if !f.SocialApp.CanAdd(c, obj, target) {
					continue
				}
				if ct, ok := target.(vocab.CollectionType); ok {
					ct.AddItemsObject(obj)
				} else if oct, ok := target.(vocab.OrderedCollectionType); ok {
					oct.AddOrderedItemsObject(obj)
				}
				if err := f.App.Set(c, target); err != nil {
					return err
				}
			}
		}
		return f.ClientCallbacker.Add(c, s)
	}
}

func (f *federator) handleClientRemove(c context.Context, deliverable *bool) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		} else if s.LenType() == 0 {
			return ErrTypeRequired
		}
		raw := s.Raw()
		ids, err := getTargetIds(raw)
		if err != nil {
			return err
		} else if len(ids) == 0 {
			return fmt.Errorf("remove target has no ids: %v", s)
		}
		objIds, err := getObjectIds(s.Raw())
		if err != nil {
			return err
		} else if len(objIds) == 0 {
			return fmt.Errorf("remove object has no ids: %v", s)
		}
		var targets []vocab.ObjectType
		for _, id := range ids {
			if !f.SocialApp.Owns(c, id) {
				continue
			}
			target, err := f.App.Get(c, id)
			if err != nil {
				return err
			}
			ct, okCollection := target.(vocab.CollectionType)
			oct, okOrdered := target.(vocab.OrderedCollectionType)
			if !okCollection && !okOrdered {
				return fmt.Errorf("cannot remove from type that is not Collection and not OrderedCollection: %v", target)
			} else if okCollection {
				targets = append(targets, ct)
			} else {
				targets = append(targets, oct)
			}
		}
		for i := 0; i < raw.ObjectLen(); i++ {
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("remove object must be object type: %v", raw)
			}
			obj := raw.GetObject(i)
			if !f.SocialApp.Owns(c, obj.GetId()) {
				continue
			}
			for _, target := range targets {
				if !f.SocialApp.CanRemove(c, obj, target) {
					continue
				}
				if ct, ok := target.(vocab.CollectionType); ok {
					removeCollectionItemWithId(ct, obj.GetId())
				} else if oct, ok := target.(vocab.OrderedCollectionType); ok {
					removeOrderedCollectionItemWithId(oct, obj.GetId())
				}
				if err := f.App.Set(c, target); err != nil {
					return err
				}
			}
		}
		return f.ClientCallbacker.Remove(c, s)
	}
}

func (f *federator) handleClientLike(ctx context.Context, deliverable *bool) func(s *streams.Like) error {
	return func(s *streams.Like) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		if err := f.addToAllActorLikedCollection(ctx, s.Raw()); err != nil {
			return err
		}
		return f.ClientCallbacker.Like(ctx, s)
	}
}

func (f *federator) handleClientUndo(c context.Context, deliverable *bool) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		// TODO: Determine if we can support common forms of undo natively.
		return f.ClientCallbacker.Undo(c, s)
	}
}

func (f *federator) handleClientBlock(c context.Context, deliverable *bool) func(s *streams.Block) error {
	return func(s *streams.Block) error {
		*deliverable = false
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		return f.ClientCallbacker.Block(c, s)
	}
}

func (f *federator) getPostInboxResolver(c context.Context) *streams.Resolver {
	return &streams.Resolver{
		CreateCallback: f.handleCreate(c),
		UpdateCallback: f.handleUpdate(c),
		DeleteCallback: f.handleDelete(c),
		FollowCallback: f.handleFollow(c),
		AcceptCallback: f.handleAccept(c),
		RejectCallback: f.handleReject(c),
		AddCallback:    f.handleAdd(c),
		RemoveCallback: f.handleRemove(c),
		LikeCallback:   f.handleLike(c),
		UndoCallback:   f.handleUndo(c),
		// TODO: Extended activity types, such as Announce, Arrive, etc.
	}
}

func (f *federator) handleCreate(c context.Context) func(s *streams.Create) error {
	return func(s *streams.Create) error {
		// Create requires the client application to persist the 'object' that
		// was created.
		raw := s.Raw()
		for i := 0; i < raw.ObjectLen(); i++ {
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("create requires object to be wholly provided at index %d", i)
			}
			obj := raw.GetObject(i)
			if err := f.App.Set(c, obj); err != nil {
				return err
			}
		}
		return f.ServerCallbacker.Create(c, s)
	}
}

func (f *federator) handleUpdate(c context.Context) func(s *streams.Update) error {
	return func(s *streams.Update) error {
		// TODO: The receiving server MUST take care to be sure that the Update
		// is authorized to modify its object.
		raw := s.Raw()
		for i := 0; i < raw.ObjectLen(); i++ {
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("update requires object to be wholly provided at index %d", i)
			}
			obj := raw.GetObject(i)
			if err := f.App.Set(c, obj); err != nil {
				return err
			}
		}
		return f.ServerCallbacker.Update(c, s)
	}
}

func (f *federator) handleDelete(c context.Context) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
		// TODO: Verify ownership. I think the spec unintentionally suggests to
		// just assume it is owned, so we will actually verify.
		ids, err := getObjectIds(s.Raw())
		if err != nil {
			return err
		} else if len(ids) == 0 {
			return fmt.Errorf("delete has no id: %v", s)
		}
		for _, id := range ids {
			pObj, err := f.App.Get(c, id)
			if err != nil {
				return err
			}
			obj, ok := pObj.(vocab.ObjectType)
			if !ok {
				return fmt.Errorf("cannot delete non-ObjectType: %T", pObj)
			}
			tomb := toTombstone(obj, id, f.Clock.Now())
			if err := f.App.Set(c, tomb); err != nil {
				return err
			}
		}
		return f.ServerCallbacker.Delete(c, s)
	}
}

func (f *federator) handleFollow(c context.Context) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		// Permit either human-triggered or automatically triggering
		// 'Accept'/'Reject'.
		todo := f.FederateApp.OnFollow(c, s)
		if todo != DoNothing {
			var activity vocab.ActivityType
			if todo == AutomaticAccept {
				activity = &vocab.Accept{}
			} else if todo == AutomaticReject {
				activity = &vocab.Reject{}
			}
			raw := s.Raw()
			activity.AddObject(raw)
			for i := 0; i < raw.ActorLen(); i++ {
				if raw.IsActorObject(i) {
					activity.AddToObject(raw.GetActorObject(i))
				} else if raw.IsActorLink(i) {
					activity.AddToLink(raw.GetActorLink(i))
				} else if raw.IsActorIRI(i) {
					activity.AddToIRI(raw.GetActorIRI(i))
				}
			}
			if err := f.deliver(activity); err != nil {
				return err
			}
			if todo == AutomaticAccept {
				// TODO: Add to followers collection.
			}
		}
		return f.ServerCallbacker.Follow(c, s)
	}
}

func (f *federator) handleAccept(c context.Context) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		// Accept can be client application specific. However, if this 'Accept'
		// is in response to a 'Follow' then the 'actor' should be added to the
		// original 'actor's 'following' collection by the client application.
		raw := s.Raw()
		for i := 0; i < raw.ObjectLen(); i++ {
			if raw.IsObject(i) {
				obj := raw.GetObject(i)
				follow, ok := obj.(vocab.FollowType)
				if !ok {
					continue
				}
				for j := 0; j < follow.ActorLen(); j++ {
					var iri url.URL
					if follow.IsActorObject(j) {
						actor := follow.GetActorObject(j)
						if !actor.HasId() {
							return fmt.Errorf("actor object on follow must have id")
						}
						iri = actor.GetId()
					} else if follow.IsActorLink(j) {
						l := follow.GetActorLink(j)
						if !l.HasHref() {
							return fmt.Errorf("actor link on follow must have href")
						}
						iri = l.GetHref()
					} else if follow.IsActorIRI(j) {
						iri = follow.GetActorIRI(j)
					}
					following, err := f.FederateApp.GetFollowing(c, iri)
					if err != nil {
						return err
					}
					// TODO: Deduplication detection.
					for k := 0; k < raw.ActorLen(); k++ {
						if raw.IsActorObject(k) {
							following.AddItemsObject(raw.GetActorObject(k))
						} else if raw.IsActorLink(k) {
							following.AddItemsLink(raw.GetActorLink(k))
						} else if raw.IsActorIRI(k) {
							following.AddItemsIRI(raw.GetActorIRI(k))
						}
					}
					f.App.Set(c, following)
				}
			}
		}
		return f.ServerCallbacker.Accept(c, s)
	}
}

func (f *federator) handleReject(c context.Context) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		// Reject can be client application specific. However, if this 'Reject'
		// is in response to a 'Follow' then the client MUST NOT go forward with
		// adding the 'actor' to the original 'actor's 'following' collection
		// by the client application.
		return f.ServerCallbacker.Reject(c, s)
	}
}

func (f *federator) handleAdd(c context.Context) func(s *streams.Add) error {
	return func(s *streams.Add) error {
		// TODO: Implement.
		// Add is client application specific, generally involving adding an
		// 'object' to a specific 'target' collection.
		return f.ServerCallbacker.Add(c, s)
	}
}

func (f *federator) handleRemove(c context.Context) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
		// TODO: Implement.
		// Remove is client application specific, generally involving removing
		// an 'object' from a specific 'target' collection.
		return f.ServerCallbacker.Remove(c, s)
	}
}

func (f *federator) handleLike(c context.Context) func(s *streams.Like) error {
	return func(s *streams.Like) error {
		if err := f.addToAllLikesCollections(c, s.Raw()); err != nil {
			return err
		}
		return f.ServerCallbacker.Like(c, s)
	}
}

func (f *federator) handleUndo(c context.Context) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
		// Undo negates a previous action. The 'actor' on the 'Undo' MUST be the
		// same as the 'actor' on the Activity being undone, and the client
		// application is responsible for enforcing this. Note that 'Undo'-ing
		// is not a deletion of a previous Activity, but the addition of its
		// opposite.
		return f.ServerCallbacker.Undo(c, s)
	}
}