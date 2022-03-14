package authstorage

import (
	"context"
	authmodel "edx_go/module/auth/model"
)

func (store *sqlStore) FindByEmail(ctx context.Context, email string) (*authmodel.User, error) {
	var user authmodel.User

	if err := store.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
