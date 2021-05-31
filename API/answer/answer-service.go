package answer

import (
	"errors"
	"fmt"
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/Excellent-Echo/stuckOverflow/API/API/user"
)

type AnswerService interface {
	PostNewAnswer(questionID int, userID int, answer entity.AnswerInput) (entity.Answers, error)
	UpdateAnswerByID(id string, dataInput entity.AnswerInput) (AnswerFormat, error)
	FindAnswerByID(id string) (AnswerFormat, error)
	DeleteAnswerByID(id string) (interface{}, error)
}

type answerService struct {
	repository AnswerRepository
}

func NewService(repository AnswerRepository) *answerService {
	return &answerService{repository}
}

func (s *answerService) PostNewAnswer(questionID int, userID int, answer entity.AnswerInput) (entity.Answers, error) {
	var newAnswer = entity.Answers{
		Content:    answer.Content,
		UserID:     uint32(userID),
		QuestionID: uint64(questionID),
		CreatedAt:  time.Now(),
	}

	createAnswer, err := s.repository.PostAnswer(newAnswer)

	if err != nil {
		return createAnswer, err
	}

	return createAnswer, nil
}

func (s *answerService) UpdateAnswerByID(id string, dataInput entity.AnswerInput) (AnswerFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return AnswerFormat{}, err
	}

	answer, err := s.repository.FindAnswerByID(id)

	if err != nil {
		return AnswerFormat{}, err
	}

	if answer.ID == 0 {
		newError := fmt.Sprintf("answer id %s is not found", id)
		return AnswerFormat{}, errors.New(newError)
	}

	if dataInput.Content != "" || len(dataInput.Content) != 0 {
		dataUpdate["content"] = dataInput.Content
	}

	dataUpdate["updated_at"] = time.Now()

	answerUpdate, err := s.repository.UpdateAnswer(id, dataUpdate)

	if err != nil {
		return AnswerFormat{}, err
	}

	formatAnswer := FormattingAnswer(answerUpdate)

	return formatAnswer, nil
}

func (s *answerService) FindAnswerByID(id string) (AnswerFormat, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return AnswerFormat{}, err
	}

	answer, err := s.repository.FindAnswerByID(id)

	if err != nil {
		return AnswerFormat{}, err
	}

	if answer.ID == 0 {
		newError := fmt.Sprintf("answer id %s is not found", id)
		return AnswerFormat{}, errors.New(newError)
	}

	answerFormat := FormattingAnswer(answer)

	return answerFormat, nil

}

func (s *answerService) DeleteAnswerByID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	answer, err := s.repository.FindAnswerByID(id)

	if err != nil {
		return nil, err
	}

	if answer.ID == 0 {
		newError := fmt.Sprintf("answer id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteAnswer(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("delete answer id %s succeed", id)

	formatDelete := user.FormatDelete(msg)

	return formatDelete, nil
}
