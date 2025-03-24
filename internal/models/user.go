package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string    `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Email          string    `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password       string    `json:"-" gorm:"size:100;not null"`
	FirstName      string    `json:"firstName" gorm:"size:50"`
	LastName       string    `json:"lastName" gorm:"size:50"`
	ProfilePicture string    `json:"profilePicture"`
	Bio            string    `json:"bio" gorm:"size:500"`
	IsActive       bool      `json:"isActive" gorm:"default:true"`
	LastLogin      time.Time `json:"lastLogin"`
	Role           string    `json:"role" gorm:"default:user"`

	Playlists  []Playlist `json:"-" gorm:"foreignKey:UserID"`
	Favourites []Song     `json:"-" gormn:"many2many:user_favorites;"`
}

func (u *User) CheckPasswords(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))

	return err == nil
}

func (User) TableName() string {
	return "users"
}
