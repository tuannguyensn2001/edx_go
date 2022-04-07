package authbusiness

import (
	"context"
	"edx_go/models"
	authstorage "edx_go/module/auth/storage"
	"gorm.io/gorm"
)

type repository interface {
	FindById(ctx context.Context, userId int) (*models.User, error)
}

type service struct {
	repository repository
}

func NewAuthService(db *gorm.DB) *service {

	repo := authstorage.NewSQLStore(db)

	return &service{
		repository: repo,
	}
}
