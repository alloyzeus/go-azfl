package azcore

// A ValueObject is a small object that represents a simple entity whose
// equality is not based on identity, i.e., two value objects are equal when
// they have the same value, not necessarily being the same object.
type ValueObject interface {
	//TODO: we should have Clone method contract here, but it requires
	// covariant support which won't happen
	// https://github.com/golang/go/issues/30602
}
