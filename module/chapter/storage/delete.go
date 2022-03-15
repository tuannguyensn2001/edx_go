package chapterstorage

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
)

func (store *sqlStore) Delete(ctx context.Context, chapterId int) error {
	if err := store.db.Where("id = ?", chapterId).Delete(&chaptermodel.Chapter{}).Error; err != nil {
		return err
	}
	return nil
}
