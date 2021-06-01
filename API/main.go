package main

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/jobsapi"
)

func main() {
	// r := gin.Default()
	// r.Use(cors.Default())

	// routes.UserRoute(r)
	// routes.QuestionRoute(r)
	// routes.CategoryRoute(r)
	// routes.AnswerRoute(r)

	// port := os.Getenv("PORT")
	// r.Run(":" + port)
	jobsapi.StoreToDB()
}
