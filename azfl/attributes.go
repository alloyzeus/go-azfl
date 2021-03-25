package azcore

import "github.com/alloyzeus/go-azfl/azfl/azob"

// Attributes abstracts attributes.
type Attributes interface {
	azob.Equatable

	AZAttributes()
}
