package chapterstorage

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
)

func (store *sqlStore) Update(ctx context.Context, id int, data chaptermodel.ChapterUpdate) error {
	if err := store.db.Model(chaptermodel.Chapter{}).Where("id = ?", id).Update("name", data.Name).Error; err != nil {
		return err
	}
	return nil
}
