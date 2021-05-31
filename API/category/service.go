package category

import (
	"errors"
	"fmt"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/question"
)

type CategoryService interface {
	GetAllCategories() ([]InputCategoryFormat, error)
	SaveNewCategory(category entity.CategoryInput) (InputCategoryFormat, error)
	FindCategoryByName(categoryName string) (CategoryFormat, error)
	UpdateCategoryByName(categoryID string, dataInput entity.UpdateCategoryInput) (InputCategoryFormat, error)
	GetAllQuestionsByCategory(categoryName string) ([]question.QuestionFormat, error)
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

func (s *categoryService) GetAllQuestionsByCategory(categoryName string) ([]question.QuestionFormat, error) {
	category, _ := s.repository.FindCategoryName(categoryName)

	categoryID := int(category.ID)

	questions, err := s.repository.GetAllQuestionsByCategory(categoryID)

	var questionsFormat []question.QuestionFormat

	for _, q := range questions {
		var questionFormat = question.FormattingQuestion(q)
		questionsFormat = append(questionsFormat, questionFormat)
	}

	if err != nil {
		return questionsFormat, err
	}

	return questionsFormat, nil
}
