package models

import (
	"time"

	"gorm.io/gorm"
)

type Music struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:255;not null"`
	Artist      string         `json:"artist" gorm:"size:255"`
	Album       string         `json:"album" gorm:"size:255"`
	Genre       string         `json:"genre" gorm:"size:100"`
	FileName    string         `json:"file_name" gorm:"size:255;not null"`
	ContentType string         `json:"content_type" gorm:"size:100"`
	FileData    []byte         `json:"-" gorm:"type:bytea"`
	FileSize    int64          `json:"file_size"`
	Duration    float64        `json:"duration"`
	Description string         `json:"description" gorm:"type:text"`
	CreateAt    time.Time      `json:"create_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Music) TableName() string {
	return "musics"
}
