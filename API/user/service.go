package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"github.com/Excellent-Echo/stuckOverflow/API/API/helper"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SaveNewUser(user entity.UserInput) (UserInputFormat, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
	GetAllUsers() ([]UserFormat, error)
	GetUserByID(id string) (UserFormat, error)
	UpdateUserByID(id string, dataInput entity.UpdateUserInput) (UserFormat, error)
	DeleteUserByID(id string) (interface{}, error)
	UpdateAvatarByID(pathFile string, id string) (UserFormat, error)
}

type userService struct {
	repository UserRepository
}

func UserNewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) SaveNewUser(user entity.UserInput) (UserInputFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserInputFormat{}, err
	}

	var newUser = entity.User{
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.CreateUser(newUser)
	formatUser := FormattingUserInput(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *userService) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByUserName(input.UserName)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]UserFormat, error) {
	users, err := s.repository.GetAll()

	var usersFormat []UserFormat

	for _, user := range users {
		var userFormat = FormattingUser(user)
		usersFormat = append(usersFormat, userFormat)
	}

	if err != nil {
		return usersFormat, err
	}

	return usersFormat, nil
}

func (s *userService) GetUserByID(id string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return UserFormat{}, errors.New(newError)
	}

	userFormat := FormattingUser(user)

	return userFormat, nil

}

func (s *userService) UpdateUserByID(id string, dataInput entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return UserFormat{}, errors.New(newError)
	}

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["first_name"] = dataInput.FirstName
	}
	if dataInput.LastName != "" || len(dataInput.LastName) != 0 {
		dataUpdate["last_name"] = dataInput.LastName
	}
	if dataInput.UserName != "" || len(dataInput.UserName) != 0 {
		dataUpdate["user_name"] = dataInput.UserName
	}
	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}
	if dataInput.Location != "" || len(dataInput.Location) != 0 {
		dataUpdate["location"] = dataInput.Location
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	userUpdated, err := s.repository.UpdateUserDetail(id, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormattingUser(userUpdated)

	return formatUser, nil
}

func (s *userService) DeleteUserByID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteUser(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("delete user id %s succeed", id)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}

func (s *userService) UpdateAvatarByID(pathFile string, id string) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	dataUpdate["avatar"] = pathFile

	userProfileUpdate, err := s.repository.UpdateAvatar(id, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormattingUser(userProfileUpdate)

	return formatUser, nil
}
