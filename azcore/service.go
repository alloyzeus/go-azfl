package azcore

// Service provides an abstraction for all services.
type Service interface {
}

// ServiceConfig provides an abstractions for all service-related configs.
type ServiceConfig interface {
}

// ServiceContext provides an abstraction for all service-related contexts.
type ServiceContext interface {
}

// ServiceModule provides an abstraction for modular services.
type ServiceModule interface {
}

// ServiceError is an abstraction for all errors emitted by a service.
type ServiceError interface {
	Error
}
