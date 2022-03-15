package chaptermodel

import (
	"gorm.io/gorm"
	"time"
)

type Chapter struct {
	ID        int        `json:"id" gorm:"column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	Order     int        `json:"order" gorm:"column:order"`
	CourseId  int        `json:"courseId" gorm:"column:course_id"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAT gorm.DeletedAt
}
