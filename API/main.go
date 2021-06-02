package main

import (
	"os"
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// jobsapi.StoreToDB()
	r := gin.Default()
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	routes.UserRoute(r)
	routes.QuestionRoute(r)
	routes.CategoryRoute(r)
	routes.AnswerRoute(r)
	routes.JobRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
