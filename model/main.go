package model

import (
	"time"
)

type Artist struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Songs     []Song `gorm:"many2many:artist_song"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type MusicAlbum struct {
	ID        uint `gorm:"primaryKey"`
	AlbumName string
	Url       string
	IsDeleted bool `gorm:"default:false;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type Song struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	MusicAlbum MusicAlbum `gorm:"references:id;foreignKey:album_id;"`
	AlbumID    uint
	Artists    []Artist  `gorm:"many2many:artist_song"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}
