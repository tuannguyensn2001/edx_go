package routes

import (
	app_ctx "edx_go/component"
	"edx_go/middleware"
	authtransport "edx_go/module/auth/transport"
	chaptertransport "edx_go/module/chapter/transport"
	coursetransport "edx_go/module/course/transport"
	"github.com/gin-gonic/gin"
)

func DeclareRoute(r *gin.Engine, appCtx app_ctx.AppContext) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authtransport.Register(appCtx))

		v1.GET("/courses", middleware.AuthMiddleware(appCtx), coursetransport.GetMyCourses(appCtx))
		v1.POST("/courses", middleware.AuthMiddleware(appCtx), coursetransport.CreateCourse(appCtx))

		v1.POST("chapters", chaptertransport.CreateChapter(appCtx))
	}
}
