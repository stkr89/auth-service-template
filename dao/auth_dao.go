package dao

import (
	"github.com/go-kit/log"
	"github.com/stkr89/authsvc/common"
	"github.com/stkr89/authsvc/config"
	"gorm.io/gorm"
)

type AuthDao interface {
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
