package api

import (
	"errors"
	"strings"
)

// Contains the methods of the userService
type UserService interface {
	New(user NewUserRequest) (int, error)
}

// Let service do Ddatabase operation without knowing the implementation
type UserRepository interface {
	CreateUser(NewUserRequest) (int, error)
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}

func (u *userService) New(user NewUserRequest) (int, error) {

	// Validations
	if user.Email == "" {
		return 0, errors.New("user service - email required")
	}

	if user.Name == "" {
		return 0, errors.New("user service - name required")
	}

	if user.WeightGoal == "" {
		return 0, errors.New("user service - weight goal required")
	}

	// Normalisation
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	userID, err := u.storage.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
