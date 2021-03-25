package azob

// Equatable defines the requirements for implementations so that they can
// be compared for value-based equality.
type Equatable interface {
	// Equals returns true if the other value matches this object perfectly.
	//
	// Note that value objects are defined by their attributes; even if the
	// types have different names, two value objects are equal if both
	// have the same attributes.
	Equals(interface{}) bool

	// Equal should be redirected to Equals. This method is required
	// to make all implementations conform the equality interface
	// needed by github.com/google/go-cmp .
	Equal(interface{}) bool
}
