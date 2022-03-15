package business

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
)

type chapterGetByCourseAndUserStore interface {
	GetChapterByCourseAndUser(ctx context.Context, userId int, courseId int) ([]chaptermodel.Chapter, error)
}

type chapterGetByCourseAndUserBiz struct {
	store chapterGetByCourseAndUserStore
}

func NewChapterGetByCourseAndUserBiz(store chapterGetByCourseAndUserStore) *chapterGetByCourseAndUserBiz {
	return &chapterGetByCourseAndUserBiz{
		store: store,
	}
}

func (biz *chapterGetByCourseAndUserBiz) GetByCourseAndUser(ctx context.Context, userId int, courseId int) ([]chaptermodel.Chapter, error) {
	chapters, err := biz.store.GetChapterByCourseAndUser(ctx, userId, courseId)

	if err != nil {
		return nil, err
	}

	return chapters, nil

}
