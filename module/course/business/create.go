package authbusiness

import (
	"context"
	"edx_go/common"
	coursemodel "edx_go/module/course/model"
	"errors"
	"gorm.io/gorm"
)

type CreateStore interface {
	Create(ctx context.Context, data *coursemodel.CourseCreate) (*coursemodel.Course, error)
	FindByNameAndUserId(ctx context.Context, name string, userId int) (*coursemodel.Course, error)
}

type CreateCourseBiz struct {
	store CreateStore
}

func NewCreateCourseBiz(store CreateStore) *CreateCourseBiz {
	return &CreateCourseBiz{
		store: store,
	}
}

func (biz *CreateCourseBiz) CreateCourse(ctx context.Context, data *coursemodel.CourseCreate) (*coursemodel.Course, error) {

	course, err := biz.store.FindByNameAndUserId(ctx, data.Name, data.UserId)

	if course != nil {
		return nil, common.ErrEntityExisted("Course", err)
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrDB(err)
	}

	user, err := biz.store.Create(ctx, data)

	if err != nil {
		return nil, err
	}

	return user, nil

}
