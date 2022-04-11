package commentrepository

import (
	"context"
	"edx_go/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *models.Comment) error
}

type repository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) Create(ctx context.Context, comment *models.Comment) error {
	err := repo.db.Create(comment).Error

	return err

}
