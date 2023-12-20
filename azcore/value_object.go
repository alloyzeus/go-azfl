package azcore

import "github.com/alloyzeus/go-azfl/v2/azob"

// A ValueObject is a small object that represents a simple entity whose
// equality is not based on identity, i.e., two value objects are equal when
// they have the same value, not necessarily being the same object.
type ValueObject interface {
	//TODO: we should have Clone method contract here, but it requires
	// covariant support which won't happen
	// https://github.com/golang/go/issues/30602
}

// ValueObjectAssert is tool to assert that a struct conforms the
// characteristic of a value-object.
//
// To use:
//
//	var _ = ValueObjectAssert[MyStruct] = MyStruct{}
type ValueObjectAssert[T any] interface {
	azob.CloneableAssert[T]
}
