package commentstruct

type CommentInput struct {
	Content          string `form:"content" binding:"required"`
	UserId           int
	CommentableId    int    `form:"commentableId" binding:"required"`
	CommentableModel string `form:"commentableModel" binding:"required"`
}
