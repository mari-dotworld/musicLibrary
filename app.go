package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/musicLibrary/database"
	"github.com/musicLibrary/handler"
)

func appStartup() {
	//Load env variables
	godotenv.Load()
	fmt.Println("connecting postgres database")
	err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("postgres database connected")
}

func main() {
	appStartup()
	app := gin.New()

	//Album CRUD
	app.GET("/album/:id",handler.GetAlbum)
	app.POST("/album",handler.CreateAlbum)
	app.PUT("/album/:id",func(ctx *gin.Context) {

	})
	app.DELETE("/album/:id",func(ctx *gin.Context) {

	})
	app.Run("localhost:8080")
}
