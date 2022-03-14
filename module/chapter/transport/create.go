package chaptertransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	"edx_go/module/chapter/business"
	chaptermodel "edx_go/module/chapter/model"
	chapterstorage "edx_go/module/chapter/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateChapter(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var chapter chaptermodel.ChapterCreate

		if err := ctx.ShouldBind(&chapter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := chapterstorage.NewSQLStore(appCtx.GetDBConnection())

		biz := business.NewChapterCreateBiz(store)

		result, err := biz.CreateChapter(ctx, &chapter)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Thêm mới chương học thành công", result))
	}
}
