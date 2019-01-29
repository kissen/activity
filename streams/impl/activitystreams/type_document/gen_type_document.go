package typedocument

import vocab "github.com/go-fed/activity/streams/vocab"

// Represents a document of any kind.
//
// Example 49 (https://www.w3.org/TR/activitystreams-vocabulary/#ex48-jsonld):
//   {
//     "name": "4Q Sales Forecast",
//     "type": "Document",
//     "url": "http://example.org/4q-sales-forecast.pdf"
//   }
type Document struct {
	Altitude     vocab.AltitudePropertyInterface
	Attachment   vocab.AttachmentPropertyInterface
	AttributedTo vocab.AttributedToPropertyInterface
	Audience     vocab.AudiencePropertyInterface
	Bcc          vocab.BccPropertyInterface
	Bto          vocab.BtoPropertyInterface
	Cc           vocab.CcPropertyInterface
	Content      vocab.ContentPropertyInterface
	Context      vocab.ContextPropertyInterface
	Duration     vocab.DurationPropertyInterface
	EndTime      vocab.EndTimePropertyInterface
	Generator    vocab.GeneratorPropertyInterface
	Icon         vocab.IconPropertyInterface
	Id           vocab.IdPropertyInterface
	Image        vocab.ImagePropertyInterface
	InReplyTo    vocab.InReplyToPropertyInterface
	Location     vocab.LocationPropertyInterface
	MediaType    vocab.MediaTypePropertyInterface
	Name         vocab.NamePropertyInterface
	Object       vocab.ObjectPropertyInterface
	Preview      vocab.PreviewPropertyInterface
	Published    vocab.PublishedPropertyInterface
	Replies      vocab.RepliesPropertyInterface
	StartTime    vocab.StartTimePropertyInterface
	Summary      vocab.SummaryPropertyInterface
	Tag          vocab.TagPropertyInterface
	To           vocab.ToPropertyInterface
	Type         vocab.TypePropertyInterface
	Updated      vocab.UpdatedPropertyInterface
	Url          vocab.UrlPropertyInterface
	alias        string
	unknown      map[string]interface{}
}

// DeserializeDocument creates a Document from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializeDocument(m map[string]interface{}, aliasMap map[string]string) (*Document, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	this := &Document{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAltitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Altitude = p
	}
	if p, err := mgr.DeserializeAttachmentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Attachment = p
	}
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.AttributedTo = p
	}
	if p, err := mgr.DeserializeAudiencePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Audience = p
	}
	if p, err := mgr.DeserializeBccPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Bcc = p
	}
	if p, err := mgr.DeserializeBtoPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Bto = p
	}
	if p, err := mgr.DeserializeCcPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Cc = p
	}
	if p, err := mgr.DeserializeContentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Content = p
	}
	if p, err := mgr.DeserializeContextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Context = p
	}
	if p, err := mgr.DeserializeDurationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Duration = p
	}
	if p, err := mgr.DeserializeEndTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.EndTime = p
	}
	if p, err := mgr.DeserializeGeneratorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Generator = p
	}
	if p, err := mgr.DeserializeIconPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Icon = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Id = p
	}
	if p, err := mgr.DeserializeImagePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Image = p
	}
	if p, err := mgr.DeserializeInReplyToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.InReplyTo = p
	}
	if p, err := mgr.DeserializeLocationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Location = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.MediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Name = p
	}
	if p, err := mgr.DeserializeObjectPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Object = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Preview = p
	}
	if p, err := mgr.DeserializePublishedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Published = p
	}
	if p, err := mgr.DeserializeRepliesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Replies = p
	}
	if p, err := mgr.DeserializeStartTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.StartTime = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Summary = p
	}
	if p, err := mgr.DeserializeTagPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Tag = p
	}
	if p, err := mgr.DeserializeToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.To = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Type = p
	}
	if p, err := mgr.DeserializeUpdatedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Updated = p
	}
	if p, err := mgr.DeserializeUrlPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Url = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "altitude" {
			continue
		} else if k == "attachment" {
			continue
		} else if k == "attributedTo" {
			continue
		} else if k == "audience" {
			continue
		} else if k == "bcc" {
			continue
		} else if k == "bto" {
			continue
		} else if k == "cc" {
			continue
		} else if k == "content" {
			continue
		} else if k == "context" {
			continue
		} else if k == "duration" {
			continue
		} else if k == "endTime" {
			continue
		} else if k == "generator" {
			continue
		} else if k == "icon" {
			continue
		} else if k == "id" {
			continue
		} else if k == "image" {
			continue
		} else if k == "inReplyTo" {
			continue
		} else if k == "location" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "object" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "published" {
			continue
		} else if k == "replies" {
			continue
		} else if k == "startTime" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "tag" {
			continue
		} else if k == "to" {
			continue
		} else if k == "type" {
			continue
		} else if k == "updated" {
			continue
		} else if k == "url" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// DocumentExtends returns true if the Document type extends from the other type.
func DocumentExtends(other vocab.Type) bool {
	extensions := []string{"Object"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// DocumentIsDisjointWith returns true if the other provided type is disjoint with
// the Document type.
func DocumentIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Link", "Mention"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// DocumentIsExtendedBy returns true if the other provided type extends from the
// Document type.
func DocumentIsExtendedBy(other vocab.Type) bool {
	extensions := []string{"Audio", "Page", "Image", "Video"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// NewDocument creates a new Document type
func NewDocument() *Document {
	return &Document{
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// GetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this Document) GetAltitude() vocab.AltitudePropertyInterface {
	return this.Altitude
}

// GetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this Document) GetAttachment() vocab.AttachmentPropertyInterface {
	return this.Attachment
}

// GetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Document) GetAttributedTo() vocab.AttributedToPropertyInterface {
	return this.AttributedTo
}

// GetAudience returns the "audience" property if it exists, and nil otherwise.
func (this Document) GetAudience() vocab.AudiencePropertyInterface {
	return this.Audience
}

// GetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this Document) GetBcc() vocab.BccPropertyInterface {
	return this.Bcc
}

// GetBto returns the "bto" property if it exists, and nil otherwise.
func (this Document) GetBto() vocab.BtoPropertyInterface {
	return this.Bto
}

// GetCc returns the "cc" property if it exists, and nil otherwise.
func (this Document) GetCc() vocab.CcPropertyInterface {
	return this.Cc
}

// GetContent returns the "content" property if it exists, and nil otherwise.
func (this Document) GetContent() vocab.ContentPropertyInterface {
	return this.Content
}

// GetContext returns the "context" property if it exists, and nil otherwise.
func (this Document) GetContext() vocab.ContextPropertyInterface {
	return this.Context
}

// GetDuration returns the "duration" property if it exists, and nil otherwise.
func (this Document) GetDuration() vocab.DurationPropertyInterface {
	return this.Duration
}

// GetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this Document) GetEndTime() vocab.EndTimePropertyInterface {
	return this.EndTime
}

// GetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this Document) GetGenerator() vocab.GeneratorPropertyInterface {
	return this.Generator
}

// GetIcon returns the "icon" property if it exists, and nil otherwise.
func (this Document) GetIcon() vocab.IconPropertyInterface {
	return this.Icon
}

// GetId returns the "id" property if it exists, and nil otherwise.
func (this Document) GetId() vocab.IdPropertyInterface {
	return this.Id
}

// GetImage returns the "image" property if it exists, and nil otherwise.
func (this Document) GetImage() vocab.ImagePropertyInterface {
	return this.Image
}

// GetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this Document) GetInReplyTo() vocab.InReplyToPropertyInterface {
	return this.InReplyTo
}

// GetLocation returns the "location" property if it exists, and nil otherwise.
func (this Document) GetLocation() vocab.LocationPropertyInterface {
	return this.Location
}

// GetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Document) GetMediaType() vocab.MediaTypePropertyInterface {
	return this.MediaType
}

// GetName returns the "name" property if it exists, and nil otherwise.
func (this Document) GetName() vocab.NamePropertyInterface {
	return this.Name
}

// GetObject returns the "object" property if it exists, and nil otherwise.
func (this Document) GetObject() vocab.ObjectPropertyInterface {
	return this.Object
}

// GetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Document) GetPreview() vocab.PreviewPropertyInterface {
	return this.Preview
}

// GetPublished returns the "published" property if it exists, and nil otherwise.
func (this Document) GetPublished() vocab.PublishedPropertyInterface {
	return this.Published
}

// GetReplies returns the "replies" property if it exists, and nil otherwise.
func (this Document) GetReplies() vocab.RepliesPropertyInterface {
	return this.Replies
}

// GetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this Document) GetStartTime() vocab.StartTimePropertyInterface {
	return this.StartTime
}

// GetSummary returns the "summary" property if it exists, and nil otherwise.
func (this Document) GetSummary() vocab.SummaryPropertyInterface {
	return this.Summary
}

// GetTag returns the "tag" property if it exists, and nil otherwise.
func (this Document) GetTag() vocab.TagPropertyInterface {
	return this.Tag
}

// GetTo returns the "to" property if it exists, and nil otherwise.
func (this Document) GetTo() vocab.ToPropertyInterface {
	return this.To
}

// GetType returns the "type" property if it exists, and nil otherwise.
func (this Document) GetType() vocab.TypePropertyInterface {
	return this.Type
}

// GetUnknownProperties returns the unknown properties for the Document type. Note
// that this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this Document) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this Document) GetUpdated() vocab.UpdatedPropertyInterface {
	return this.Updated
}

// GetUrl returns the "url" property if it exists, and nil otherwise.
func (this Document) GetUrl() vocab.UrlPropertyInterface {
	return this.Url
}

// IsExtending returns true if the Document type extends from the other type.
func (this Document) IsExtending(other vocab.Type) bool {
	return DocumentExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this Document) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.Altitude, m)
	m = this.helperJSONLDContext(this.Attachment, m)
	m = this.helperJSONLDContext(this.AttributedTo, m)
	m = this.helperJSONLDContext(this.Audience, m)
	m = this.helperJSONLDContext(this.Bcc, m)
	m = this.helperJSONLDContext(this.Bto, m)
	m = this.helperJSONLDContext(this.Cc, m)
	m = this.helperJSONLDContext(this.Content, m)
	m = this.helperJSONLDContext(this.Context, m)
	m = this.helperJSONLDContext(this.Duration, m)
	m = this.helperJSONLDContext(this.EndTime, m)
	m = this.helperJSONLDContext(this.Generator, m)
	m = this.helperJSONLDContext(this.Icon, m)
	m = this.helperJSONLDContext(this.Id, m)
	m = this.helperJSONLDContext(this.Image, m)
	m = this.helperJSONLDContext(this.InReplyTo, m)
	m = this.helperJSONLDContext(this.Location, m)
	m = this.helperJSONLDContext(this.MediaType, m)
	m = this.helperJSONLDContext(this.Name, m)
	m = this.helperJSONLDContext(this.Object, m)
	m = this.helperJSONLDContext(this.Preview, m)
	m = this.helperJSONLDContext(this.Published, m)
	m = this.helperJSONLDContext(this.Replies, m)
	m = this.helperJSONLDContext(this.StartTime, m)
	m = this.helperJSONLDContext(this.Summary, m)
	m = this.helperJSONLDContext(this.Tag, m)
	m = this.helperJSONLDContext(this.To, m)
	m = this.helperJSONLDContext(this.Type, m)
	m = this.helperJSONLDContext(this.Updated, m)
	m = this.helperJSONLDContext(this.Url, m)

	return m
}

// LessThan computes if this Document is lesser, with an arbitrary but stable
// determination.
func (this Document) LessThan(o vocab.DocumentInterface) bool {
	// Begin: Compare known properties
	// Compare property "altitude"
	if this.Altitude.LessThan(o.GetAltitude()) {
		return true
	} else if o.GetAltitude().LessThan(this.Altitude) {
		return false
	}
	// Compare property "attachment"
	if this.Attachment.LessThan(o.GetAttachment()) {
		return true
	} else if o.GetAttachment().LessThan(this.Attachment) {
		return false
	}
	// Compare property "attributedTo"
	if this.AttributedTo.LessThan(o.GetAttributedTo()) {
		return true
	} else if o.GetAttributedTo().LessThan(this.AttributedTo) {
		return false
	}
	// Compare property "audience"
	if this.Audience.LessThan(o.GetAudience()) {
		return true
	} else if o.GetAudience().LessThan(this.Audience) {
		return false
	}
	// Compare property "bcc"
	if this.Bcc.LessThan(o.GetBcc()) {
		return true
	} else if o.GetBcc().LessThan(this.Bcc) {
		return false
	}
	// Compare property "bto"
	if this.Bto.LessThan(o.GetBto()) {
		return true
	} else if o.GetBto().LessThan(this.Bto) {
		return false
	}
	// Compare property "cc"
	if this.Cc.LessThan(o.GetCc()) {
		return true
	} else if o.GetCc().LessThan(this.Cc) {
		return false
	}
	// Compare property "content"
	if this.Content.LessThan(o.GetContent()) {
		return true
	} else if o.GetContent().LessThan(this.Content) {
		return false
	}
	// Compare property "context"
	if this.Context.LessThan(o.GetContext()) {
		return true
	} else if o.GetContext().LessThan(this.Context) {
		return false
	}
	// Compare property "duration"
	if this.Duration.LessThan(o.GetDuration()) {
		return true
	} else if o.GetDuration().LessThan(this.Duration) {
		return false
	}
	// Compare property "endTime"
	if this.EndTime.LessThan(o.GetEndTime()) {
		return true
	} else if o.GetEndTime().LessThan(this.EndTime) {
		return false
	}
	// Compare property "generator"
	if this.Generator.LessThan(o.GetGenerator()) {
		return true
	} else if o.GetGenerator().LessThan(this.Generator) {
		return false
	}
	// Compare property "icon"
	if this.Icon.LessThan(o.GetIcon()) {
		return true
	} else if o.GetIcon().LessThan(this.Icon) {
		return false
	}
	// Compare property "id"
	if this.Id.LessThan(o.GetId()) {
		return true
	} else if o.GetId().LessThan(this.Id) {
		return false
	}
	// Compare property "image"
	if this.Image.LessThan(o.GetImage()) {
		return true
	} else if o.GetImage().LessThan(this.Image) {
		return false
	}
	// Compare property "inReplyTo"
	if this.InReplyTo.LessThan(o.GetInReplyTo()) {
		return true
	} else if o.GetInReplyTo().LessThan(this.InReplyTo) {
		return false
	}
	// Compare property "location"
	if this.Location.LessThan(o.GetLocation()) {
		return true
	} else if o.GetLocation().LessThan(this.Location) {
		return false
	}
	// Compare property "mediaType"
	if this.MediaType.LessThan(o.GetMediaType()) {
		return true
	} else if o.GetMediaType().LessThan(this.MediaType) {
		return false
	}
	// Compare property "name"
	if this.Name.LessThan(o.GetName()) {
		return true
	} else if o.GetName().LessThan(this.Name) {
		return false
	}
	// Compare property "object"
	if this.Object.LessThan(o.GetObject()) {
		return true
	} else if o.GetObject().LessThan(this.Object) {
		return false
	}
	// Compare property "preview"
	if this.Preview.LessThan(o.GetPreview()) {
		return true
	} else if o.GetPreview().LessThan(this.Preview) {
		return false
	}
	// Compare property "published"
	if this.Published.LessThan(o.GetPublished()) {
		return true
	} else if o.GetPublished().LessThan(this.Published) {
		return false
	}
	// Compare property "replies"
	if this.Replies.LessThan(o.GetReplies()) {
		return true
	} else if o.GetReplies().LessThan(this.Replies) {
		return false
	}
	// Compare property "startTime"
	if this.StartTime.LessThan(o.GetStartTime()) {
		return true
	} else if o.GetStartTime().LessThan(this.StartTime) {
		return false
	}
	// Compare property "summary"
	if this.Summary.LessThan(o.GetSummary()) {
		return true
	} else if o.GetSummary().LessThan(this.Summary) {
		return false
	}
	// Compare property "tag"
	if this.Tag.LessThan(o.GetTag()) {
		return true
	} else if o.GetTag().LessThan(this.Tag) {
		return false
	}
	// Compare property "to"
	if this.To.LessThan(o.GetTo()) {
		return true
	} else if o.GetTo().LessThan(this.To) {
		return false
	}
	// Compare property "type"
	if this.Type.LessThan(o.GetType()) {
		return true
	} else if o.GetType().LessThan(this.Type) {
		return false
	}
	// Compare property "updated"
	if this.Updated.LessThan(o.GetUpdated()) {
		return true
	} else if o.GetUpdated().LessThan(this.Updated) {
		return false
	}
	// Compare property "url"
	if this.Url.LessThan(o.GetUrl()) {
		return true
	} else if o.GetUrl().LessThan(this.Url) {
		return false
	}
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this Document) Serialize() (interface{}, error) {
	m := make(map[string]interface{})
	// Begin: Serialize known properties
	// Maybe serialize property "altitude"
	if i, err := this.Altitude.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Altitude.Name()] = i
	}
	// Maybe serialize property "attachment"
	if i, err := this.Attachment.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Attachment.Name()] = i
	}
	// Maybe serialize property "attributedTo"
	if i, err := this.AttributedTo.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.AttributedTo.Name()] = i
	}
	// Maybe serialize property "audience"
	if i, err := this.Audience.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Audience.Name()] = i
	}
	// Maybe serialize property "bcc"
	if i, err := this.Bcc.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Bcc.Name()] = i
	}
	// Maybe serialize property "bto"
	if i, err := this.Bto.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Bto.Name()] = i
	}
	// Maybe serialize property "cc"
	if i, err := this.Cc.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Cc.Name()] = i
	}
	// Maybe serialize property "content"
	if i, err := this.Content.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Content.Name()] = i
	}
	// Maybe serialize property "context"
	if i, err := this.Context.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Context.Name()] = i
	}
	// Maybe serialize property "duration"
	if i, err := this.Duration.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Duration.Name()] = i
	}
	// Maybe serialize property "endTime"
	if i, err := this.EndTime.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.EndTime.Name()] = i
	}
	// Maybe serialize property "generator"
	if i, err := this.Generator.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Generator.Name()] = i
	}
	// Maybe serialize property "icon"
	if i, err := this.Icon.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Icon.Name()] = i
	}
	// Maybe serialize property "id"
	if i, err := this.Id.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Id.Name()] = i
	}
	// Maybe serialize property "image"
	if i, err := this.Image.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Image.Name()] = i
	}
	// Maybe serialize property "inReplyTo"
	if i, err := this.InReplyTo.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.InReplyTo.Name()] = i
	}
	// Maybe serialize property "location"
	if i, err := this.Location.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Location.Name()] = i
	}
	// Maybe serialize property "mediaType"
	if i, err := this.MediaType.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.MediaType.Name()] = i
	}
	// Maybe serialize property "name"
	if i, err := this.Name.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Name.Name()] = i
	}
	// Maybe serialize property "object"
	if i, err := this.Object.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Object.Name()] = i
	}
	// Maybe serialize property "preview"
	if i, err := this.Preview.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Preview.Name()] = i
	}
	// Maybe serialize property "published"
	if i, err := this.Published.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Published.Name()] = i
	}
	// Maybe serialize property "replies"
	if i, err := this.Replies.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Replies.Name()] = i
	}
	// Maybe serialize property "startTime"
	if i, err := this.StartTime.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.StartTime.Name()] = i
	}
	// Maybe serialize property "summary"
	if i, err := this.Summary.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Summary.Name()] = i
	}
	// Maybe serialize property "tag"
	if i, err := this.Tag.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Tag.Name()] = i
	}
	// Maybe serialize property "to"
	if i, err := this.To.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.To.Name()] = i
	}
	// Maybe serialize property "type"
	if i, err := this.Type.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Type.Name()] = i
	}
	// Maybe serialize property "updated"
	if i, err := this.Updated.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Updated.Name()] = i
	}
	// Maybe serialize property "url"
	if i, err := this.Url.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Url.Name()] = i
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this Document) SetAltitude(i vocab.AltitudePropertyInterface) {
	this.Altitude = i
}

// SetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this Document) SetAttachment(i vocab.AttachmentPropertyInterface) {
	this.Attachment = i
}

// SetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Document) SetAttributedTo(i vocab.AttributedToPropertyInterface) {
	this.AttributedTo = i
}

// SetAudience returns the "audience" property if it exists, and nil otherwise.
func (this Document) SetAudience(i vocab.AudiencePropertyInterface) {
	this.Audience = i
}

// SetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this Document) SetBcc(i vocab.BccPropertyInterface) {
	this.Bcc = i
}

// SetBto returns the "bto" property if it exists, and nil otherwise.
func (this Document) SetBto(i vocab.BtoPropertyInterface) {
	this.Bto = i
}

// SetCc returns the "cc" property if it exists, and nil otherwise.
func (this Document) SetCc(i vocab.CcPropertyInterface) {
	this.Cc = i
}

// SetContent returns the "content" property if it exists, and nil otherwise.
func (this Document) SetContent(i vocab.ContentPropertyInterface) {
	this.Content = i
}

// SetContext returns the "context" property if it exists, and nil otherwise.
func (this Document) SetContext(i vocab.ContextPropertyInterface) {
	this.Context = i
}

// SetDuration returns the "duration" property if it exists, and nil otherwise.
func (this Document) SetDuration(i vocab.DurationPropertyInterface) {
	this.Duration = i
}

// SetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this Document) SetEndTime(i vocab.EndTimePropertyInterface) {
	this.EndTime = i
}

// SetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this Document) SetGenerator(i vocab.GeneratorPropertyInterface) {
	this.Generator = i
}

// SetIcon returns the "icon" property if it exists, and nil otherwise.
func (this Document) SetIcon(i vocab.IconPropertyInterface) {
	this.Icon = i
}

// SetId returns the "id" property if it exists, and nil otherwise.
func (this Document) SetId(i vocab.IdPropertyInterface) {
	this.Id = i
}

// SetImage returns the "image" property if it exists, and nil otherwise.
func (this Document) SetImage(i vocab.ImagePropertyInterface) {
	this.Image = i
}

// SetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this Document) SetInReplyTo(i vocab.InReplyToPropertyInterface) {
	this.InReplyTo = i
}

// SetLocation returns the "location" property if it exists, and nil otherwise.
func (this Document) SetLocation(i vocab.LocationPropertyInterface) {
	this.Location = i
}

// SetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Document) SetMediaType(i vocab.MediaTypePropertyInterface) {
	this.MediaType = i
}

// SetName returns the "name" property if it exists, and nil otherwise.
func (this Document) SetName(i vocab.NamePropertyInterface) {
	this.Name = i
}

// SetObject returns the "object" property if it exists, and nil otherwise.
func (this Document) SetObject(i vocab.ObjectPropertyInterface) {
	this.Object = i
}

// SetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Document) SetPreview(i vocab.PreviewPropertyInterface) {
	this.Preview = i
}

// SetPublished returns the "published" property if it exists, and nil otherwise.
func (this Document) SetPublished(i vocab.PublishedPropertyInterface) {
	this.Published = i
}

// SetReplies returns the "replies" property if it exists, and nil otherwise.
func (this Document) SetReplies(i vocab.RepliesPropertyInterface) {
	this.Replies = i
}

// SetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this Document) SetStartTime(i vocab.StartTimePropertyInterface) {
	this.StartTime = i
}

// SetSummary returns the "summary" property if it exists, and nil otherwise.
func (this Document) SetSummary(i vocab.SummaryPropertyInterface) {
	this.Summary = i
}

// SetTag returns the "tag" property if it exists, and nil otherwise.
func (this Document) SetTag(i vocab.TagPropertyInterface) {
	this.Tag = i
}

// SetTo returns the "to" property if it exists, and nil otherwise.
func (this Document) SetTo(i vocab.ToPropertyInterface) {
	this.To = i
}

// SetType returns the "type" property if it exists, and nil otherwise.
func (this Document) SetType(i vocab.TypePropertyInterface) {
	this.Type = i
}

// SetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this Document) SetUpdated(i vocab.UpdatedPropertyInterface) {
	this.Updated = i
}

// SetUrl returns the "url" property if it exists, and nil otherwise.
func (this Document) SetUrl(i vocab.UrlPropertyInterface) {
	this.Url = i
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this Document) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}
