package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

type UserIDNumMethods interface {
	AZUserIDNum()
}

// UserID abstracts the identifiers of User entity instances.
type UserIDNum interface {
	azid.IDNum

	UserIDNumMethods
}

// UserRefKey is used to refer to a User entity instance.
type UserRefKey[IDNumT UserIDNum] interface {
	azid.RefKey[IDNumT]

	// UserIDNum returns only the ID part of this ref-key.
	UserIDNum() IDNumT
}
