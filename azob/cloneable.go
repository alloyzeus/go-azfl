package azob

// CloneableAssert is an interface to assert that type implements
// the method Clone that returns a copy of itself.
//
// This is interface is partial solution to the fact that Go doesn't
// support covariance.
type CloneableAssert[T any] interface {
	// Clone returns a deep copy of the instance. Some types, like mutex and
	// context, are not supposed to be copied.
	Clone() T
}
