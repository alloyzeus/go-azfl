package azcore

import (
	"golang.org/x/text/language"
)

// Context is a generalized context for all service methods.
type Context interface{}

// HumanContext is a specialized context, where current processing
// was initiated by a human.
type HumanContext interface {
	Context

	// AcceptLanguage is analogous to HTTP Accept-Language header. The
	// languages must be ordered by the human's preference.
	// If the languages comes as weighted, as found in HTTP Accept-Language,
	// sort the languages by their weights then drop the weights.
	AcceptLanguage() []language.Tag
}
