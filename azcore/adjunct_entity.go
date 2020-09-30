package azcore

// AdjunctEntityID abstracts adjunct entity IDs.
type AdjunctEntityID interface {
	AZAdjunctEntityID() AdjunctEntityID
}

// AdjunctEntityRefKey abstracts adjunct entity ref keys.
type AdjunctEntityRefKey interface {
	AZAdjunctEntityRefKey() AdjunctEntityRefKey
}

// AdjunctEntityService abstracts adjunct entity services.
type AdjunctEntityService interface {
	Service

	AZAdjunctEntityService() AdjunctEntityService
}

// AdjunctEntityServiceBase provides a base
// for AdjunctEntityService implementations. This implementation is shared
// by client and server implementations.
type AdjunctEntityServiceBase struct {
}

var _ AdjunctEntityService = &AdjunctEntityServiceBase{}

// AZAdjunctEntityService is required for conformance with AdjunctEntityService.
func (svc *AdjunctEntityServiceBase) AZAdjunctEntityService() AdjunctEntityService { return svc }

// AZService is required for conformance with Service.
func (svc *AdjunctEntityServiceBase) AZService() Service { return svc }

/**/ /**/

// AdjunctEntityServiceClient abstracts adjunct entity
// service client implementations.
type AdjunctEntityServiceClient interface {
	AdjunctEntityService
	ServiceClient

	AZAdjunctEntityServiceClient() AdjunctEntityServiceClient
}

// AdjunctEntityServiceClientBase provides a base
// for AdjunctEntityServiceClient implementations.
type AdjunctEntityServiceClientBase struct {
	AdjunctEntityServiceBase
}

var _ AdjunctEntityServiceClient = &AdjunctEntityServiceClientBase{}

// AZAdjunctEntityServiceClient is required
// for comformance with AdjunctEntityServiceClient.
func (svc *AdjunctEntityServiceClientBase) AZAdjunctEntityServiceClient() AdjunctEntityServiceClient {
	return svc
}

// AZServiceClient is required
// for conformance with ServiceClient.
func (svc *AdjunctEntityServiceClientBase) AZServiceClient() ServiceClient { return svc }
