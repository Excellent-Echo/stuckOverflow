package routes

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/answer"
	"github.com/Excellent-Echo/stuckOverflow/API/API/handler"
	"github.com/gin-gonic/gin"
)

var (
	answerRepository = answer.NewRepository(DB)
	answerService    = answer.NewService(answerRepository)
	answerHandler    = handler.NewAnswerHandler(answerService, authService)
)

func AnswerRoute(r *gin.Engine) {
	r.POST("/questions/:id", handler.Middleware(userService, authService), answerHandler.CreateAnswerHandler)
	r.PUT("/answers/:id", handler.Middleware(userService, authService), answerHandler.UpdateAnswerHandler)
	r.DELETE("/answers/:id", handler.Middleware(userService, authService), answerHandler.DeleteAnswerHandler)
}
