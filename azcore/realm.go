package azcore

import "net/url"

// RealmService provides information about the realm.
//
// A realm is a site or the whole service.
type RealmService interface {
	AZRealmService()
}

// RealmServiceServer provides contracts for a realm service.
type RealmServiceServer interface {
	RealmService

	RealmName() string

	WebsiteURL() url.URL
}
