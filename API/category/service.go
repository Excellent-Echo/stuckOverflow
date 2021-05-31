package category

import (
	"errors"
	"fmt"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
)

type CategoryService interface {
	GetAllCategories() ([]CategoryFormat, error)
	SaveNewCategory(category entity.CategoryInput) (InputCategoryFormat, error)
	FindCategoryByName(categoryName string) (CategoryFormat, error)
}

type categoryService struct {
	repository CategoryRepository
}

func NewService(repository CategoryRepository) *categoryService {
	return &categoryService{repository}
}

func (s *categoryService) GetAllCategories() ([]CategoryFormat, error) {
	categories, err := s.repository.GetAll()

	var categoriesFormat []CategoryFormat

	for _, category := range categories {
		var categoryFormat = FormattingCategory(category)
		categoriesFormat = append(categoriesFormat, categoryFormat)
	}

	if err != nil {
		return categoriesFormat, err
	}

	return categoriesFormat, nil
}

func (s *categoryService) SaveNewCategory(category entity.CategoryInput) (InputCategoryFormat, error) {
	var newCategory = entity.Categories{
		CategoryName: category.CategoryName,
	}

	createCategory, err := s.repository.NewCategory(newCategory)
	formatCategory := FormattingInputCategory(createCategory)

	if err != nil {
		return formatCategory, err
	}

	return formatCategory, nil
}

func (s *categoryService) FindCategoryByName(categoryName string) (CategoryFormat, error) {
	category, err := s.repository.FindCategoryName(categoryName)

	if err != nil {
		return CategoryFormat{}, err
	}

	if category.CategoryName == "" {
		newError := fmt.Sprintf("category %s is not found", categoryName)
		return CategoryFormat{}, errors.New(newError)
	}

	categoryFormat := FormattingCategory(category)

	return categoryFormat, nil
}
