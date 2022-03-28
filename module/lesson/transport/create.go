package lessontransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	lessonbusiness "edx_go/module/lesson/business"
	lessonmodel "edx_go/module/lesson/model"
	lessonstorage "edx_go/module/lesson/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LessonCreate(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var lesson lessonmodel.LessonCreate
		if err := ctx.ShouldBind(&lesson); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := lessonstorage.NewSQLStore(appCtx.GetDBConnection())

		biz := lessonbusiness.NewLessonCreateBiz(store)

		result, err := biz.LessonCreate(ctx.Request.Context(), &lesson)

		if err != nil {
			panic(common.ErrCannotCreateEntity("Lesson", err))
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Thêm mới bài học thành công", result))
	}
}
