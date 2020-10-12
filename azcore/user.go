package azcore

// UserRefKey represents a reference to a User.
type UserRefKey interface {
	RefKey
	AZUserRefKey()
}
