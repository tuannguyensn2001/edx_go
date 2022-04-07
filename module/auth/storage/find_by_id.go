package authstorage

import (
	"context"
	"edx_go/models"
)

func (store *sqlStore) FindById(ctx context.Context, userId int) (*models.User, error) {
	var user models.User

	if err := store.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}
