package artist

import (
	"github.com/musicLibrary/model"
	"github.com/musicLibrary/server"
	"gorm.io/gorm"
)

func CreateArtist(artistData *model.Artist)(*gorm.DB,error){
	tx := server.DbConnection.Model(&model.Artist{}).Create(&model.Artist{
		Name: artistData.Name,
	})
	if tx.Error != nil {
		tx.Rollback()
		return tx,tx.Error
	}
	return tx,nil
}