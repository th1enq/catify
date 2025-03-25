package repository

import (
	"catify/internal/models"

	"gorm.io/gorm"
)

type PlaylistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{
		db: db,
	}
}

func (r *PlaylistRepository) Create(playlist models.Playlist) error {
	
}
