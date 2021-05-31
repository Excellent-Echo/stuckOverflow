package user

import (
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
)

type UserInputFormat struct {
	UserID   uint32 `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"delete_time"`
}

func FormattingUser(user entity.User) UserInputFormat {
	userFormat := UserInputFormat{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}

	return userFormat
}

func FormatDelete(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
