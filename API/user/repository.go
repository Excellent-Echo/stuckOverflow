package user

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	GetAll() ([]entity.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateUser(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetAll() ([]entity.User, error) {
	var Users []entity.User

	err := r.db.Preload("User").Find(&Users).Error
	if err != nil {
		return Users, err
	}

	return Users, nil
}
