package job

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/gin-gonic/gin"
)

type JobService interface {
	FindAllJobs() ([]entity.Jobs, error)
	GeneratePaginationFromRequest(c *gin.Context) entity.Pagination
	FindJobByID(id string) (entity.Jobs, error)
}

type jobService struct {
	repository JobRepository
}

func JobNewService(repository JobRepository) *jobService {
	return &jobService{repository}
}

func (s *jobService) FindAllJobs() ([]entity.Jobs, error) {
	jobs, err := s.repository.GetAllJobs()

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (s *jobService) GeneratePaginationFromRequest(c *gin.Context) entity.Pagination {
	limit := 500
	page := 1
	sort := "publication_date desc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return entity.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func (s *jobService) FindJobByID(id string) (entity.Jobs, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.Jobs{}, err
	}

	job, err := s.repository.GetJobByID(id)

	if err != nil {
		return entity.Jobs{}, err
	}

	if job.ID == 0 {
		newError := fmt.Sprintf("job id %s is not found", id)
		return entity.Jobs{}, errors.New(newError)
	}

	return job, nil
}
