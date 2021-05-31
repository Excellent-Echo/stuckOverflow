package handler

import (
	"strconv"

	"github.com/Excellent-Echo/stuckOverflow/API/API/answer"
	"github.com/Excellent-Echo/stuckOverflow/API/API/auth"
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/gin-gonic/gin"
)

type answerHandler struct {
	answerService answer.AnswerService
	authService   auth.Service
}

func NewAnswerHandler(answerService answer.AnswerService, authService auth.Service) *answerHandler {
	return &answerHandler{answerService, authService}
}

func (h *answerHandler) CreateAnswerHandler(c *gin.Context) {
	var inputAnswer entity.AnswerInput

	if err := c.ShouldBindJSON(&inputAnswer); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userID := int(c.MustGet("currentUser").(int))
	id := c.Params.ByName("id")
	questionID, _ := strconv.Atoi(id)

	response, err := h.answerService.PostNewAnswer(questionID, userID, inputAnswer)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	newResponse := helper.APIResponse("insert new answer succeed", 201, "success", response)
	c.JSON(201, newResponse)
}

func (h *answerHandler) UpdateAnswerHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	var updateAnswer entity.AnswerInput

	if err := c.ShouldBindJSON(&updateAnswer); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	answer, err := h.answerService.UpdateAnswerByID(id, updateAnswer)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	idParam := int(answer.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("update answer succeed", 200, "success", answer)
	c.JSON(200, response)
}

func (h *answerHandler) DeleteAnswerHandler(c *gin.Context) {
	id := c.Param("id")

	answerDetail, _ := h.answerService.FindAnswerByID(id)

	idParam := int(answerDetail.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	answer, err := h.answerService.DeleteAnswerByID(id)

	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("question was deleted successfully", 200, "success", answer)
	c.JSON(200, response)
}
