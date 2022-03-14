package routes

import (
	app_ctx "edx_go/component"
	authtransport "edx_go/module/auth/transport"
	"github.com/gin-gonic/gin"
)

func DeclareRoute(r *gin.Engine, appCtx app_ctx.AppContext) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authtransport.Register(appCtx))
	}
}
