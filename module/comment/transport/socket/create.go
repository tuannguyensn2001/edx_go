package commenttransportsocket

import (
	"context"
	"edx_go/websocket"
	"gorm.io/gorm"
)

type appContext interface {
	GetDBConnection() *gorm.DB
	GetHub() *websocket.Hub
}

func CreateComment(appCtx appContext) func(ctx context.Context, message websocket.MessageWs) {
	return func(ctx context.Context, message websocket.MessageWs) {
		//repo := commentrepository.NewCommentRepository(appCtx.GetDBConnection())
		//service := commentservice.NewCommentService(repo)
		//
		//comment, err := service.Create(ctx, message)
		//
		//if err != nil {
		//	return
		//}
		//
		//dataResponse := make(map[string]interface{})
		//
		//dataResponse["content"] = comment.Content
		//dataResponse["lessonId"] = comment.CommentableId
		//dataResponse["userId"] = comment.UserId
		//
		//appCtx.GetHub().EmitToRoom("lesson-detail-"+string(rune(comment.CommentableId)), websocket.MessageWs{
		//	Channel: "receive-comment",
		//	RoomId:  "lesson-detail-" + string(rune(comment.CommentableId)),
		//	Data:    dataResponse,
		//})

	}
}
