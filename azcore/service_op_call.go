package azcore

import (
	"context"
	"time"

	"github.com/alloyzeus/go-azfl/v2/azob"
	"golang.org/x/text/language"
)

type ServiceOpCallInputMetadata interface {
	//TODO: request id / correlation id / idemptotency key, receive time by
	// the handler.
}

// ServiceOpCallOriginInfo holds information about a call's origin.
//
// TODO: key-value custom data
type ServiceOpCallOriginInfo struct {
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

//region ServiceOpIdempotencyKey

// ServiceOpIdempotencyKey represents the identifier of a method call. This
// identifier doubles as idempotency token.
type ServiceOpIdempotencyKey interface {
	azob.Equatable
}

//endregion

//region ServiceOpCallInputData

// ServiceOpCallData abstracts method request body.
type ServiceOpCallData interface {
}

//endregion

//region ServiceOpCallInvocationError

// ServiceOpCallInvocationError is a sub-class of ServiceOpError which
// indicates that there's an error in the request.
//
// This error class is analogous to HTTP's 4xx status codes.
//
// TODO: sub-classes: acces, parameters, context
type ServiceOpCallInvocationError interface {
	ServiceOpError
}

//endregion

type ServiceOpCallInput[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT],
	SessionT Session[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT, SessionT],
	ServiceOpIdempotencyKeyT ServiceOpIdempotencyKey,
	ContextT ServiceOpCallContext[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT, SessionSubjectT, SessionT,
		ServiceOpIdempotencyKeyT],
	DataT ServiceOpCallData,
] struct {
	Context ContextT
	Data    DataT
}

// ServiceOpCallContext provides an abstraction for all input contexts
// in method call inputs.
type ServiceOpCallContext[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT],
	SessionT Session[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT, SessionT],
	ServiceOpIdempotencyKeyT ServiceOpIdempotencyKey,
] interface {
	ServiceOpContext

	// Session returns the session for this context.
	Session() SessionT

	// IdempotencyKey is a key used to ensure that a distinct operation is
	// performed at most once.
	//
	// This key is different from request ID, where for the same operation,
	// there could be more than one requests in attempt to retry in the event
	// of transit error.
	IdempotencyKey() ServiceOpIdempotencyKeyT

	// OriginInfo returns information about the system that made the call.
	OriginInfo() ServiceOpCallOriginInfo
}

// ServiceOpCallContextError provides information for
// request-context-related error. It is a sub-class of
// ServiceOpCallInputError.
type ServiceOpCallContextError interface {
	ServiceOpCallInvocationError
}

// ServiceOpCallSessionError is a sub-class of
// ServiceOpCallContextError specialized for indicating error
// in the session.
type ServiceOpCallSessionError interface {
	ServiceOpCallContextError
}

//region ServiceOpResultContext

// ServiceOpResultContext provides an abstraction for all output contexts
// in method call outputs.
//
// TODO: listing of affected states with their respective revision ID.
// TODO: directive: done/end, redirect, retry (on failure; optionally with
// timing and retry count parameters or exponentially back-off parameters)
type ServiceOpResultContext interface {
	//TODO: ServiceOpCallContext, or keep it like this and
	// add a method to access the input context: OpCallContext()
	ServiceOpContext

	// // Succeed returns true when the method achieved its objective.
	// Succeed() bool

	// Returns the error, if any.
	ServiceOpErr() ServiceOpError

	// Mutated returns true if the method made any changes to any state in the
	// server, even when the method did not succeed. It should not
	// return true if the change has been completely rolled-back before the
	// method returned this context.
	Mutated() bool
}

//endregion

//region ServiceOpResultContextBase

// ServiceOpResultContextBase is a base
// for ServiceOpResultContext implementations.
type ServiceOpResultContextBase struct {
	context.Context

	opName     string
	resourceID string
	err        ServiceOpError
	mutated    bool
}

var _ ServiceOpResultContext = ServiceOpResultContextBase{}

// NewOpResultContext creates a new instance
// of ServiceOpResultContext.
func NewOpResultContext(
	opName string,
	resourceID string,
	err ServiceOpError,
	mutated bool,
) ServiceOpResultContextBase {
	return ServiceOpResultContextBase{
		opName:     opName,
		resourceID: resourceID,
		err:        err,
		mutated:    mutated}
}

// OperationName is required for conformance with ServiceOpContext.
func (ctx ServiceOpResultContextBase) OperationName() string { return ctx.opName }

// ResourceID is required for conformance with ServiceOpContext.
func (ctx ServiceOpResultContextBase) ResourceID() string { return ctx.resourceID }

// ServiceOpErr is required for conformance with ServiceOpResultContext.
func (ctx ServiceOpResultContextBase) ServiceOpErr() ServiceOpError { return ctx.err }

// Mutated is required for conformance with ServiceOpResultContext.
func (ctx ServiceOpResultContextBase) Mutated() bool { return ctx.mutated }

//endregion
