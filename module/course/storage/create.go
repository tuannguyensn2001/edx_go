package coursestorage

import (
	"context"
	coursemodel "edx_go/module/course/model"
)

func (store *sqlStore) Create(ctx context.Context, data *coursemodel.CourseCreate) (*coursemodel.Course, error) {
	user := coursemodel.Course{
		Name:        data.Name,
		Description: data.Description,
		ImageUrl:    data.ImageUrl,
		Status:      data.Status,
		Price:       data.Price,
		UserId:      data.UserId,
	}

	if err := store.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
