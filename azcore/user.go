package azcore

import "github.com/alloyzeus/go-azfl/v2/azid"

type UserIDNumMethods interface {
}

// UserID abstracts the identifiers of User entity instances.
type UserIDNum interface {
	azid.IDNum

	UserIDNumMethods
}

// UserID is used to refer to a User entity instance.
type UserID[IDNumT UserIDNum] interface {
	azid.ID[IDNumT]

	// UserIDNum returns only the ID part of this ref-key.
	UserIDNum() IDNumT
}
