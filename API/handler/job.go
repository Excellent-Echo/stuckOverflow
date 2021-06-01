package handler

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/Excellent-Echo/stuckOverflow/API/API/job"
	"github.com/gin-gonic/gin"
)

type jobHandler struct {
	jobService    job.JobService
	jobRepository job.JobRepository
}

func NewJobHandler(jobService job.JobService, jobRepository job.JobRepository) *jobHandler {
	return &jobHandler{jobService, jobRepository}
}

func (h *jobHandler) ShowAllJobsHandler(c *gin.Context) {
	jobs, err := h.jobService.FindAllJobs()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("get all jobs succeed", 200, "success", jobs)
	c.JSON(200, userResponse)

}

func (h *jobHandler) GetAllJobsWithQuery(c *gin.Context) {
	pagination := h.jobService.GeneratePaginationFromRequest(c)
	var job entity.Jobs
	jobLists, err := h.jobRepository.GetAllJobsQuery(&job, &pagination)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("get all jobs succeed", 200, "success", jobLists)
	c.JSON(200, userResponse)
}
