package lessonmodel

type LessonCreate struct {
	Name      string `form:"name" binding:"required"`
	VideoURL  string `form:"videoURL" binding:"required"`
	ChapterId int    `form:"chapterId" binding:"required"`
}
