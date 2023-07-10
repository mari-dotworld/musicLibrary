package database

import (
	"fmt"
	"os"

	"github.com/musicLibrary/model"
	"github.com/musicLibrary/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() error {
	var err error
	host := os.Getenv("HOST")
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASSWORD")
	db := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	connDetails := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", host, user, pass, db, port)
	server.DbConnection, err = gorm.Open(postgres.Open(connDetails), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if os.Getenv("MIGRATION") != "" && os.Getenv("MIGRATION") == "true"{
        fmt.Println("Migration started")
		server.DbConnection.AutoMigrate(&model.Artist{},&model.MusicAlbum{},&model.Song{})
		fmt.Println("Migration completed")
	}
	return err
}
