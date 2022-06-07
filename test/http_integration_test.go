package test

import (
	"github.com/joho/godotenv"
	"github.com/stkr89/authsvc/cmd/server"
	"github.com/stkr89/authsvc/common"
	"github.com/stkr89/authsvc/endpoints"
	"github.com/stkr89/authsvc/service"
	transport "github.com/stkr89/authsvc/transports"
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
