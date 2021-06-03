package handler

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/auth"
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/Excellent-Echo/stuckOverflow/API/API/question"
	"github.com/gin-gonic/gin"
)

type questionHandler struct {
	questionService question.QuestionService
	authService     auth.Service
}

func NewQuestionHandler(questionService question.QuestionService, authService auth.Service) *questionHandler {
	return &questionHandler{questionService, authService}
}

func (h *questionHandler) ShowAllQuestionsHandler(c *gin.Context) {
	questions, err := h.questionService.FindAllQuestions()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("get all questions succeed", 200, "success", questions)
	c.JSON(200, userResponse)

}

func (h *questionHandler) CreateQuestionHandler(c *gin.Context) {
	var inputQuestion entity.QuestionInput

	if err := c.ShouldBindJSON(&inputQuestion); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	// userID := int(c.MustGet("currentUser").(int))

	response, err := h.questionService.SaveNewQuestion(inputQuestion)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	newResponse := helper.APIResponse("insert new question succeed", 201, "success", response)
	c.JSON(201, newResponse)
}

func (h *questionHandler) ShowQuestionByIdHandler(c *gin.Context) {
	id := c.Param("id")

	question, err := h.questionService.FindQuestionById(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("get question succeed", 200, "success", question)
	c.JSON(200, response)
}

func (h *questionHandler) UpdateQuestionByIdHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	var updateQuestionInput entity.UpdateQuestionInput

	if err := c.ShouldBindJSON(&updateQuestionInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	question, err := h.questionService.UpdateQuestionById(id, updateQuestionInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	idParam := int(question.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("update question succeed", 200, "success", question)
	c.JSON(200, response)
}

func (h *questionHandler) DeleteByQuestionHandler(c *gin.Context) {
	id := c.Param("id")

	questionDetail, _ := h.questionService.FindQuestionById(id)

	idParam := int(questionDetail.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	question, err := h.questionService.DeleteQuestionById(id)

	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("question was deleted successfully", 200, "success", question)
	c.JSON(200, response)
}
