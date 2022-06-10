package azcore

import (
	"context"
	"time"

	"github.com/alloyzeus/go-azfl/azob"
	"golang.org/x/text/language"
)

//region ServiceMethodCallContext

type ServiceMethodCallContext interface {
	ServiceMethodContext

	AZServiceMethodCallContext()

	// MethodName returns the name of the method or the endpoint.
	//
	// For HTTP, this method returns the method. For other protocols, it should
	// be the name of the method e.g., `getUser`.
	MethodName() string

	// ResourceID returns the identifier of the resource being accessed by
	// the call. For HTTP, it's the path. For other protocols, it should
	// be the identifier (ID) of the entity. If there's more than one
	// resources, e.g., a relationship between two entities, then it returns
	// the identifiers of the entities separated by commas.
	ResourceID() string
}

//endregion

type ServiceMethodCallInputMetadata interface {
	//TODO: request id / correlation id / idemptotency key, receive time by
	// the handler.
}

// ServiceMethodCallOriginInfo holds information about a call's origin.
//
//TODO: key-value custom data
type ServiceMethodCallOriginInfo struct {
	// Address returns the IP address or hostname where this call was initiated
	// from. This field might be empty if it's not possible to resolve
	// the address (e.g., the server is behind a proxy or a load-balancer and
	// they didn't forward the the origin IP).
	Address string

	// EnvironmentString returns some details of the environment,
	// might include application's version information, where the application
	// which made the request runs on. For web app, this method usually
	// returns the browser's user-agent string.
	EnvironmentString string

	// AcceptLanguage is analogous to HTTP Accept-Language header field. The
	// languages must be ordered by the human's preference.
	// If the languages comes as weighted, as found in HTTP Accept-Language,
	// sort the languages by their weights then drop the weights.
	AcceptLanguage []language.Tag

	// DateTime is the time of the caller device when the call was initiated.
	//
	// Analogous to HTTP Date header field.
	DateTime *time.Time
}

//region ServiceMethodIdempotencyKey

// ServiceMethodIdempotencyKey represents the identifier of a method call. This
// identifier doubles as idempotency token.
type ServiceMethodIdempotencyKey interface {
	azob.Equatable

	AZServiceMethodIdempotencyKey()
}

//endregion

//region ServiceMethodCallInputData

// ServiceMethodCallInputData abstracts method request body.
type ServiceMethodCallInputData interface {
}

//endregion

//region ServiceMethodCallInvocationError

// ServiceMethodCallInvocationError is a sub-class of ServiceMethodError which
// indicates that there's an error in the request.
//
// This error class is analogous to HTTP's 4xx status codes.
//
//TODO: sub-classes: acces, parameters, context
type ServiceMethodCallInvocationError interface {
	ServiceMethodError

	AZServiceMethodCallInvocationError()
}

//endregion

type ServiceMethodCallInput[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT],
	SessionT Session[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT],
	ServiceMethodIdempotencyKeyT ServiceMethodIdempotencyKey,
	InputContextT ServiceMethodCallInputContext[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT, SessionSubjectT, SessionT,
		ServiceMethodIdempotencyKeyT],
	InputDataT ServiceMethodCallInputData,
] struct {
	Context InputContextT
	Data    InputDataT
}

// ServiceMethodCallInputContext provides an abstraction for all input contexts
// in method call inputs.
type ServiceMethodCallInputContext[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT],
	SessionT Session[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT],
	ServiceMethodIdempotencyKeyT ServiceMethodIdempotencyKey,
] interface {
	ServiceMethodCallContext

	AZServiceMethodCallInputContext()

	// Session returns the session for this context.
	Session() SessionT

	// IdempotencyKey is a key used to ensure that a distinct operation is
	// performed at most once.
	//
	// This key is different from request ID, where for the same operation,
	// there could be more than one requests in attempt to retry in the event
	// of transit error.
	IdempotencyKey() ServiceMethodIdempotencyKeyT

	// OriginInfo returns information about the system that made the call.
	OriginInfo() ServiceMethodCallOriginInfo
}

// ServiceMethodCallInputContextError provides information for
// request-context-related error. It is a sub-class of
// ServiceMethodCallInputError.
type ServiceMethodCallInputContextError interface {
	ServiceMethodCallInvocationError

	AZServiceMethodCallInputContextError()
}

// ServiceMethodCallSessionError is a sub-class of
// ServiceMethodCallInputContextError specialized for indicating error
// in the session.
type ServiceMethodCallSessionError interface {
	ServiceMethodCallInputContextError

	AZServiceMethodCallSessionError()
}

//region ServiceMethodCallOutputContext

// ServiceMethodCallOutputContext provides an abstraction for all output contexts
// in method call outputs.
//
//TODO: listing of affected states with their respective revision ID.
//TODO: directive: done/end, redirect, retry (on failure; optionally with
// timing and retry count parameters or exponentially back-off parameters)
type ServiceMethodCallOutputContext interface {
	//TODO: ServiceMethodCallContext, or keep it like this and
	// add a method to access the input context: MethodCallInputContext()
	ServiceMethodContext

	AZServiceMethodCallOutputContext()

	// // Succeed returns true when the method achieved its objective.
	// Succeed() bool

	// Returns the error, if any.
	ServiceMethodErr() ServiceMethodError

	// Mutated returns true if the method made any changes to any state in the
	// server, even when the method did not succeed. It should not
	// return true if the change has been completely rolled-back before the
	// method returned this context.
	Mutated() bool
}

//endregion

//region ServiceMethodCallOutputContextBase

// ServiceMethodCallOutputContextBase is a base
// for ServiceMethodCallOutputContext implementations.
type ServiceMethodCallOutputContextBase struct {
	context.Context

	err     ServiceMethodError
	mutated bool
}

var _ ServiceMethodCallOutputContext = ServiceMethodCallOutputContextBase{}

// NewMethodCallOutputContext creates a new instance
// of ServiceMethodCallOutputContext.
func NewMethodCallOutputContext(
	err ServiceMethodError,
	mutated bool,
) ServiceMethodCallOutputContextBase {
	return ServiceMethodCallOutputContextBase{err: err, mutated: mutated}
}

// AZContext is required for conformance with Context.
func (ServiceMethodCallOutputContextBase) AZContext() {}

// AZServiceContext is required for conformance with ServiceContext.
func (ServiceMethodCallOutputContextBase) AZServiceContext() {}

// AZServiceMethodContext is required
// for conformance with ServiceMethodContext.
func (ServiceMethodCallOutputContextBase) AZServiceMethodContext() {}

// AZServiceMethodCallOutputContext is required
// for conformance with ServiceMethodCallOutputContext.
func (ServiceMethodCallOutputContextBase) AZServiceMethodCallOutputContext() {}

// ServiceMethodErr is required for conformance with ServiceMethodCallOutputContext.
func (ctx ServiceMethodCallOutputContextBase) ServiceMethodErr() ServiceMethodError { return ctx.err }

// Mutated is required for conformance with ServiceMethodCallOutputContext.
func (ctx ServiceMethodCallOutputContextBase) Mutated() bool { return ctx.mutated }

//endregion
