package authbusiness

import (
	"context"
	"edx_go/common"
	coursemodel "edx_go/module/course/model"
)

type store interface {
	FindByUserId(ctx context.Context, userId int) ([]coursemodel.Course, error)
}

type GetMyCoursesBiz struct {
	store store
}

func NewGetMyCoursesBiz(store store) *GetMyCoursesBiz {
	return &GetMyCoursesBiz{
		store: store,
	}
}

func (biz *GetMyCoursesBiz) FindByUserId(ctx context.Context, userId int) ([]coursemodel.Course, error) {
	course, err := biz.store.FindByUserId(ctx, userId)

	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	return course, nil
}
