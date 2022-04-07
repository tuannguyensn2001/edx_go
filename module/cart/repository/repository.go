package cartrepository

import (
	"context"
	"edx_go/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	Create(ctx context.Context, userId int, courseId int) (*models.Cart, error)
	FindByUserId(ctx context.Context, userId int) ([]int, error)
	Delete(ctx context.Context, userId int, courseId int) error
}

type repository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) Create(ctx context.Context, userId int, courseId int) (*models.Cart, error) {
	cart := models.Cart{UserId: userId, CourseId: courseId}

	if err := repo.db.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (repo *repository) FindByUserId(ctx context.Context, userId int) ([]int, error) {

	var courseIds []int

	if err := repo.db.Raw("SELECT course_id FROM carts WHERE user_id = ?", userId).Scan(&courseIds).Error; err != nil {
		return nil, err
	}

	return courseIds, nil

}

func (repo *repository) Delete(ctx context.Context, userId int, courseId int) error {
	err := repo.db.Exec("DELETE FROM carts WHERE user_id = ? AND course_id = ?", userId, courseId).Error

	if err != nil {
		return err
	}

	return nil
}
