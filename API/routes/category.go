package routes

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/category"
	"github.com/Excellent-Echo/stuckOverflow/API/API/handler"
	"github.com/gin-gonic/gin"
)

var (
	categoryRepository = category.NewRepository(DB)
	categoryService    = category.NewService(categoryRepository)
	categoryHandler    = handler.NewCategoryHandler(categoryService)
)

func CategoryRoute(r *gin.Engine) {
	r.GET("/categories", categoryHandler.ShowAllCategoryHandler)
	r.POST("/categories", handler.Middleware(userService, authService), categoryHandler.CreateCategoryHandler)
	r.GET("/categories/:category_name", categoryHandler.ShowCategoryByNameHandler)
	r.PUT("/categories/:category_name", handler.Middleware(userService, authService), categoryHandler.UpdateCategoryByNameHandler)
}
