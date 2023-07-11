package musicalbum

import (
	"github.com/musicLibrary/controller/song"
	"github.com/musicLibrary/helpers"
	"github.com/musicLibrary/model"
	"github.com/musicLibrary/server"
)


func GetAlbum(albumId uint)(GetMusicAlbum,error){
	var err error
	var albumData GetMusicAlbum
	reqFields := "music_album.album_name,music_album.url,song.id,song.name"
	rows,err := server.DbConnection.Table("music_album").Select(reqFields).Joins("join song on song.album_id=music_album.id").Where("music_album.id=? and is_deleted=false",albumId).Rows()
	if err != nil {
		return albumData,err
	}

	for rows.Next(){
		var row song.Song
		var tSong model.Song

		err = rows.Scan(&albumData.Name,&albumData.Url,&row.SongId,&row.Name)
		if err != nil {
			return albumData,err
     	}
        
		err = server.DbConnection.Preload("Artists").Find(&tSong).Where("song_id=?",row.SongId).Error
		if err != nil {
			return albumData,err
     	}
		row.ArtistList = append(row.ArtistList, tSong.Artists...)
		albumData.SongList = append(albumData.SongList, row)
	}

	return albumData,err
}

func CreateAlbum(albumData CreateMusicAlbum)(uint,error){
	var err error
	var musicAlbumData model.MusicAlbum
	musicAlbumData.AlbumName = albumData.Name
	musicAlbumData.Url       = helpers.RandomString(12)

	tx := server.DbConnection.Model(&model.MusicAlbum{}).Create(&musicAlbumData)

	if tx.Error != nil {
		tx.Rollback()
		return 0,tx.Error
	}

	for _,curr := range albumData.SongList {
		cSong := model.Song{
			AlbumID: musicAlbumData.ID,
			Name: curr.Name,
		}

		for _,cArtist := range curr.ArtistList {
			var tempArtist model.Artist
			tempArtist.ID = cArtist
			cSong.Artists = append(cSong.Artists, tempArtist)
		}

		ttx := server.DbConnection.Create(&cSong)
		if ttx.Error != nil {
			ttx.Rollback()
			return 0,ttx.Error
		}
		server.DbConnection.Save(&cSong)
	}

	return musicAlbumData.ID,err
}

func DeleteAlbum(albumId uint)error{
	err := server.DbConnection.Model(&model.MusicAlbum{}).Where("id=?",albumId).Update("is_deleted",true).Error
	return err
}