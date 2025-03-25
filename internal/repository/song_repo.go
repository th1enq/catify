package repository

import (
	"catify/internal/models"

	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{
		db: db,
	}
}

func (r *SongRepository) Create(song *models.Song) error {
	return r.db.Save(song).Error
}

func (r *SongRepository) Delete(id uint) error {
	return r.db.Delete(&models.Song{}, id).Error
}

func (r *SongRepository) Update(song *models.Song) error {
	return r.db.Save(song).Error
}

func (r *SongRepository) FindByID(id uint) (*models.Song, error) {
	var song models.Song
	if err := r.db.Where("id = ?", id).First(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *SongRepository) FindByArtist(name string) (*models.Song, error) {

}
