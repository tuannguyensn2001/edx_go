package lessontransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	"edx_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetById(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var lesson models.Lesson

		db := appCtx.GetDBConnection()

		err = db.Raw("SELECT * FROM lessons WHERE id = ?", id).Scan(&lesson).Error

		if err != nil {
			panic(common.ErrDB(err))
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Lấy thông tin bài học thành công", lesson))

	}
}
