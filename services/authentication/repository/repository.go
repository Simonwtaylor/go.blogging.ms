package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	CreateUser(user interface{}) error
	Login(username string, password string) error
}

type AuthenticationRepository struct {
	DB *gorm.DB
}

func (a *AuthenticationRepository) CreateUser(user interface{}) error {
	return nil
}

func (a *AuthenticationRepository) Login(username string, password string) error {
	return nil
}

func ProvideAuthenticationRepository(db *gorm.DB) Repository {
	return &AuthenticationRepository{DB:db}
}