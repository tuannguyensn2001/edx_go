package business

import "context"

type chapterDeleteStore interface {
	Delete(ctx context.Context, chapterId int) error
}

type chapterDeleteBiz struct {
	store chapterDeleteStore
}

func NewChapterDeleteBiz(store chapterDeleteStore) *chapterDeleteBiz {
	return &chapterDeleteBiz{store: store}
}

func (biz *chapterDeleteBiz) Delete(ctx context.Context, chapterId int) error {
	if err := biz.store.Delete(ctx, chapterId); err != nil {
		return err
	}

	return nil
}
