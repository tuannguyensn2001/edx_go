package authtransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	authmodel "edx_go/module/auth/model"
	hashprovider "edx_go/providers/hash"
	tokenprovider "edx_go/providers/token"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var loginDTO authmodel.LoginDTO

		if err := ctx.ShouldBind(&loginDTO); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var user authmodel.User

		db := appCtx.GetDBConnection()

		err := db.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", loginDTO.Email).Scan(&user).Error

		if err != nil {
			panic(common.ErrDB(err))
		}

		check := hashprovider.Compare(loginDTO.Password, user.Password)

		if !check {
			panic(common.ErrInvalidRequest(errors.New("Email hoặc mật khẩu không hợp lệ")))
		}

		accessToken, err := tokenprovider.GenerateToken(user.ID, 3600)

		if err != nil {
			panic(common.ErrInternal(err))
		}

		refreshToken, err := tokenprovider.GenerateToken(user.ID, 36000)

		if err != nil {
			panic(common.ErrInternal(err))
		}

		response := authmodel.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			User:         user,
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Đăng nhập thành công", response))

	}
}
