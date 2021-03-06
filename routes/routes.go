package routes

import (
	app_ctx "edx_go/component"
	"edx_go/middleware"
	authtransport "edx_go/module/auth/transport"
	carttransport "edx_go/module/cart/transport"
	chaptertransport "edx_go/module/chapter/transport"
	commenttransport "edx_go/module/comment/transport"
	coursetransport "edx_go/module/course/transport"
	lessontransport "edx_go/module/lesson/transport"
	uploadtransport "edx_go/module/upload/transport"
	"edx_go/websocket"
	"github.com/gin-gonic/gin"
)

func DeclareRoute(r *gin.Engine, appCtx app_ctx.AppContext) {

	r.GET("/ws", func(context *gin.Context) {
		websocket.WsHandler(context.Writer, context.Request, appCtx.GetHub())
	})
	go appCtx.GetHub().Run()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authtransport.Register(appCtx))
		v1.POST("/auth/login", authtransport.Login(appCtx))
		v1.GET("/auth/me", middleware.AuthMiddleware(appCtx), authtransport.GetMe(appCtx))

		v1.GET("/courses", middleware.AuthMiddleware(appCtx), coursetransport.GetMyCourses(appCtx))
		v1.POST("/courses", middleware.AuthMiddleware(appCtx), coursetransport.CreateCourse(appCtx))
		v1.GET("/courses/lesson/:id", coursetransport.CourseGetByLesson(appCtx))

		v1.POST("/chapters", chaptertransport.CreateChapter(appCtx))
		v1.GET("/chapters/course/:id", middleware.AuthMiddleware(appCtx), chaptertransport.GetChapterByCourse(appCtx))
		v1.PUT("/chapters/:id", middleware.AuthMiddleware(appCtx), chaptertransport.ChapterUpdate(appCtx))
		v1.DELETE("/chapters/:id", middleware.AuthMiddleware(appCtx), chaptertransport.ChapterDelete(appCtx))

		v1.POST("/lessons", middleware.AuthMiddleware(appCtx), lessontransport.LessonCreate(appCtx))
		v1.GET("/lessons/chapter/:id", middleware.AuthMiddleware(appCtx), lessontransport.LessonGetByChapterId(appCtx))
		v1.PUT("/lessons/:id", middleware.AuthMiddleware(appCtx), lessontransport.LessonUpdate(appCtx))
		v1.PUT("/lessons/order", middleware.AuthMiddleware(appCtx), lessontransport.UpdateOrder(appCtx))
		v1.GET("/lessons/:id", middleware.AuthMiddleware(appCtx), lessontransport.GetById(appCtx))

		v1.POST("/carts", middleware.AuthMiddleware(appCtx), carttransport.AddToCart(appCtx))
		v1.DELETE("/carts", middleware.AuthMiddleware(appCtx), carttransport.DeleteInCart(appCtx))

		v1.POST("/upload", uploadtransport.Upload(appCtx))

		v1.POST("/comments", middleware.AuthMiddleware(appCtx), commenttransport.CreateComment(appCtx))

	}
}
