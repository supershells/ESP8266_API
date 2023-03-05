package services

import (
	"esp8266_api/errs"
	"esp8266_api/logs"
	"esp8266_api/repositories"
	passwordhash "esp8266_api/util/passwordhash"

	"github.com/google/uuid"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s userService) GetUserOne(username string) (*UserResponse, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("User not found")
	}

	userResponse := UserResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
		Dept:     user.Dept,
	}

	return &userResponse, nil
}

func (s userService) GetUserAll() ([]UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("User not found")
	}

	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Role:     user.Role,
			Dept:     user.Dept,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (s userService) GetUser(request UserLogin) (*UserResponse, error) {
	user, err := s.userRepo.GetByUsername(request.Username)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("Authentication failed")
	}
	passwordIsCorrect := passwordhash.CheckPasswordHash(request.Password, user.Password)
	if !passwordIsCorrect {
		logs.Error(err)
		return nil, errs.NewUnauthorizedError("User or password is incorrect")
	}

	userResponse := &UserResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
		Dept:     user.Dept,
	}

	return userResponse, nil
}

func (s userService) NewUser(request NewUserRequest) (*UserResponse, error) {

	passwordhash, err := passwordhash.HashPassword(request.Password)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError("Error while hashing password")
	}

	user := repositories.User{
		UserID:   uuid.New().String(),
		Username: request.Username,
		Password: string(passwordhash),
		Role:     request.Role,
		Dept:     request.Dept,
		Status:   1,
	}

	newUser, err := s.userRepo.Create(user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError("Error while creating user")
	}
	response := UserResponse{
		UserID:   newUser.UserID,
		Username: newUser.Username,
		Role:     newUser.Role,
		Dept:     newUser.Dept,
	}

	return &response, nil
}
