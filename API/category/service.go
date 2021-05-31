package category

import (
	"errors"
	"fmt"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
)

type CategoryService interface {
	GetAllCategories() ([]InputCategoryFormat, error)
	SaveNewCategory(category entity.CategoryInput) (InputCategoryFormat, error)
	FindCategoryByName(categoryName string) (entity.Categories, error)
	UpdateCategoryByName(categoryID string, dataInput entity.UpdateCategoryInput) (InputCategoryFormat, error)
}

type categoryService struct {
	repository CategoryRepository
}

func NewService(repository CategoryRepository) *categoryService {
	return &categoryService{repository}
}

func (s *categoryService) GetAllCategories() ([]InputCategoryFormat, error) {
	categories, err := s.repository.GetAll()

	var categoriesFormat []InputCategoryFormat

	for _, category := range categories {
		var categoryFormat = FormattingInputCategory(category)
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
		Description:  category.Description,
	}

	createCategory, err := s.repository.NewCategory(newCategory)
	formatCategory := FormattingInputCategory(createCategory)

	if err != nil {
		return formatCategory, err
	}

	return formatCategory, nil
}

func (s *categoryService) FindCategoryByName(categoryName string) (entity.Categories, error) {
	category, err := s.repository.FindCategoryName(categoryName)

	if err != nil {
		return entity.Categories{}, err
	}

	if category.CategoryName == "" {
		newError := fmt.Sprintf("category %s is not found", categoryName)
		return entity.Categories{}, errors.New(newError)
	}

	// categoryFormat := FormattingCategory(category)

	return category, nil
}

func (s *categoryService) UpdateCategoryByName(categoryName string, dataInput entity.UpdateCategoryInput) (InputCategoryFormat, error) {
	var dataUpdate = map[string]interface{}{}

	category, err := s.repository.FindCategoryName(categoryName)

	if err != nil {
		return InputCategoryFormat{}, err
	}

	if category.ID == 0 {
		newError := fmt.Sprintf("category name %s not found", categoryName)
		return InputCategoryFormat{}, errors.New(newError)
	}

	if dataInput.CategoryName != "" || len(dataInput.CategoryName) != 0 {
		dataUpdate["category_name"] = dataInput.CategoryName
	}
	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}

	catUpdated, err := s.repository.UpdateByID(categoryName, dataUpdate)
	formatCatUpdated := FormattingInputCategory(catUpdated)

	if err != nil {
		return formatCatUpdated, err
	}

	return formatCatUpdated, nil
}
