package azcore

/**/ /**/

// ServiceClientConfig holds the configuration for a service client.
type ServiceClientConfig interface {
	ServiceConfig
}

// ServiceClient provides an abstraction for all service clients.
type ServiceClient interface {
	Service
	AZServiceClient()
}

// ServiceClientModule provides all the required to instantiate a service
// client.
type ServiceClientModule struct {
	ServiceClientConfigSkeleton func() ServiceClientConfig
	NewServiceClient            func(ServiceClientConfig) (ServiceClient, Error)
}

var _ ServiceModule = ServiceClientModule{}

// AZServiceModule is required for conformance with ServiceModule.
func (ServiceClientModule) AZServiceModule() {}
