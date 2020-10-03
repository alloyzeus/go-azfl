package azcore

//region Service

// AdjunctEntityService abstracts adjunct entity services.
type AdjunctEntityService interface {
	Service

	AZAdjunctEntityService()
}

//endregion

//region ServiceBase

// AdjunctEntityServiceBase provides a base
// for AdjunctEntityService implementations. This implementation is shared
// by client and server implementations.
type AdjunctEntityServiceBase struct{}

var _ AdjunctEntityService = &AdjunctEntityServiceBase{}

// AZAdjunctEntityService is required
// for conformance with AdjunctEntityService.
func (*AdjunctEntityServiceBase) AZAdjunctEntityService() {}

// AZService is required
// for conformance with Service.
func (*AdjunctEntityServiceBase) AZService() {}

//endregion

//region ServiceClient

// AdjunctEntityServiceClient abstracts adjunct entity
// service client implementations.
type AdjunctEntityServiceClient interface {
	AdjunctEntityService
	ServiceClient

	AZAdjunctEntityServiceClient()
}

//endregion

//region ServiceClientBase

// AdjunctEntityServiceClientBase provides a base
// for AdjunctEntityServiceClient implementations.
type AdjunctEntityServiceClientBase struct {
	AdjunctEntityServiceBase
}

var _ AdjunctEntityServiceClient = &AdjunctEntityServiceClientBase{}

// AZAdjunctEntityServiceClient is required
// for comformance with AdjunctEntityServiceClient.
func (*AdjunctEntityServiceClientBase) AZAdjunctEntityServiceClient() {}

// AZServiceClient is required
// for conformance with ServiceClient.
func (*AdjunctEntityServiceClientBase) AZServiceClient() {}

//endregion

//region ServiceServer

// AdjunctEntityServiceServer abstracts adjunct entity
// service client implementations.
type AdjunctEntityServiceServer interface {
	AdjunctEntityService
	ServiceServer

	AZAdjunctEntityServiceServer()
}

//endregion

//region ServiceServerBase

// AdjunctEntityServiceServerBase provides a base
// for AdjunctEntityServiceServer implementations.
type AdjunctEntityServiceServerBase struct {
	AdjunctEntityServiceBase
}

var _ AdjunctEntityServiceServer = &AdjunctEntityServiceServerBase{}

// AZAdjunctEntityServiceServer is required
// for comformance with AdjunctEntityServiceServer.
func (*AdjunctEntityServiceServerBase) AZAdjunctEntityServiceServer() {}

// AZServiceServer is required
// for conformance with ServiceServer.
func (*AdjunctEntityServiceServerBase) AZServiceServer() {}

//endregion
