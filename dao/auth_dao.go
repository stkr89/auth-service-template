package dao

import (
	"github.com/go-kit/log"
	"github.com/stkr89/authsvc/common"
	"github.com/stkr89/authsvc/config"
	"github.com/stkr89/authsvc/models"
	"gorm.io/gorm"
)

type AuthDao interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type AuthDaoImpl struct {
	logger log.Logger
	db     gorm.DB
}

func NewAuthDaoImpl() *AuthDaoImpl {
	return &AuthDaoImpl{
		logger: common.NewLogger(),
		db:     config.NewDB(),
	}
}

func (a AuthDaoImpl) GetUserByEmail(email string) (*models.User, error) {
	obj := models.User{}
	result := a.db.Where("email = ?", email).First(&obj)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		a.logger.Log("message", "failed to get by email", "email", email, "error", result.Error)
		return nil, common.SomethingWentWrong
	}

	return &obj, nil
}

func (a AuthDaoImpl) CreateUser(user *models.User) (*models.User, error) {
	result := a.db.Create(&user)
	if result.Error != nil {
		a.logger.Log("message", "failed to create", "error", result.Error)
		return nil, common.SomethingWentWrong
	}

	a.logger.Log("message", "created successfully", "return", user.ID)

	return user, nil
}
