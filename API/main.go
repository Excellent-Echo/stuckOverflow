package main

import (
	"os"

	"github.com/Excellent-Echo/stuckOverflow/API/API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
