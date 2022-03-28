package lessonstorage

import (
	"context"
	"edx_go/models"
	lessonmodel "edx_go/module/lesson/model"
)

func (store *sqlStore) Create(ctx context.Context, lesson *lessonmodel.LessonCreate, order int) (*models.Lesson, error) {
	newLesson := models.Lesson{
		Name:      lesson.Name,
		VideoURL:  lesson.VideoURL,
		VideoType: "youtube",
		Order:     order,
		ChapterId: lesson.ChapterId,
		Status:    "ACTIVE",
	}

	if err := store.db.Create(&newLesson).Error; err != nil {
		return nil, err
	}

	return &newLesson, nil
}
