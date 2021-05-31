package answer

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
)

type AnswerFormat struct {
	ID         uint64 `json:"id"`
	Content    string `json:"content"`
	UserID     uint32 `json:"user_id"`
	QuestionID uint64 `json:"question_id"`
}

func FormattingAnswer(answer entity.Answers) AnswerFormat {
	answerFormat := AnswerFormat{
		ID:         answer.ID,
		Content:    answer.Content,
		UserID:     answer.UserID,
		QuestionID: answer.QuestionID,
	}

	return answerFormat
}
