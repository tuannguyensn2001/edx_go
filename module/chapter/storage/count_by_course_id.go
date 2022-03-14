package chapterstorage

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
)

func (store *sqlStore) CountByCourseId(ctx context.Context, courseId int) (int64, error) {
	var count int64
	if err := store.db.Model(&chaptermodel.Chapter{}).Where("course_id = ?", courseId).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}
