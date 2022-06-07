package service

import (
	"github.com/stkr89/authsvc/common"
	"github.com/stkr89/authsvc/dao"
	"github.com/stkr89/authsvc/types"

	"github.com/go-kit/kit/log"
)

// AuthService interface
type AuthService interface {
	Add(request *types.MathRequest) (*types.MathResponse, error)
}

type AuthServiceImpl struct {
	logger  log.Logger
	authDao dao.AuthDao
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{
		logger:  common.NewLogger(),
		authDao: dao.NewAuthDaoImpl(),
	}
}

func (s AuthServiceImpl) Add(request *types.MathRequest) (*types.MathResponse, error) {
	return nil, nil
}
