package coursestorage

import (
	"context"
	coursemodel "edx_go/module/course/model"
)

func (store *sqlStore) FindByNameAndUserId(ctx context.Context, name string, userId int) (*coursemodel.Course, error) {
	var course coursemodel.Course

	if err := store.db.Where("name = ? and user_id = ?", name, userId).First(&course).Error; err != nil {
		return nil, err
	}

	return &course, nil
}
