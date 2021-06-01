package job

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetAllJobs() ([]entity.Jobs, error)
	GetAllJobsQuery(jobs *entity.Jobs, pagination *entity.Pagination) (*[]entity.Jobs, error)
	GetJobByID(id string) (entity.Jobs, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllJobs() ([]entity.Jobs, error) {
	var Jobs []entity.Jobs

	err := r.db.Find(&Jobs).Error
	if err != nil {
		return Jobs, err
	}

	return Jobs, nil
}

func (r *Repository) GetAllJobsQuery(job *entity.Jobs, pagination *entity.Pagination) (*[]entity.Jobs, error) {
	var jobs []entity.Jobs

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&entity.Jobs{}).Where(job).Find(&jobs)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}

	return &jobs, nil
}

func (r *Repository) GetJobByID(id string) (entity.Jobs, error) {
	var job entity.Jobs

	if err := r.db.Where("id = ?", id).Find(&job).Error; err != nil {
		return job, err
	}

	return job, nil
}
