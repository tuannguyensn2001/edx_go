package models

type Cart struct {
	UserId   int `gorm:"user_id" json:"userId"`
	CourseId int `gorm:"course_id" json:"courseId"`
}
