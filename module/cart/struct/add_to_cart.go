package cartstruct

type AddToCartInput struct {
	CourseId int `form:"course_id" binding:"required"`
}
