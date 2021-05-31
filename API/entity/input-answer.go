package entity

type AnswerInput struct {
	Content    string `json:"content" binding:"required"`
	UserID     uint32 `json:"user_id"`
	QuestionID uint64 `json:"question_id"`
}
