package azcore

/**/ /**/

// ServiceClientConfig holds the configuration for a service client.
type ServiceClientConfig interface {
	ServiceConfig
}

// ServiceClient provides an abstraction for all service clients.
type ServiceClient interface {
	Service
}

// ServiceClientModule provides all the required to instantiate a service
// client.
type ServiceClientModule struct {
	ServiceClientConfigSkeleton func() ServiceClientConfig
	NewServiceClient            func(ServiceClientConfig) (ServiceClient, ServiceClientError)
}

var _ ServiceModule = ServiceClientModule{}

// ServiceClientError is an abstraction for all errors emitted by a
// service server.
type ServiceClientError interface {
	ServiceError
}
