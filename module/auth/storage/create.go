package authstorage

import (
	"context"
	authmodel "edx_go/module/auth/model"
)

func (store *sqlStore) Create(ctx context.Context, data *authmodel.UserRegister) (*authmodel.User, error) {
	db := store.db

	user := authmodel.User{
		Password: data.Password,
		Email:    data.Email,
		Username: data.Username,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
