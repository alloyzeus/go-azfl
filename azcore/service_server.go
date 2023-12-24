package azcore

/**/ /**/

// ServiceServerConfig holds the configuration for a service client.
type ServiceServerConfig interface {
	ServiceConfig
}

// ServiceServer provides an abstraction for all service clients.
type ServiceServer interface {
	Service
}

// ServiceServerModule provides all the required to instantiate a service
// client.
type ServiceServerModule struct {
	ServiceServerConfigSkeleton func() ServiceServerConfig
	NewServiceServer            func(ServiceServerConfig) (ServiceServer, ServiceServerError)
}

var _ ServiceModule = ServiceServerModule{}

// ServiceServerError is an abstraction for all errors emitted by a
// service server.
type ServiceServerError interface {
	ServiceError
}
