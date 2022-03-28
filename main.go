package main

import (
	app_ctx "edx_go/component"
	"edx_go/middleware"
	"edx_go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("env loaded err")
	}

	dsn := os.Getenv("DATABASE")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("DB connect failed", err)
	}

	appCtx := app_ctx.NewAppContext(db)

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.Use(middleware.Recover(appCtx))

	routes.DeclareRoute(r, appCtx)

	err = r.Run(":5000")
	if err != nil {
		return
	}
}
