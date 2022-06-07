package service

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/go-kit/kit/log"
	"github.com/stkr89/authsvc/common"
	"github.com/stkr89/authsvc/dao"
	"github.com/stkr89/authsvc/models"
	"github.com/stkr89/authsvc/types"
	"gorm.io/gorm"
)

// AuthService interface
type AuthService interface {
	CreateUser(ctx context.Context, request *types.CreateUserRequest) (*types.CreateUserResponse, error)
}

type AuthServiceImpl struct {
	logger   log.Logger
	authDao  dao.AuthDao
	firebase *firebase.App
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{
		logger:   common.NewLogger(),
		authDao:  dao.NewAuthDaoImpl(),
		firebase: common.NewFirebaseApp(),
	}
}

func (s AuthServiceImpl) CreateUser(ctx context.Context, request *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	existingUser, err := s.authDao.GetUserByEmail(request.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Log("error", err)
		return nil, common.UserAlreadyExists
	}

	if existingUser.Email != "" {
		return nil, common.UserAlreadyExists
	}

	createdUser, err := s.authDao.CreateUser(&models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	})
	if err != nil {
		s.logger.Log("error", err)
		return nil, err
	}

	s.logger.Log("msg", "user created successfully", "id", createdUser.ID)

	firebaseClient, err := s.firebase.Auth(ctx)
	if err != nil {
		s.logger.Log("error", err)
		return nil, common.SomethingWentWrong
	}

	params := (&auth.UserToCreate{}).
		Email(request.Email).
		Password(request.Password).
		Disabled(false)
	u, err := firebaseClient.CreateUser(ctx, params)
	if err != nil {
		s.logger.Log("error", err)
		return nil, common.SomethingWentWrong
	}

	s.logger.Log("msg", "firebase user created successfully", "id", u.UID)

	return &types.CreateUserResponse{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
	}, nil
}
