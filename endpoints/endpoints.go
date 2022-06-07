package endpoints

import (
	"context"
	"github.com/stkr89/authsvc/types"

	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/authsvc/service"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

func MakeEndpoints(s service.AuthService) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
	}
}

func makeAddEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.MathRequest)
		return s.Add(req)
	}
}
