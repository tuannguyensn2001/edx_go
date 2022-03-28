package lessonstorage

import (
	"context"
	"edx_go/models"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) GetByChapterId(context context.Context, chapterId int) ([]models.Lesson, error) {
	var lessons []models.Lesson

	if err := store.db.Where("chapter_id = ?", chapterId).Find(&lessons).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else {
			return lessons, nil
		}
	}

	return lessons, nil

}
