package entity

type QuestionInput struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint32 `json:"category_id" binding:"required"`
	UserID     uint32 `json:"user_id"`
}

type UpdateQuestionInput struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID uint32 `json:"category_id"`
}
