package middleware

import (
	app_ctx "edx_go/component"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(appCtx app_ctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("user_id", 1)
		context.Next()
	}
}
