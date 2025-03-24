package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	Name        string         `json:"name" gorm:"index"`
	Description string         `json:"description"`
	IsPublic    bool           `json:"isPublic" gorm:"default:false"`
	CoverImage  string         `json:"coverImage"`
	UserID      int            `json:"userID"`
	Users       User           `json:"-" gorm:"foreignKey:UserID"`
	Songs       []PlaylistSong `json:"songs"`
}

type PlaylistSong struct {
	gorm.Model
	PlaylistID int  `json:"playlistID" gorm:"index"`
	SongID     uint `json:"songID" gorm:"index"`
	Position   int  `json:"position"`
	AddedBy    uint `json:"addedBy"`
}

func (Playlist) TableName() string {
	return "playlist"
}

func (PlaylistSong) TableName() string {
	return "playlist_songs"
}
