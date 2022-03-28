package lessonstorage

import (
	"context"
	"edx_go/models"
)

func (store *sqlStore) GetLatestOrder(ctx context.Context, chapterId int) (int64, error) {
	var count int64
	if err := store.db.Model(&models.Lesson{}).Where("chapter_id = ?", chapterId).Count(&count).Error; err != nil {
		return -1, err
	}

	return count, nil
}
