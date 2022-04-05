package lessontransport

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	"edx_go/models"
	lessonmodel "edx_go/module/lesson/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func UpdateOrder(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateOrderDTO lessonmodel.UpdateOrderDTO

		if err := ctx.ShouldBind(&updateOrderDTO); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()

		err := db.Transaction(func(tx *gorm.DB) error {
			for index, item := range updateOrderDTO.Ids {
				err := tx.Model(&models.Lesson{}).Where(" id = ? ", item).Update("order", index+1).Error
				if err != nil {
					return err
				}
			}
			return nil
		})

		if err != nil {
			panic(common.ErrDB(err))
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse("Cập nhật thứ tự thành công", nil))

	}
}
