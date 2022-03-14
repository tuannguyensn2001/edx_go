package coursetransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	authbusiness "edx_go/module/course/business"
	coursemodel "edx_go/module/course/model"
	coursestorage "edx_go/module/course/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourse(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var course coursemodel.CourseCreate

		if err := ctx.ShouldBind(&course); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userId, _ := ctx.Get("user_id")

		course.UserId = userId.(int)

		store := coursestorage.NewSQLStore(appCtx.GetDBConnection())

		biz := authbusiness.NewCreateCourseBiz(store)

		result, err := biz.CreateCourse(ctx.Request.Context(), &course)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Thêm mới khóa học thành công", result))

	}
}
