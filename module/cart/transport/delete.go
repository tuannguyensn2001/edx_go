package carttransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	cartservice "edx_go/module/cart/service"
	cartstruct "edx_go/module/cart/struct"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DeleteInCart(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		log.Println(ctx.PostForm("course_id"))

		var deleteInCartInput cartstruct.AddToCartInput

		if err := ctx.ShouldBind(&deleteInCartInput); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userId, _ := ctx.Get("user_id")

		service := cartservice.NewCartService(appCtx.GetDBConnection())

		err := service.Delete(ctx.Request.Context(), userId.(int), deleteInCartInput.CourseId)

		if err != nil {
			panic(common.ErrDB(err))
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Xóa khỏi giỏ hàng thành công", nil))

	}
}
