package repository

import (
	"catify/internal/models"

	"gorm.io/gorm"
)

type FollowRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{
		db: db,
	}
}

func (r *FollowRepository) Create(u *models.UserFollow) error {
	return r.db.Save(u).Error
}

func (r *FollowRepository) Delete(u *models.UserFollow) error {
	return r.db.Delete(u).Error
}

func (r *FollowRepository) IsFollwing(FollowerID, FollowingID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&models.UserFollow{}).Where("follower_id = ? AND following_id = ?", FollowerID, FollowingID).Count(&count).Error; err != nil {
		return false, nil
	}
	return count > 0, nil
}

func (r *FollowRepository) GetFollowersID(id uint) ([]uint, error) {
	var users []models.UserFollow
	if err := r.db.Where("following_id = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}

	var result = make([]uint, len(users))
	for i, followers := range users {
		result[i] = followers.FollowerID
	}

	return result, nil
}

func (r *FollowRepository) GetFollowingsID(id uint) ([]uint, error) {
	var users []models.UserFollow
	if err := r.db.Where("follower_id = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}

	var result = make([]uint, len(users))
	for i, followings := range users {
		result[i] = followings.FollowingID
	}

	return result, nil
}
