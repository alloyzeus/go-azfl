package azcore

//region Service

// AdjunctEntityService abstracts adjunct entity services.
type AdjunctEntityService interface {
	Service
}

//endregion

//region ServiceBase

// AdjunctEntityServiceBase provides a base
// for AdjunctEntityService implementations. This implementation is shared
// by client and server implementations.
type AdjunctEntityServiceBase struct{}

var _ AdjunctEntityService = &AdjunctEntityServiceBase{}

//endregion

//region ServiceClient

// AdjunctEntityServiceClient abstracts adjunct entity
// service client implementations.
type AdjunctEntityServiceClient interface {
	AdjunctEntityService
	ServiceClient
}

//endregion

//region ServiceClientBase

// AdjunctEntityServiceClientBase provides a base
// for AdjunctEntityServiceClient implementations.
type AdjunctEntityServiceClientBase struct {
	AdjunctEntityServiceBase
}

var _ AdjunctEntityServiceClient = &AdjunctEntityServiceClientBase{}

//endregion

//region ServiceServer

// AdjunctEntityServiceServer abstracts adjunct entity
// service client implementations.
type AdjunctEntityServiceServer interface {
	AdjunctEntityService
	ServiceServer
}

//endregion

//region ServiceServerBase

// AdjunctEntityServiceServerBase provides a base
// for AdjunctEntityServiceServer implementations.
type AdjunctEntityServiceServerBase struct {
	AdjunctEntityServiceBase
}

var _ AdjunctEntityServiceServer = &AdjunctEntityServiceServerBase{}

//endregion
