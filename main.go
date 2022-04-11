package main

import (
	app_ctx "edx_go/component"
	"edx_go/middleware"
	"edx_go/routes"
	"edx_go/websocket"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	godotenv.Load()

	dsn := os.Getenv("DATABASE")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("DB connect failed", err)
	}
	//
	hub := websocket.NewHub()
	appCtx := app_ctx.NewAppContext(db, hub)
	//
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.Use(middleware.Recover(appCtx))

	routes.DeclareRoute(r, appCtx)

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
