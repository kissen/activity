package boolean

import "fmt"

// SerializeBoolean converts a boolean value to an interface representation
// suitable for marshalling into a text or binary format.
func SerializeBoolean(this bool) (interface{}, error) {
	return this, nil
}

// DeserializeBoolean creates boolean value from an interface representation that
// has been unmarshalled from a text or binary format.
func DeserializeBoolean(this interface{}) (bool, error) {
	if b, ok := this.(bool); ok {
		return b, nil
	} else if f, ok := this.(float64); ok {
		if f == 0 {
			return false, nil
		} else if f == 1 {
			return true, nil
		} else {
			return false, fmt.Errorf("%v cannot be interpreted as a bool float64 for xsd:boolean", this)
		}
	} else {
		return false, fmt.Errorf("%v cannot be interpreted as a bool for xsd:boolean", this)
	}
}

// LessBoolean returns true if the left boolean value is less than the right value.
func LessBoolean(lhs, rhs bool) bool {
	// Booleans don't have a natural ordering, so we pick that truth is greater than falsehood.
	return !lhs && rhs
}
