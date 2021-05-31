package question

import "github.com/Excellent-Echo/stuckOverflow/API/API/entity"

type QuestionFormat struct {
	ID           uint64           `json:"id"`
	Title        string           `json:"title"`
	Content      string           `json:"content"`
	UserID       uint32           `json:"user_id"`
	UserName     string           `json:"user_name"`
	CategoryID   uint32           `json:"category_id"`
	CategoryName string           `json:"category_name"`
	Answers      []entity.Answers `json:"answers"`
}

type InputQuestionFormat struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserID     uint32 `json:"user_id"`
	CategoryID uint32 `json:"category_id"`
}

func FormattingQuestion(question entity.Questions) QuestionFormat {
	questionFormat := QuestionFormat{
		ID:           question.ID,
		Title:        question.Title,
		Content:      question.Content,
		UserID:       question.UserID,
		UserName:     question.User.UserName,
		CategoryID:   question.CategoryID,
		CategoryName: question.Category.CategoryName,
		Answers:      question.Answers,
	}

	return questionFormat
}

func FormattingInputQuestion(question entity.Questions) InputQuestionFormat {
	questionFormat := InputQuestionFormat{
		ID:         question.ID,
		Title:      question.Title,
		Content:    question.Content,
		UserID:     question.UserID,
		CategoryID: question.CategoryID,
	}

	return questionFormat
}
