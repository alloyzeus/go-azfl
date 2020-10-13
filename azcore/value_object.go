package azcore

// A ValueObject is a small object that represents a simple entity whose
// equality is not based on identity, i.e., two value objects are equal when
// they have the same value, not necessarily being the same object.
type ValueObject interface {
	Equatable
}
