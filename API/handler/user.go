package handler

import (
	"fmt"

	"github.com/Excellent-Echo/stuckOverflow/API/API/auth"
	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"github.com/Excellent-Echo/stuckOverflow/API/API/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	response, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		emailDuplicateError := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'email'", inputUser.Email)
		userNameDuplicateError := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'user_name'", inputUser.UserName)
		if err.Error() == emailDuplicateError {
			responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": "Email sudah digunakan"})

			c.JSON(500, responseError)
			return
		}
		if err.Error() == userNameDuplicateError {
			responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": "Username sudah digunakan"})

			c.JSON(500, responseError)
			return
		}
	}

	userResponse := helper.APIResponse("insert user data succeed", 201, "success", response)
	c.JSON(201, userResponse)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("input data error", 401, "bad request", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}

	token, err := h.authService.GenerateToken(int(userData.ID))
	if err != nil {
		responseError := helper.APIResponse("input data error", 401, "bad request", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}
	response := helper.APIResponse("login user succeed", 200, "success", gin.H{"token": token})
	c.JSON(200, response)
}

func (h *userHandler) ShowAllUsersHandler(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("get all users succeed", 200, "success", users)
	c.JSON(200, userResponse)
}
