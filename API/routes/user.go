package routes

import (
	"github.com/Excellent-Echo/stuckOverflow/API/API/auth"
	"github.com/Excellent-Echo/stuckOverflow/API/API/config"
	"github.com/Excellent-Echo/stuckOverflow/API/API/handler"
	"github.com/Excellent-Echo/stuckOverflow/API/API/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.UserNewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowAllUsersHandler)
	r.GET("/users/:id", handler.Middleware(userService, authService), userHandler.ShowUserByIdHandler)
	r.PUT("/users/:id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
	r.DELETE("/users/:id", handler.Middleware(userService, authService), userHandler.DeleteByUserIDHandler)
	r.PUT("/users/:id/avatar", handler.Middleware(userService, authService), userHandler.UpdateAvatarByIDHandler)

}
