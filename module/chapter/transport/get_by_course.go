package chaptertransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	"edx_go/module/chapter/business"
	chapterstorage "edx_go/module/chapter/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetChapterByCourse(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		courseId, err := strconv.Atoi(ctx.Param("id"))
		userId, _ := ctx.Get("user_id")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := chapterstorage.NewSQLStore(appCtx.GetDBConnection())

		biz := business.NewChapterGetByCourseAndUserBiz(store)

		result, err := biz.GetByCourseAndUser(ctx.Request.Context(), userId.(int), courseId)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Lấy danh sách khóa học thành công", result))
	}
}
