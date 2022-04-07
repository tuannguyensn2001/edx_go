package authtransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	authbusiness "edx_go/module/auth/business"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMe(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		userId, _ := ctx.Get("user_id")

		service := authbusiness.NewAuthService(appCtx.GetDBConnection())

		user, err := service.GetMe(ctx.Request.Context(), userId.(int))

		if err != nil {
			panic(common.ErrDB(err))
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Lấy thông tin người dùng thành công", user))

	}
}
