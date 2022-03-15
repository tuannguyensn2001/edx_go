package business

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
)

type chapterCreateStore interface {
	Create(ctx context.Context, data *chaptermodel.Chapter) error
	CountByCourseId(ctx context.Context, courseId int) (int64, error)
}

type ChapterCreateBiz struct {
	store chapterCreateStore
}

func NewChapterCreateBiz(store chapterCreateStore) *ChapterCreateBiz {
	return &ChapterCreateBiz{
		store: store,
	}
}

func (biz *ChapterCreateBiz) CreateChapter(ctx context.Context, data *chaptermodel.ChapterCreate) (*chaptermodel.Chapter, error) {

	count, err := biz.store.CountByCourseId(ctx, data.CourseId)

	if err != nil {
		return nil, err
	}

	chapter := &chaptermodel.Chapter{
		Name:     data.Name,
		CourseId: data.CourseId,
		Order:    int(count + 1),
	}

	err = biz.store.Create(ctx, chapter)

	if err != nil {
		return nil, err
	}

	return chapter, nil

}
