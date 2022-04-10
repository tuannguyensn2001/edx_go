package main

import (
	app_ctx "edx_go/component"
	"edx_go/middleware"
	"edx_go/socket"
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

	//dsn := os.Getenv("DATABASE")

	dsn := "admin:java2001@tcp(database-1.ckh8wqhaxqge.us-east-1.rds.amazonaws.com:3306)/edx?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("DB connect failed", err)
	}
	//
	appCtx := app_ctx.NewAppContext(db)
	//
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.Use(middleware.Recover(appCtx))

	r.StaticFile("demo", "./demo.html")

	//routes.DeclareRoute(r, appCtx)

	engine := socket.NewEngine()

	err = engine.Run(appCtx, r)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
