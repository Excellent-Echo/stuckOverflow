package routes

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/handler"
	"github.com/Excellent-Echo/stuckOverflow/API/API/job"
	"github.com/gin-gonic/gin"
)

var (
	jobRepository = job.NewRepository(DB)
	jobService    = job.JobNewService(jobRepository)
	jobHandler    = handler.NewJobHandler(jobService, jobRepository)
)

func JobRoute(r *gin.Engine) {
	// r.GET("/jobs/all", jobHandler.ShowAllJobsHandler)
	r.GET("/jobs/all", jobHandler.GetAllJobsWithQuery)
	r.GET("/jobs/:id", jobHandler.ShowJobByIDHandler)

}
