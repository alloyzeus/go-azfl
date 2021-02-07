package azcore

// ServiceConfig provides a contract for all of its implementations.
type ServiceConfig interface {
	AZServiceConfig()
}

// Service provides an abstraction for all services.
type Service interface {
	AZService()
}

// ServiceModule provides an abstraction for all kind of service modules.
type ServiceModule interface {
	AZServiceModule()
}

// ServiceError is an abstraction for all errors emitted by a service.
type ServiceError interface {
	Error

	AZServiceError()
}

// ServiceMethodError is a specialization of ServiceError which focuses
// on method-related errors.
type ServiceMethodError interface {
	ServiceError

	AZServiceMethodError()
}
