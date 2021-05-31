package handler

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/category"
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.CategoryService
}

func NewCategoryHandler(categoryService category.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) ShowAllCategoryHandler(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("get all categories succeed", 200, "success", categories)
	c.JSON(200, response)

}

func (h *categoryHandler) CreateCategoryHandler(c *gin.Context) {
	var categoryInput entity.CategoryInput

	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	response, err := h.categoryService.SaveNewCategory(categoryInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	newResponse := helper.APIResponse("insert new category succeed", 201, "success", response)
	c.JSON(201, newResponse)
}

func (h *categoryHandler) ShowCategoryByNameHandler(c *gin.Context) {
	categoryName := c.Param("category_name")

	category, err := h.categoryService.FindCategoryByName(categoryName)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("get category succeed", 200, "success", category)
	c.JSON(200, response)
}
