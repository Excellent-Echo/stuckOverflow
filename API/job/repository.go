package job

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/jobsapi"
	"gorm.io/gorm"
)

type JobRepository interface {
	SaveAllJobs()
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) SaveAllJobs() {
	var jobs entity.Jobs

	datas := jobsapi.JobsApi()
	for _, data := range datas.Jobs {
		jobs = entity.Jobs{
			Title:           data.Title,
			Url:             data.Url,
			CompanyName:     data.CompanyName,
			JobType:         data.JobType,
			PublicationDate: data.PublicationDate,
		}

		if err := r.db.Select("title", "url", "company_name", "job_type", "publication_date").Create(&jobs).Error; err != nil {
			return
		}
	}
}
