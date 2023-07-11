package musicalbum

import (
	"github.com/musicLibrary/controller/song"
)


type GetMusicAlbum struct { 
	Name  string `json:"albumName"`
	Url   string `json:"url"`
	SongList []song.Song `json:"songList"`
}

type CreateMusicAlbum struct { 
	Name  string `json:"albumName"`
	SongList []song.CreateSong `json:"songList"`
}