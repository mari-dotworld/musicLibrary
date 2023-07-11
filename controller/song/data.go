package song

import "github.com/musicLibrary/model"

type CreateSong struct {
	Name string `json:"name"`
	SongId uint
	ArtistList []uint `json:"artistList"`
}

type Song struct{
	Name string `json:"name"`
	SongId uint
	ArtistList []model.Artist `json:"artistList"`
}