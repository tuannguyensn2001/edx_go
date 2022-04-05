package authstorage

import (
	"context"
	authmodel "edx_go/module/auth/model"
	hashprovider "edx_go/providers/hash"
)

func (store *sqlStore) Create(ctx context.Context, data *authmodel.UserRegister) (*authmodel.User, error) {
	db := store.db

	hashPassword, err := hashprovider.Hash(data.Password)

	if err != nil {
		return nil, err
	}

	user := authmodel.User{
		Password: hashPassword,
		Email:    data.Email,
		Username: data.Username,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
