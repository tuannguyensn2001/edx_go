package cartservice

import (
	"context"
	"edx_go/common"
	"edx_go/models"
	"edx_go/module/cart/repository"
	"errors"
	"gorm.io/gorm"
)

type service struct {
	repository cartrepository.CartRepository
}

func NewCartService(db *gorm.DB) *service {
	repository := cartrepository.NewCartRepository(db)

	return &service{repository: repository}

}

func (s *service) AddToCart(ctx context.Context, userId int, courseId int) (*models.Cart, error) {

	courseIds, err := s.repository.FindByUserId(ctx, userId)

	if err != nil {
		return nil, err
	}

	for _, item := range courseIds {
		if item == courseId {
			return nil, common.NewCustomError(errors.New("Đã tồn tại trong giỏ hàng"), "Đã tồn tại trong giỏ hàng", "cart")
		}
	}

	cart, err := s.repository.Create(ctx, userId, courseId)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *service) Delete(ctx context.Context, userId int, courseId int) error {
	return s.repository.Delete(ctx, userId, courseId)
}
