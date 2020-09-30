package azcore

// ServiceConfig provides a contract for all of its implementations.
type ServiceConfig interface {
	AZServiceConfig() ServiceConfig
}

// Service provides an abstraction for all services.
type Service interface {
	AZService() Service
}

// ServiceModule provides an abstraction for all kind of service modules.
type ServiceModule interface {
	AZServiceModule() ServiceModule
}

/**/ /**/

// ServiceServerConfig holds the configuration for a service server.
type ServiceServerConfig interface {
	ServiceConfig
}

/**/ /**/

// ServiceClientConfig holds the configuration for a service client.
type ServiceClientConfig interface {
	ServiceConfig
}

// ServiceClient provides an abstraction for all service clients.
type ServiceClient interface {
	AZService() Service
	AZServiceClient() ServiceClient
}

// ServiceClientModule provides all the required to instantiate a service
// client.
type ServiceClientModule struct {
	ServiceClientConfigSkeleton func() ServiceClientConfig
	NewServiceClient            func(ServiceClientConfig) (ServiceClient, Error)
}

var _ ServiceModule = ServiceClientModule{}

// AZServiceModule is required for conformance with ServiceModule.
func (clientMod ServiceClientModule) AZServiceModule() ServiceModule { return clientMod }
