package errors

import "strings"

// EntityError is a class of errors describing errors in entities.
//
// An entity here is defined as anything that has identifier associated to it.
// For example, a web page is an entity with its URL is the identifier.
type EntityError interface {
	Unwrappable
	EntityIdentifier() string
}

// Ent creates an instance of error that conforms EntityError. It takes
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

// EntFields is used to create an error that describes multiple field errors
// of an entity.
func EntFields(entityIdentifier string, fields ...EntityError) EntityError {
	return &entityError{
		identifier: entityIdentifier,
		fields:     fields[:],
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

// EntInvalid creates an EntityError with err set to DataErrInvalid.
func EntInvalid(entityIdentifier string, details error) EntityError {
	return &entityError{
		identifier: entityIdentifier,
		err: descriptorDetailsError{
			descriptor: DataErrInvalid,
			details:    details,
		},
	}
}

func IsEntInvalid(err error) bool {
	if err == DataErrInvalid {
		return true
	}
	if d, ok := err.(hasDescriptor); ok {
		desc := d.Descriptor()
		if desc == DataErrInvalid {
			return true
		}
	}
	return false
}

// EntErrNotFound is used to describet that the entity with the identifier
// could not be found in the system.
const EntErrNotFound = dataErrorConstantDescriptor("not found")

func EntNotFound(entityIdentifier string, details error) EntityError {
	return &entityError{
		identifier: entityIdentifier,
		err: descriptorDetailsError{
			descriptor: EntErrNotFound,
			details:    details,
		},
	}
}

func IsEntNotFound(err error) bool {
	if err == EntErrNotFound {
		return true
	}
	if d, ok := err.(hasDescriptor); ok {
		desc := d.Descriptor()
		if desc == EntErrNotFound {
			return true
		}
	}
	return false
}

type entityError struct {
	identifier string
	err        error
	fields     []EntityError
}

var (
	_ error         = &entityError{}
	_ Unwrappable   = &entityError{}
	_ EntityError   = &entityError{}
	_ hasDescriptor = &entityError{}
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
func (e *entityError) Descriptor() ErrorDescriptor {
	if e == nil {
		return nil
	}
	if desc, ok := e.err.(ErrorDescriptor); ok {
		return desc
	}
	if d, ok := e.err.(hasDescriptor); ok {
		return d.Descriptor()
	}
	if wrap, ok := e.err.(Unwrappable); ok {
		inner := wrap.Unwrap()
		if inner != nil {
			if desc, ok := inner.(ErrorDescriptor); ok {
				return desc
			}
		}
	}
	return nil
}
