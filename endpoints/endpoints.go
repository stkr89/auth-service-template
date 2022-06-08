package endpoints

import (
	"context"
	"github.com/stkr89/auth-service-template/types"

	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/auth-service-template/service"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
}

func MakeEndpoints(s service.AuthService) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.CreateUserRequest)
		return s.SignUp(ctx, req)
	}
}
