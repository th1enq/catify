package database

import (
	"catify/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgres(config struct {
	Port     string
	Host     string
	User     string
	Password string
	Name     string
}) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s, port=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Song{},
		&models.Playlist{},
		&models.PlaylistSong{},
		&models.User{},
		&models.UserFollow{},
		&models.UserFavorite{},
		&models.Activity{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}
