package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title      string    `json:"title" gorm:"index"`
	Artist     string    `json:"artist" gorm:"index"`
	Album      string    `json:"album" gorm:"index"`
	FilePath   string    `json:"filePath"`
	Duration   int       `json:"duration"`
	FileSize   int64     `json:"fileSize"`
	Genre      string    `json:"genre" gorm:"index"`
	Year       int       `json:"year"`
	TrackNum   int       `json:"trackNum"`
	PlayCount  int       `json:"playCount" gorm:"default:0"`
	LastPlayed time.Time `json:"lastPlayed"`
	UploadedBy uint      `json:"uploadedBy" gorm:"index"`

	FavoritedBy []User `json:"-" gorm:"many2many:user_favorites;"`
}

func (Song) TableName() string {
	return "songs"
}
