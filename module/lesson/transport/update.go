package lessontransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	lessonmodel "edx_go/module/lesson/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LessonUpdate(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		chapterId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var lesson lessonmodel.LessonEdit

		if err := ctx.ShouldBind(&lesson); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()

		if err := db.Exec("UPDATE lessons SET name = ?, video_url = ? WHERE id = ?", lesson.Name, lesson.VideoURL, chapterId).Error; err != nil {
			panic(common.ErrCannotUpdateEntity("lesson", err))
		}

		lesson.Id = chapterId
		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Cập nhật chương học thành công", lesson))

	}
}
