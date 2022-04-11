package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id               int        `gorm:"column:id" json:"id"`
	Content          string     `gorm:"column:content" json:"content"`
	UserId           int        `gorm:"column:user_id" json:"userId"`
	ParentId         int        `gorm:"column:parent_id" json:"parentId"`
	CommentableModel string     `gorm:"column:commentable_model" json:"commentableModel"`
	CommentableId    int        `gorm:"column:commentable_id" json:"commentableId"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt        gorm.DeletedAt
}
