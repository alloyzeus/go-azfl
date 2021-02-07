package azcore

/**/ /**/

// ServiceServerConfig holds the configuration for a service client.
type ServiceServerConfig interface {
	ServiceConfig
}

// ServiceServer provides an abstraction for all service clients.
type ServiceServer interface {
	Service
	AZServiceServer()
}

// ServiceServerModule provides all the required to instantiate a service
// client.
type ServiceServerModule struct {
	ServiceServerConfigSkeleton func() ServiceServerConfig
	NewServiceServer            func(ServiceServerConfig) (ServiceServer, ServiceServerError)
}

var _ ServiceModule = ServiceServerModule{}

// AZServiceModule is required for conformance with ServiceModule.
func (ServiceServerModule) AZServiceModule() {}

// ServiceServerError is an abstraction for all errors emitted by a
// service server.
type ServiceServerError interface {
	ServiceError

	AZServiceServerError()
}

// ServiceServerMethodError is a specialization of ServiceServerError
// which focuses on method-related errors.
type ServiceServerMethodError interface {
	ServiceServerError
	ServiceMethodError

	AZServiceServerMethodError()
}
