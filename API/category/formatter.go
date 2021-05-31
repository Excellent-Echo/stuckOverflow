package category

import "github.com/Excellent-Echo/stuckOverflow/API/API/entity"

type CategoryFormat struct {
	ID           uint64             `json:"id"`
	CategoryName string             `json:"category_name"`
	Description  string             `json:"description"`
	Questions    []entity.Questions `json:"questions"`
}

type QuestionsList struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageFile string `json:"image_file"`
	UserID    uint32 `json:"user_id"`
}

type InputCategoryFormat struct {
	ID           uint64 `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

func FormattingCategory(category entity.Categories) CategoryFormat {
	categoryFormat := CategoryFormat{
		ID:           uint64(category.ID),
		CategoryName: category.CategoryName,
		Description:  category.Description,
	}

	return categoryFormat
}

func FormattingInputCategory(category entity.Categories) InputCategoryFormat {
	categoryFormat := InputCategoryFormat{
		ID:           uint64(category.ID),
		CategoryName: category.CategoryName,
		Description:  category.Description,
	}

	return categoryFormat
}
