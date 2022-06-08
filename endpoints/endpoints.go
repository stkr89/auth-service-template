package endpoints

import (
	"context"
	"github.com/stkr89/auth-service-template/types"

	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/auth-service-template/service"
)

type Endpoints struct {
	SignUp endpoint.Endpoint
}

func MakeEndpoints(s service.AuthService) Endpoints {
	return Endpoints{
		SignUp: makeSignUpEndpoint(s),
	}
}

func makeSignUpEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.SignUpRequest)
		return s.SignUp(ctx, req)
	}
}
