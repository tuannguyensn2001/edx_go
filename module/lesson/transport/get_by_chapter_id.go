package lessontransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	lessonbusiness "edx_go/module/lesson/business"
	lessonstorage "edx_go/module/lesson/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LessonGetByChapterId(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		chapterId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := lessonstorage.NewSQLStore(appCtx.GetDBConnection())

		biz := lessonbusiness.NewLessonGetByChapterIdBiz(store)

		lessons, err := biz.GetByChapterIdBiz(ctx.Request.Context(), chapterId)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Lấy danh sách bài học thành công", lessons))

	}
}
