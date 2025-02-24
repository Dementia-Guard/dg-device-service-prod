package services

import (
	"api/dto"
	"api/models"
	"api/repositories"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func GetUsersService() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func CreateUserService(userDTO dto.CreateUserDTO) (*models.User, error) {
	if err := validate.Struct(userDTO); err != nil {
		return nil, err
	}

	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password, // Hash password in production
	}

	// Add user to repository (implement in repository)
	return &newUser, nil
}
