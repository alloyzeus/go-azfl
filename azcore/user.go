package azcore

// UserID abstracts the identifiers of User entity instances.
type UserID interface {
	EID
	AZUserID()
}

// UserRefKey is used to refer to a User entity instance.
type UserRefKey interface {
	RefKey

	// UserID returns only the ID part of this ref-key.
	UserID() UserID
}
