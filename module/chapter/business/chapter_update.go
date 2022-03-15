package business

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
)

type chapterUpdateStore interface {
	Update(ctx context.Context, id int, data chaptermodel.ChapterUpdate) error
}

type chapterUpdateBiz struct {
	store chapterUpdateStore
}

func NewChapterUpdateBiz(store chapterUpdateStore) *chapterUpdateBiz {
	return &chapterUpdateBiz{store: store}
}

func (biz *chapterUpdateBiz) Update(ctx context.Context, id int, data chaptermodel.ChapterUpdate) error {
	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}
	return nil
}
