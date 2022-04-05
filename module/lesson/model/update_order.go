package lessonmodel

type UpdateOrderDTO struct {
	Ids []int `form:"ids" binding:"required"`
	//ChapterId int   `form:"chapterId" binding:"required"`
}
