package category

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]entity.Categories, error)
	NewCategory(category entity.Categories) (entity.Categories, error)
	FindCategoryName(categoryName string) (entity.Categories, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]entity.Categories, error) {
	var Category []entity.Categories

	err := r.db.Find(&Category).Error
	if err != nil {
		return Category, err
	}

	return Category, nil
}

func (r *Repository) NewCategory(category entity.Categories) (entity.Categories, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *Repository) FindCategoryName(categoryName string) (entity.Categories, error) {
	var category entity.Categories

	if err := r.db.Where("category_name = ?", categoryName).Preload("Questions.Answers").Preload("Questions.Category").Preload("Questions").Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}
