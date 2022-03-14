package authbusiness

import (
	"context"
	"edx_go/common"
	authmodel "edx_go/module/auth/model"
	"errors"
)

type CreateStore interface {
	Create(ctx context.Context, data *authmodel.UserRegister) (*authmodel.User, error)
	FindByEmail(ctx context.Context, email string) (*authmodel.User, error)
}

type registerUserBiz struct {
	store CreateStore
}

func NewRegisterUserBiz(store CreateStore) *registerUserBiz {
	return &registerUserBiz{
		store: store,
	}
}

func (biz *registerUserBiz) Register(ctx context.Context, data *authmodel.UserRegister) (*authmodel.User, error) {

	if user, _ := biz.store.FindByEmail(ctx, data.Email); user != nil {
		return nil, common.ErrEntityExisted("user", errors.New("Người dùng này đã tồn tại"))
	}

	user, err := biz.store.Create(ctx, data)

	if err != nil {
		return nil, err
	}

	return user, nil
}
