package coursetransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	"edx_go/models"
	chaptermodel "edx_go/module/chapter/model"
	coursemodel "edx_go/module/course/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CourseGetByLesson(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		lessonId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()

		var course coursemodel.Course

		db.Raw("SELECT courses.id FROM chapters LEFT JOIN lessons ON lessons.chapter_id = chapters.id LEFT JOIN courses ON courses.id = chapters.course_id WHERE lessons.id = ?", lessonId).Scan(&course)

		var chapters []chaptermodel.Chapter

		db.Raw("SELECT * FROM chapters WHERE course_id = ?", course.ID).Scan(&chapters)

		course.Chapters = chapters

		for index, chapter := range chapters {
			var lessons []models.Lesson
			db.Raw("SELECT * FROM lessons WHERE chapter_id = ?", chapter.ID).Scan(&lessons)

			if lessons == nil {
				chapters[index].Lessons = []models.Lesson{}
			} else {
				chapters[index].Lessons = lessons
			}
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Lấy data thành công", course))

	}
}
