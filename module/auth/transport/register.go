package authtransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	authbusiness "edx_go/module/auth/business"
	authmodel "edx_go/module/auth/model"
	authstorage "edx_go/module/auth/storage"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Register(appContext app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var userRegister authmodel.UserRegister

		if err := ctx.ShouldBind(&userRegister); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		validate := validator.New()

		errValidate := validate.Struct(userRegister)

		if errValidate != nil {
			panic(common.ErrInvalidRequest(errValidate))
		}

		store := authstorage.NewSQLStore(appContext.GetDBConnection())

		biz := authbusiness.NewRegisterUserBiz(store)

		result, err := biz.Register(ctx.Request.Context(), &userRegister)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Thêm mới người dùng thành công", result))
	}
}
