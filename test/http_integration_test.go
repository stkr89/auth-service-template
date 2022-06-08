package test

import (
	"github.com/joho/godotenv"
	"github.com/stkr89/go-auth-service-template/cmd/server"
	"github.com/stkr89/go-auth-service-template/common"
	"github.com/stkr89/go-auth-service-template/endpoints"
	"github.com/stkr89/go-auth-service-template/service"
	transport "github.com/stkr89/go-auth-service-template/transports"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type HTTPIntegrationTestSuite struct {
	suite.Suite
	handler http.Handler
}

func (suite *HTTPIntegrationTestSuite) SetupSuite() {
	err := godotenv.Load("../.env")
	suite.NoError(err)

	e := endpoints.MakeEndpoints(service.NewAuthServiceImpl())
	server.StartServer(common.NewLogger(), e, false, true)

	suite.handler = transport.NewHTTPHandler(e)
}

// new test cases

func TestHTTPIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPIntegrationTestSuite))
}
