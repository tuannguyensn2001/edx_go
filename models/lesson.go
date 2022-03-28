package models

import (
	"gorm.io/gorm"
	"time"
)

type Lesson struct {
	ID        int        `json:"id" gorm:"column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	VideoURL  string     `json:"videoURL" gorm:"column:video_url"`
	VideoType string     `json:"videoType" gorm:"column:video_type"`
	Order     int        `json:"order" gorm:"column:order"`
	ChapterId int        `json:"chapterId" gorm:"column:chapter_id"`
	Status    string     `json:"status" gorm:"status"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt
}
