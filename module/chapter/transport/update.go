package chaptertransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	"edx_go/module/chapter/business"
	chaptermodel "edx_go/module/chapter/model"
	chapterstorage "edx_go/module/chapter/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ChapterUpdate(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var chapter chaptermodel.ChapterUpdate

		chapterId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := ctx.ShouldBind(&chapter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := chapterstorage.NewSQLStore(appCtx.GetDBConnection())

		biz := business.NewChapterUpdateBiz(store)

		if err := biz.Update(ctx.Request.Context(), chapterId, chapter); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Cập nhật chương học thành công", gin.H{
			"id": chapterId,
		}))

	}
}
