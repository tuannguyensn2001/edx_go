package authbusiness

import (
	"context"
	"edx_go/common"
	"edx_go/models"
)

func (service *service) GetMe(ctx context.Context, userId int) (*models.User, error) {
	user, err := service.repository.FindById(ctx, userId)

	if err != nil {
		return nil, common.ErrEntityNotFound("user", err)
	}

	return user, nil

}
