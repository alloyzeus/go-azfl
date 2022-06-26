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

func IsEntityError(err error) bool {
	_, ok := err.(EntityError)
	return ok
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
	if fields != nil {
		fields = fields[:]
	}
	return &entityError{
		identifier: entityIdentifier,
		fields:     fields,
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
		err:        DescWrap(ErrValueInvalid, details),
	}
}

func IsEntInvalidError(err error) bool {
	if !IsEntityError(err) {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueInvalid
	}
	return false
}

// ErrEntityNotFound is used to describet that the entity with the identifier
// could not be found in the system.
const ErrEntityNotFound = valueConstantErrorDescriptor("not found")

func EntNotFound(entityIdentifier string, details error) EntityError {
	return &entityError{
		identifier: entityIdentifier,
		err:        DescWrap(ErrEntityNotFound, details),
	}
}

func IsEntNotFoundError(err error) bool {
	if !IsEntityError(err) {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrEntityNotFound
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
	detailsStr := errorString(e.err)

	if e.identifier != "" {
		if detailsStr != "" {
			return e.identifier + ": " + detailsStr + suffix
		}
		return e.identifier + suffix
	}
	if detailsStr != "" {
		return "entity " + detailsStr + suffix
	}
	if suffix != "" {
		return "entity" + suffix
	}
	return "entity error"
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
	if desc := UnwrapDescriptor(e.err); desc != nil {
		return desc
	}
	return nil
}

// EntityErrorSet is an interface to combine multiple EntityError instances
// into a single error.
type EntityErrorSet interface {
	ErrorSet
	EntityErrors() []EntityError
}

// UnwrapEntityErrorSet returns the contained EntityError instances if err is
// indeed a EntityErrorSet.
func UnwrapEntityErrorSet(err error) []EntityError {
	if errSet := asEntityErrorSet(err); errSet != nil {
		return errSet.EntityErrors()
	}
	return nil
}

// asEntityErrorSet returns err as an EntityErrorSet if err is indeed an
// EntityErrorSet, otherwise it returns nil.
func asEntityErrorSet(err error) EntityErrorSet {
	e, _ := err.(EntityErrorSet)
	return e
}

// EntSet creates a compound error comprised of multiple instances of
// EntityError. Note that the resulting error is not an EntityError because
// it has no identity.
func EntSet(entityErrors ...EntityError) EntityErrorSet {
	ours := make([]EntityError, len(entityErrors))
	copy(ours, entityErrors)
	return entErrorSet(ours)
}

type entErrorSet []EntityError

var (
	_ error          = entErrorSet{}
	_ ErrorSet       = entErrorSet{}
	_ EntityErrorSet = entErrorSet{}
)

func (e entErrorSet) Error() string {
	if len(e) > 0 {
		errs := make([]string, 0, len(e))
		for _, ce := range e {
			errs = append(errs, ce.Error())
		}
		s := strings.Join(errs, ", ")
		if s != "" {
			return s
		}
	}
	return ""
}

func (e entErrorSet) Errors() []error {
	if len(e) > 0 {
		errs := make([]error, 0, len(e))
		for _, i := range e {
			errs = append(errs, i)
		}
		return errs
	}
	return nil
}

func (e entErrorSet) EntityErrors() []EntityError {
	if len(e) > 0 {
		errs := make([]EntityError, len(e))
		copy(errs, e)
		return errs
	}
	return nil
}

func copyFieldSet(fields []EntityError) []EntityError {
	var copiedFields []EntityError
	if len(fields) > 0 {
		copiedFields = make([]EntityError, 0, len(fields))
		for _, e := range fields {
			if e != nil {
				copiedFields = append(copiedFields, e)
			}
		}
	}
	return copiedFields
}
