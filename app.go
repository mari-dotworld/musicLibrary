package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/musicLibrary/database"
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
	app.Run("localhost:8080")
}
