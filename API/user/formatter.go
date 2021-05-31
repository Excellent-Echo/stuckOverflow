package user

import (
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
)

type UserFormat struct {
	UserID    uint32 `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Location  string `json:"location"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"delete_time"`
}

func FormattingUser(user entity.User) UserFormat {
	userFormat := UserFormat{
		UserID:    user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.User.FirstName,
		LastName:  user.User.LastName,
		Avatar:    user.User.Avatar,
		Location:  user.User.Location,
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
