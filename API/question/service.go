package question

import (
	"errors"
	"fmt"
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/Excellent-Echo/stuckOverflow/API/API/user"
)

type QuestionService interface {
	FindAllQuestions() ([]QuestionFormat, error)
	SaveNewQuestion(question entity.QuestionInput) (InputQuestionFormat, error)
	FindQuestionById(id string) (QuestionFormat, error)
	UpdateQuestionById(id string, dataInput entity.UpdateQuestionInput) (QuestionFormat, error)
	DeleteQuestionById(id string) (interface{}, error)
}

type questionService struct {
	repository QuestionRepository
}

func QuestionNewService(repository QuestionRepository) *questionService {
	return &questionService{repository}
}

func (s *questionService) FindAllQuestions() ([]QuestionFormat, error) {
	questions, err := s.repository.FindAllQuestions()

	var questionsFormat []QuestionFormat

	for _, question := range questions {
		var questionFormat = FormattingQuestion(question)
		questionsFormat = append(questionsFormat, questionFormat)
	}

	if err != nil {
		return questionsFormat, err
	}

	return questionsFormat, nil
}

func (s *questionService) SaveNewQuestion(question entity.QuestionInput) (InputQuestionFormat, error) {
	var newQuestion = entity.Questions{
		Title:      question.Title,
		Content:    question.Content,
		CategoryID: question.CategoryID,
		UserID:     question.UserID,
		CreatedAt:  time.Now(),
	}

	createQuestion, err := s.repository.PostQuestion(newQuestion)
	formatQuestion := FormattingInputQuestion(createQuestion)

	if err != nil {
		return formatQuestion, err
	}

	return formatQuestion, nil
}

func (s *questionService) FindQuestionById(id string) (QuestionFormat, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return QuestionFormat{}, err
	}

	question, err := s.repository.FindQuestionById(id)

	if err != nil {
		return QuestionFormat{}, err
	}

	if question.ID == 0 {
		newError := fmt.Sprintf("question id %s is not found", id)
		return QuestionFormat{}, errors.New(newError)
	}

	questionFormat := FormattingQuestion(question)

	return questionFormat, nil

}

func (s *questionService) UpdateQuestionById(id string, dataInput entity.UpdateQuestionInput) (QuestionFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return QuestionFormat{}, err
	}

	question, err := s.repository.FindQuestionById(id)

	if err != nil {
		return QuestionFormat{}, err
	}

	if question.ID == 0 {
		newError := fmt.Sprintf("question id %s is not found", id)
		return QuestionFormat{}, errors.New(newError)
	}

	if dataInput.Title != "" || len(dataInput.Title) != 0 {
		dataUpdate["title"] = dataInput.Title
	}
	if dataInput.Content != "" || len(dataInput.Content) != 0 {
		dataUpdate["content"] = dataInput.Content
	}
	if dataInput.CategoryID != 0 {
		dataUpdate["category_id"] = dataInput.CategoryID
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	questionUpdate, err := s.repository.UpdateQuestion(id, dataUpdate)

	if err != nil {
		return QuestionFormat{}, err
	}

	formatQuestion := FormattingQuestion(questionUpdate)

	return formatQuestion, nil
}

func (s *questionService) DeleteQuestionById(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	question, err := s.repository.FindQuestionById(id)

	if err != nil {
		return nil, err
	}

	if question.ID == 0 {
		newError := fmt.Sprintf("question id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteQuestion(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("delete question id %s succeed", id)

	formatDelete := user.FormatDelete(msg)

	return formatDelete, nil
}
