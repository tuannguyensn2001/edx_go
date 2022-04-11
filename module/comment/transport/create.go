package commenttransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	commentrepository "edx_go/module/comment/repository"
	commentservice "edx_go/module/comment/service"
	commentstruct "edx_go/module/comment/struct"
	"edx_go/websocket"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateComment(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var input commentstruct.CommentInput

		if err := ctx.ShouldBind(&input); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userId, _ := ctx.Get("user_id")

		input.UserId = userId.(int)

		repository := commentrepository.NewCommentRepository(appCtx.GetDBConnection())
		service := commentservice.NewCommentService(repository)

		comment, err := service.Create(ctx.Request.Context(), input)

		if err != nil {
			panic(err)
		}
		
		if comment.CommentableModel == "lesson" {
			room := "lesson-detail-" + strconv.Itoa(comment.CommentableId)
			appCtx.GetHub().EmitToRoom(room, websocket.MessageWs{
				Channel: "receive-comment",
				RoomId:  room,
				Data:    comment,
			})
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Thêm mới bình luận thành công", comment))

	}
}
