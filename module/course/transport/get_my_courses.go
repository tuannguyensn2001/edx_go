package coursetransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	authbusiness "edx_go/module/course/business"
	coursestorage "edx_go/module/course/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMyCourses(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user_id, _ := ctx.Get("user_id")

		userId := user_id.(int)

		store := coursestorage.NewSQLStore(appCtx.GetDBConnection())

		biz := authbusiness.NewGetMyCoursesBiz(store)

		result, err := biz.FindByUserId(ctx.Request.Context(), userId)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Lấy danh sách khóa học thành công", result))

	}
}
