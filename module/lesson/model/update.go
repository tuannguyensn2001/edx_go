package lessonmodel

type LessonEdit struct {
	Name     string `form:"name" binding:"required" json:"name"`
	VideoURL string `form:"videoURL" binding:"required" json:"videoURL"`
	Id       int    `json:"id"`
}
