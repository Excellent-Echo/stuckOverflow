package answer

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	PostAnswer(answer entity.Answers) (entity.Answers, error)
	FindAnswerByID(id string) (entity.Answers, error)
	UpdateAnswer(id string, dataUpdate map[string]interface{}) (entity.Answers, error)
	DeleteAnswer(id string) (string, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) PostAnswer(answer entity.Answers) (entity.Answers, error) {
	if err := r.db.Create(&answer).Error; err != nil {
		return answer, err
	}

	return answer, nil
}

func (r *Repository) FindAnswerByID(id string) (entity.Answers, error) {
	var answer entity.Answers

	if err := r.db.Where("id = ?", id).Find(&answer).Error; err != nil {
		return answer, err
	}

	return answer, nil
}

func (r *Repository) UpdateAnswer(id string, dataUpdate map[string]interface{}) (entity.Answers, error) {
	var answer entity.Answers

	if err := r.db.Model(&answer).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return answer, err
	}

	if err := r.db.Where("id = ?", id).Find(&answer).Error; err != nil {
		return answer, err
	}

	return answer, nil
}

func (r *Repository) DeleteAnswer(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.Answers{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
