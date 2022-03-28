package lessonbusiness

import (
	"context"
	"edx_go/models"
	lessonmodel "edx_go/module/lesson/model"
)

type lessonCreateStore interface {
	Create(ctx context.Context, lesson *lessonmodel.LessonCreate, order int) (*models.Lesson, error)
	GetLatestOrder(ctx context.Context, chapterId int) (int64, error)
}

type lessonCreateBiz struct {
	store lessonCreateStore
}

func NewLessonCreateBiz(store lessonCreateStore) *lessonCreateBiz {
	return &lessonCreateBiz{
		store: store,
	}
}

func (biz *lessonCreateBiz) LessonCreate(ctx context.Context, lesson *lessonmodel.LessonCreate) (*models.Lesson, error) {

	order, err := biz.store.GetLatestOrder(ctx, lesson.ChapterId)

	if err != nil {
		return nil, err
	}

	newLesson, err := biz.store.Create(ctx, lesson, int(order+1))

	if err != nil {
		return nil, err
	}

	return newLesson, nil
}
