package repositories

import (
	"be-journey/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Journey").Preload("Journey.User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Journey").Preload("Journey.User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Preload("Journey").Preload("Journey.User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Preload("Journey").Preload("Journey.User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Preload("Journey").Preload("Journey.User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Delete(&user).Error

	return user, err
}
