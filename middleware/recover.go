package middleware

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	"github.com/gin-gonic/gin"
)

func Recover(appCtx app_ctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}

				appErr := common.ErrInternal(err.(error))
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)

				return
			}
		}()

		context.Next()
	}
}
