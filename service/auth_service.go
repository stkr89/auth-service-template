package service

import (
	"github.com/stkr89/authsvc/common"
	"github.com/stkr89/authsvc/dao"
	"github.com/stkr89/authsvc/models"
	"github.com/stkr89/authsvc/types"

	"github.com/go-kit/kit/log"
)

// AuthService interface
type AuthService interface {
	CreateUser(request *types.CreateUserRequest) (*types.CreateUserResponse, error)
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

func (s AuthServiceImpl) CreateUser(request *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	createdUser, err := s.authDao.CreateUser(&models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateUserResponse{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
	}, nil
}
