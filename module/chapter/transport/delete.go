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

func ChapterDelete(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		chapterId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := chapterstorage.NewSQLStore(appCtx.GetDBConnection())

		biz := business.NewChapterDeleteBiz(store)

		if err := biz.Delete(ctx.Request.Context(), chapterId); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Xóa thành công", chapterId))
	}
}
