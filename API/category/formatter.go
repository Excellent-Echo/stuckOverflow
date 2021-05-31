package category

import "github.com/Excellent-Echo/stuckOverflow/API/API/entity"

type CategoryFormat struct {
	ID           uint64             `json:"id"`
	CategoryName string             `json:"category_name"`
	Questions    []entity.Questions `json:"questions"`
}

type InputCategoryFormat struct {
	ID           uint64 `json:"id"`
	CategoryName string `json:"category_name"`
}

func FormattingCategory(category entity.Categories) CategoryFormat {
	categoryFormat := CategoryFormat{
		ID:           uint64(category.ID),
		CategoryName: category.CategoryName,
		Questions:    category.Questions,
	}

	return categoryFormat
}

func FormattingInputCategory(category entity.Categories) InputCategoryFormat {
	categoryFormat := InputCategoryFormat{
		ID:           uint64(category.ID),
		CategoryName: category.CategoryName,
	}

	return categoryFormat
}
