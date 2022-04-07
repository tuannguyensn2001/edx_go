package carttransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	cartservice "edx_go/module/cart/service"
	cartstruct "edx_go/module/cart/struct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddToCart(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var addToCartInput cartstruct.AddToCartInput

		if err := ctx.ShouldBind(&addToCartInput); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userId, _ := ctx.Get("user_id")

		service := cartservice.NewCartService(appCtx.GetDBConnection())

		result, err := service.AddToCart(ctx, userId.(int), addToCartInput.CourseId)

		if err != nil {
			switch err.(type) {
			case *common.AppError:
				panic(err)
			default:
				panic(common.ErrInternal(err))
			}
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Thêm vào giỏ hàng thành công", result))

	}
}
