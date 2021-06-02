package category

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]entity.Categories, error)
	NewCategory(category entity.Categories) (entity.Categories, error)
	FindCategoryName(categoryName string) (entity.Categories, error)
	UpdateByID(category_name string, dataUpdate map[string]interface{}) (entity.Categories, error)
	GetAllQuestionsByCategory(categoryID int) ([]entity.Questions, error)
	DeleteCategory(categoryName string) (string, error)
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

	if err := r.db.Where("category_name = ?", categoryName).Preload("Questions").Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *Repository) UpdateByID(categoryName string, dataUpdate map[string]interface{}) (entity.Categories, error) {
	var category entity.Categories

	if err := r.db.Model(&category).Where("category_name = ?", categoryName).Updates(dataUpdate).Error; err != nil {
		return category, err
	}

	if err := r.db.Where("category_name = ?", categoryName).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *Repository) GetAllQuestionsByCategory(categoryID int) ([]entity.Questions, error) {
	var Questions []entity.Questions

	err := r.db.Where("category_id = ?", categoryID).Preload("User").Preload("Category").Preload("Answers").Find(&Questions).Error
	if err != nil {
		return Questions, err
	}

	return Questions, nil
}

func (r *Repository) DeleteCategory(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.Categories{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
