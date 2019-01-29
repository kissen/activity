package vocab

import (
	"net/url"
	"time"
)

// When the object describes a time-bound resource, such as an audio or video, a
// meeting, etc, the duration property indicates the object's approximate
// duration. The value MUST be expressed as an xsd:duration as defined by [
// xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as
// "PT5S").
//
// Example 119 (https://www.w3.org/TR/activitystreams-vocabulary/#ex134-jsonld):
//   {
//     "duration": "PT2H",
//     "name": "Birds Flying",
//     "type": "Video",
//     "url": "http://example.org/video.mkv"
//   }
type DurationPropertyInterface interface {
	// Clear ensures no value of this property is set. Calling IsDuration
	// afterwards will return false.
	Clear()
	// Get returns the value of this property. When IsDuration returns false,
	// Get will return any arbitrary value.
	Get() time.Duration
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsDuration returns true if this property is set and not an IRI.
	IsDuration() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o DurationPropertyInterface) bool
	// Name returns the name of this property: "duration".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets the value of this property. Calling IsDuration afterwards will
	// return true.
	Set(v time.Duration)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}
