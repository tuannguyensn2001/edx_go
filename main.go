package main

import (
	app_ctx "edx_go/component"
	"edx_go/middleware"
	"edx_go/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/edx?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("DB connect failed", err)
	}

	appCtx := app_ctx.NewAppContext(db)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	routes.DeclareRoute(r, appCtx)

	err = r.Run(":5000")
	if err != nil {
		return
	}
}
