package cartstruct

type DeleteInCartInput struct {
	CourseId int `form:"course_id" binding:"required"`
}
