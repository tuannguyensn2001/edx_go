package chapterstorage

import (
	"context"
	chaptermodel "edx_go/module/chapter/model"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) GetChapterByCourseAndUser(ctx context.Context, userId int, courseId int) ([]chaptermodel.Chapter, error) {
	var chapters []chaptermodel.Chapter
	if err := store.db.Where("course_id = ?", courseId).Find(&chapters).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return chapters, nil
}
