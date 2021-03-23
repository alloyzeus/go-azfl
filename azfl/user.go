package azcore

// UserID abstracts the identifiers of User entity instances.
type UserIDNum interface {
	IDNum

	AZUserIDNum()
}

// UserRefKey is used to refer to a User entity instance.
type UserRefKey interface {
	RefKey

	// UserIDNum returns only the ID part of this ref-key.
	UserIDNum() UserIDNum
}
