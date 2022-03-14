package coursemodel

type CourseCreate struct {
	Name        string  `json:"name" form:"name" binding:"required"`
	Description string  `json:"description" form:"description" binding:"required"`
	ImageUrl    string  `json:"imageUrl" form:"imageUrl" binding:"required"`
	Price       float64 `form:"price" binding:"required"`
	Status      string  `form:"status" binding:"required"`
	UserId      int
}
