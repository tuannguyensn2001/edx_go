package lessonbusiness

import (
	"context"
	"edx_go/models"
)

type store interface {
	GetByChapterId(context context.Context, chapterId int) ([]models.Lesson, error)
}

type biz struct {
	store store
}

func NewLessonGetByChapterIdBiz(store store) *biz {
	return &biz{
		store: store,
	}
}

func (biz *biz) GetByChapterIdBiz(ctx context.Context, chapterId int) ([]models.Lesson, error) {
	lessons, err := biz.store.GetByChapterId(ctx, chapterId)

	if err != nil {
		return nil, err
	}

	return lessons, nil
}
