package errors

import "strings"

// EntityError is a class of errors describing errors in entities.
// An entity here is defined as anything which has identifier associated to it.
// For example, a web page is an entity which its URL is the identifier.
type EntityError interface {
	Unwrappable
	EntityIdentifier() string
}

// Ent creates an instance of error which conforms EntityError. It takes
// entityIdentifier which could be the name, key or URL of an entity. The
// entityIdentifier should describe the 'what' while err describes the 'why'.
//
//   // Describes that the file ".config.yaml" does not exist.
//   errors.Ent("./config.yaml", os.ErrNotExist)
//
//   // Describes that the site "https://example.com" is unreachable.
//   errors.Ent("https://example.com/", errors.Msg("unreachable"))
//
func Ent(entityIdentifier string, err error) EntityError {
	return &entityError{
		identifier: entityIdentifier,
		err:        err,
	}
}

// EntMsg creates an instance of error from an entitity identifier and the
// error message which describes why the entity is considered error.
func EntMsg(entityIdentifier string, errMsg string) EntityError {
	return &entityError{
		identifier: entityIdentifier,
		err:        Msg(errMsg),
	}
}

type entityError struct {
	identifier string
	err        error
	fields     []EntityError
}

var (
	_ error       = &entityError{}
	_ Unwrappable = &entityError{}
	_ EntityError = &entityError{}
)

func (e *entityError) Error() string {
	suffix := e.fieldErrorsAsString()
	if suffix != "" {
		suffix = ": " + suffix
	}

	if e.identifier != "" {
		if errMsg := e.innerMsg(); errMsg != "" {
			return e.identifier + ": " + errMsg + suffix
		}
		return e.identifier + " invalid" + suffix
	}
	if errMsg := e.innerMsg(); errMsg != "" {
		return "entity " + errMsg + suffix
	}
	return "entity invalid" + suffix
}

func (e *entityError) innerMsg() string {
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}

func (e entityError) fieldErrorsAsString() string {
	if flen := len(e.fields); flen > 0 {
		parts := make([]string, 0, flen)
		for _, sub := range e.fields {
			parts = append(parts, sub.Error())
		}
		return strings.Join(parts, ", ")
	}
	return ""
}

func (e *entityError) Unwrap() error            { return e.err }
func (e *entityError) EntityIdentifier() string { return e.identifier }
