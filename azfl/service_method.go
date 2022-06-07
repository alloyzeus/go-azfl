package azcore

import (
	"context"
	"time"

	"github.com/alloyzeus/go-azfl/azfl/azob"
	"golang.org/x/text/language"
)

//region ServiceMethodError

// ServiceMethodError is a specialization of ServiceError which focuses
// on method-related errors.
type ServiceMethodError interface {
	ServiceError

	AZServiceMethodError()
}

// ServiceMethodErrorMsg is a basic implementation of ServiceMethodError
// which provides an error message string.
type ServiceMethodErrorMsg struct {
	msg string
}

var _ ServiceMethodError = ServiceMethodErrorMsg{}

// AZServiceMethodError is required for conformance with ServiceMethodError.
func (ServiceMethodErrorMsg) AZServiceMethodError() {}

// AZServiceError is required for conformance with ServiceError.
func (ServiceMethodErrorMsg) AZServiceError() {}

func (err ServiceMethodErrorMsg) Error() string { return err.msg }

// ServiceMethodInternalError represents error in the method, service, or in
// any dependency required to achieve the objective.
//
// This is analogous to HTTP's 5xx status code.
type ServiceMethodInternalError interface {
	ServiceMethodError

	AZServiceMethodInternalError()
}

// ServiceMethodInternalErrorMsg is a basic implementation of
// ServiceMethodInternalError which provides error message string.
type ServiceMethodInternalErrorMsg struct {
	msg string
}

var _ ServiceMethodInternalError = ServiceMethodInternalErrorMsg{}

// AZServiceMethodInternalError is required for conformance
// with ServiceMethodInternalError.
func (ServiceMethodInternalErrorMsg) AZServiceMethodInternalError() {}

// AZServiceMethodError is required for conformance with ServiceMethodError.
func (ServiceMethodInternalErrorMsg) AZServiceMethodError() {}

// AZServiceError is required for conformance with ServiceError.
func (ServiceMethodInternalErrorMsg) AZServiceError() {}

func (err ServiceMethodInternalErrorMsg) Error() string { return err.msg }

//endregion

// ErrServiceMethodNotImplemented is usually used when a method is unable to
// achieve its objective because some part of it is unimplemented.
//
// Analogous to HTTP's 501 status code and gRPC's 12 status code.
var ErrServiceMethodNotImplemented = &ServiceMethodInternalErrorMsg{msg: "not implemented"}

//region ServiceMethodContext

// ServiceMethodContext is an abstraction for input and output contexts used
// when calling a method.
type ServiceMethodContext interface {
	ServiceContext

	AZServiceMethodContext()
}

//endregion

//region ServiceMethodCallContext

type ServiceMethodCallContext interface {
	ServiceMethodContext

	AZServiceMethodCallContext()
}

//endregion

//region ServiceMethodCallInfo

// A ServiceMethodCallInfo describes a method call.
type ServiceMethodCallInfo interface {
	AZServiceMethodCallInfo()

	//TODO: only on mutating operation
	OpID() ServiceMethodOpID

	// RequestOriginTime is the time assigned by the terminal which made the
	// request for the method call.
	//
	// The value is untrusted.
	RequestOriginTime() *time.Time
}

//endregion

//region ServiceMethodCallOriginInfo

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

	// DateTime is the time of the device where this operation was initiated
	// from at the time the operation was posted.
	//
	// Analogous to HTTP Date header field.
	DateTime *time.Time
}

//endregion

//region ServiceMethodOpID

// ServiceMethodOpID represents the identifier of a method call. This
// identifier doubles as idempotency token.
type ServiceMethodOpID interface {
	azob.Equatable

	AZServiceMethodOpID()
}

//endregion

//region ServiceMethodMessage

// ServiceMethodMessage abstracts the messages, i.e., requests and responses.
type ServiceMethodMessage interface {
	AZServiceMethodMessage()

	// ServiceMethodContext returns the context of this message.
	//
	// Implementations must return most specialized context implementation.
	MethodContext() ServiceMethodContext
}

//endregion

//region ServiceMethodCallInput

// ServiceMethodCallInput abstracts method request messages.
type ServiceMethodCallInput[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	ServiceMethodMessage

	CallInputContext() ServiceMethodCallInputContext[
		SessionIDNumT, TerminalIDNumT, UserIDNumT, Session[
			SessionIDNumT, SessionRefKey[SessionIDNumT],
			TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
			UserIDNumT, UserRefKey[UserIDNumT],
			SessionSubject[
				TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
				UserIDNumT, UserRefKey[UserIDNumT],
			],
		],
	]
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

//region ServiceMethodCallInputContext

// ServiceMethodCallInputContext provides an abstraction for all input contexts
// in method call inputs.
type ServiceMethodCallInputContext[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
	SessionT Session[
		SessionIDNumT, SessionRefKey[SessionIDNumT],
		TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
		UserIDNumT, UserRefKey[UserIDNumT],
		SessionSubject[
			TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
			UserIDNumT, UserRefKey[UserIDNumT],
		],
	],
] interface {
	ServiceMethodCallContext

	AZServiceMethodCallInputContext()

	// Session returns the session for this context.
	Session() SessionT

	CallOriginInfo() ServiceMethodCallOriginInfo
}

//endregion

//region ServiceMethodCallInputContextError

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

//endregion

//region ServiceMethodCallOutput

// ServiceMethodCallOutput abstracts method response messages.
type ServiceMethodCallOutput interface {
	ServiceMethodMessage

	MethodCallOutputContext() ServiceMethodCallOutputContext
}

//endregion

//region ServiceMethodCallOutputContext

// ServiceMethodCallOutputContext provides an abstraction for all output contexts
// in method call outputs.
//
//TODO: listing of affected states with their respective revision ID.
//TODO: directive: done/end, redirect, retry (on failure; optionally with
// timing and retry count parameters or exponentially back-off parameters)
type ServiceMethodCallOutputContext interface {
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

//region Mutating method

// ServiceMutatingMethodMessage abstracts mutating method requests
// and responses.
type ServiceMutatingMethodMessage interface {
	ServiceMethodMessage

	MutatingMethodContext() ServiceMutatingMethodContext
}

// ServiceMutatingMethodContext abstracts contexts of mutating
// method requests and responses.
type ServiceMutatingMethodContext interface {
	ServiceMethodContext
}

// ServiceMutatingMethodCallInput abstracts mutating method requests.
type ServiceMutatingMethodCallInput[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	ServiceMutatingMethodMessage
	ServiceMethodCallInput[SessionIDNumT, TerminalIDNumT, UserIDNumT]

	MutatingOpCallInputContext() ServiceMutatingOpCallInputContext[SessionIDNumT, TerminalIDNumT, UserIDNumT]
}

// ServiceMutatingOpCallInputContext abstracts mutating method request contexts.
type ServiceMutatingOpCallInputContext[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	ServiceMutatingMethodContext
	ServiceMethodCallInputContext[
		SessionIDNumT, TerminalIDNumT, UserIDNumT, Session[
			SessionIDNumT, SessionRefKey[SessionIDNumT],
			TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
			UserIDNumT, UserRefKey[UserIDNumT],
			SessionSubject[
				TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
				UserIDNumT, UserRefKey[UserIDNumT],
			],
		],
	]
}

// ServiceMutatingMethodCallOutput abstracts mutating method responses.
type ServiceMutatingMethodCallOutput interface {
	ServiceMutatingMethodMessage
	ServiceMethodCallOutput

	MutatingMethodCallOutputContext() ServiceMutatingMethodCallOutputContext
}

// ServiceMutatingMethodCallOutputContext abstracts mutating
// method response contexts.
type ServiceMutatingMethodCallOutputContext interface {
	ServiceMutatingMethodContext
	ServiceMethodCallOutputContext
}

//endregion
