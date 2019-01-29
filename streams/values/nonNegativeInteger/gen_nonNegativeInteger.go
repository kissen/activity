package nonnegativeinteger

import "fmt"

// SerializeNonNegativeInteger converts a nonNegativeInteger value to an interface
// representation suitable for marshalling into a text or binary format.
func SerializeNonNegativeInteger(this int) (interface{}, error) {
	return this, nil
}

// DeserializeNonNegativeInteger creates nonNegativeInteger value from an
// interface representation that has been unmarshalled from a text or binary
// format.
func DeserializeNonNegativeInteger(this interface{}) (int, error) {
	if i, ok := this.(float64); ok {
		n := int(i)
		if n >= 0 {
			return n, nil
		} else {
			return 0, fmt.Errorf("%v is a negative integer for xsd:nonNegativeInteger", this)
		}
	} else {
		return 0, fmt.Errorf("%v cannot be interpreted as a float for xsd:nonNegativeInteger", this)
	}
}

// LessNonNegativeInteger returns true if the left nonNegativeInteger value is
// less than the right value.
func LessNonNegativeInteger(lhs, rhs int) bool {
	return lhs < rhs
}
