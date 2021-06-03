package routes

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/handler"
	"github.com/Excellent-Echo/stuckOverflow/API/API/question"
	"github.com/gin-gonic/gin"
)

var (
	questionRepository = question.NewRepository(DB)
	questionService    = question.QuestionNewService(questionRepository)
	questionHandler    = handler.NewQuestionHandler(questionService, authService)
)

func QuestionRoute(r *gin.Engine) {
	r.GET("/questions", questionHandler.ShowAllQuestionsHandler)
	r.POST("/questions/ask", questionHandler.CreateQuestionHandler)
	r.GET("/questions/:id", questionHandler.ShowQuestionByIdHandler)
	r.PUT("/questions/:id/edit", handler.Middleware(userService, authService), questionHandler.UpdateQuestionByIdHandler)
	r.DELETE("/questions/:id", handler.Middleware(userService, authService), questionHandler.DeleteByQuestionHandler)

}
