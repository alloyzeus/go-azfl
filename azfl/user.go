package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// UserID abstracts the identifiers of User entity instances.
type UserIDNum interface {
	azid.IDNum

	AZUserIDNum()
}

// UserRefKey is used to refer to a User entity instance.
type UserRefKey interface {
	azid.RefKey

	// UserIDNum returns only the ID part of this ref-key.
	UserIDNum() UserIDNum
}
