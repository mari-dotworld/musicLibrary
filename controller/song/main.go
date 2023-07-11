package song

import (
	"github.com/musicLibrary/model"
	"github.com/musicLibrary/server"
)

func GetSongsByAlbumID(albumID uint)([]Song,error){
	var song []Song
	tx := server.DbConnection.Model(&model.Song{}).Find(&song).Where("album_id=?",albumID)
	return song,tx.Error
}
