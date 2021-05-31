package question

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"gorm.io/gorm"
)

type QuestionRepository interface {
	FindAllQuestions() ([]entity.Questions, error)
	PostQuestion(question entity.Questions) (entity.Questions, error)
	FindQuestionById(id string) (entity.Questions, error)
	UpdateQuestion(id string, dataUpdate map[string]interface{}) (entity.Questions, error)
	DeleteQuestion(id string) (string, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) FindAllQuestions() ([]entity.Questions, error) {
	var Questions []entity.Questions

	err := r.db.Preload("User").Preload("Category").Preload("Answers").Find(&Questions).Error
	if err != nil {
		return Questions, err
	}

	return Questions, nil
}

func (r *Repository) PostQuestion(question entity.Questions) (entity.Questions, error) {
	if err := r.db.Preload("Category").Create(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func (r *Repository) FindQuestionById(id string) (entity.Questions, error) {
	var question entity.Questions

	if err := r.db.Where("id = ?", id).Preload("User").Preload("Category").Preload("Answers").Find(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func (r *Repository) UpdateQuestion(id string, dataUpdate map[string]interface{}) (entity.Questions, error) {
	var question entity.Questions

	if err := r.db.Model(&question).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return question, err
	}

	if err := r.db.Where("id = ?", id).Preload("User").Preload("Category").Preload("Answers").Find(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func (r *Repository) DeleteQuestion(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.Questions{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
