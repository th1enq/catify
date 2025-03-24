package models

import (
	"time"

	"gorm.io/gorm"
)

type UserFollow struct {
	gorm.Model
	FollowerID  uint `json:"followerID" gorm:"index"`
	FollowingID uint `json:"followingID" gorm:"index"`
}

type UserFavorite struct {
	gorm.Model
	UserID  uint      `json:"userID" gorm:"index"`
	SongID  uint      `json:"songID" gorm:"index"`
	AddedAt time.Time `json:"addedAt" gorm:"autoCreateTime"`
}

type Activity struct {
	gorm.Model
	UserID       uint      `json:"userID" gorm:"index"`
	ActivityType string    `json:"activityType" gorm:"index"`
	SongID       *uint     `json:"songID"`
	PlaylistID   *uint     `json:"playlistID"`
	Timestamp    time.Time `json:"timestamp" gorm:"autoCreateTime"`
}

func (UserFollow) TableName() string {
	return "user_follows"
}

func (UserFavorite) TableName() string {
	return "user_favorites"
}

func (Activity) TableName() string {
	return "activities"
}
