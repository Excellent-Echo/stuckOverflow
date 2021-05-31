package entity

type CategoryInput struct {
	CategoryName string `json:"category_name" binding:"required"`
	Description  string `json:"description"`
}

type UpdateCategoryInput struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}
