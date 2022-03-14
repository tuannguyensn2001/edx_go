package coursestorage

import (
	"context"
	coursemodel "edx_go/module/course/model"
)

func (store *sqlStore) FindByUserId(ctx context.Context, userId int) ([]coursemodel.Course, error) {
	var result []coursemodel.Course

	if err := store.db.Where("user_id = ?", userId).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}
