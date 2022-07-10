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

type EntityErrorBuilder interface {
	EntityError

	// Desc returns a copy with descriptor is set to desc.
	Desc(desc EntityErrorDescriptor) EntityErrorBuilder

	// DescMsg sets the descriptor with the provided string. For the best
	// experience, descMsg should be defined as a constant so that the error
	// users could use it to identify an error. For non-constant descriptor
	// use the Wrap method.
	DescMsg(descMsg string) EntityErrorBuilder

	// Wrap returns a copy with wrapped error is set to detailingError.
	Wrap(detailingError error) EntityErrorBuilder

	Fieldset(fields ...EntityError) EntityErrorBuilder
}

//TODO: custom type
type EntityErrorDescriptor = ErrorDescriptor

func IsEntityError(err error) bool {
	_, ok := err.(EntityError)
	return ok
}

// Ent creates an instance of error that conforms EntityError. It takes
// entityIdentifier which could be the name, key or URL of an entity. The
// entityIdentifier should describe the 'what' while err describes the 'why'.
//
//   // Describes that the file ".config.yaml" does not exist.
//   errors.Ent("./config.yaml").Wrap(os.ErrNotExist)
//
//   // Describes that the site "https://example.com" is unreachable.
//   errors.Ent("https://example.com/").DescMsg("unreachable")
//
func Ent(entityIdentifier string) EntityErrorBuilder {
	return &entityError{
		identifier: entityIdentifier,
	}
}

const (
	// ErrEntityNotFound is used when the entity with the specified name
	// cannot be found in the system.
	ErrEntityNotFound = valueConstantErrorDescriptor("not found")

	// ErrEntityConflict is used when the process of creating a new entity
	// fails because an entity with the same name already exists.
	ErrEntityConflict = valueConstantErrorDescriptor("conflict")
)

func IsEntityNotFoundError(err error) bool {
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
	descriptor ErrorDescriptor
	wrapped    error
	fields     []EntityError
}

var (
	_ error              = &entityError{}
	_ Unwrappable        = &entityError{}
	_ EntityError        = &entityError{}
	_ EntityErrorBuilder = &entityError{}
	_ hasDescriptor      = &entityError{}
)

func (e *entityError) Error() string {
	suffix := e.fieldErrorsAsString()
	if suffix != "" {
		suffix = ": " + suffix
	}
	var descStr string
	if e.descriptor != nil {
		descStr = e.descriptor.Error()
	}
	causeStr := errorString(e.wrapped)
	if causeStr == "" {
		causeStr = descStr
	} else if descStr != "" {
		causeStr = descStr + ": " + causeStr
	}

	if e.identifier != "" {
		if causeStr != "" {
			return e.identifier + ": " + causeStr + suffix
		}
		return e.identifier + suffix
	}
	if causeStr != "" {
		return "entity " + causeStr + suffix
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

func (e entityError) Unwrap() error            { return e.wrapped }
func (e entityError) EntityIdentifier() string { return e.identifier }
func (e entityError) Descriptor() ErrorDescriptor {
	if e.descriptor != nil {
		return e.descriptor
	}
	if desc, ok := e.wrapped.(ErrorDescriptor); ok {
		return desc
	}
	return nil
}

func (e entityError) Desc(desc EntityErrorDescriptor) EntityErrorBuilder {
	e.descriptor = desc
	return &e
}

func (e entityError) DescMsg(descMsg string) EntityErrorBuilder {
	e.descriptor = constantErrorDescriptor(descMsg)
	return &e
}

func (e entityError) Fieldset(fields ...EntityError) EntityErrorBuilder {
	e.fields = fields // copy?
	return &e
}

func (e entityError) Wrap(detailingError error) EntityErrorBuilder {
	e.wrapped = detailingError
	return &e
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
	return copyFieldSet(e)
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
