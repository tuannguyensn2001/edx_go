package chaptermodel

type ChapterUpdate struct {
	Name string `form:"name" binding:"required"`
}
