package chaptermodel

type ChapterCreate struct {
	Name     string `form:"name" binding:"required"`
	CourseId int    `form:"courseId" binding:"required"`
}
