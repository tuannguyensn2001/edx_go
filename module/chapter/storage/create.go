package chapterstorage

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
)

func (store *sqlStore) Create(ctx context.Context, data *chaptermodel.Chapter) error {
	if err := store.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
