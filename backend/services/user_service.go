package services

import (
	"jopanel/backend/config"
	"jopanel/backend/models"
	"jopanel/backend/utils"
)

type UserService interface {
	CreateUser(input models.User, password string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	UpdateUser(id uint, input models.User) (models.User, error)
	DeleteUser(id uint) error
	ChangeStatus(id uint, status string) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(user models.User, password string) (models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return user, err
	}
	user.PasswordHash = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}

	// TODO: Trigger Agent to create system user
	return user, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return user, err
}

func (s *userService) UpdateUser(id uint, input models.User) (models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	user.Username = input.Username // Ideally username shouldn't change easily as it affects system user
	user.Email = input.Email
	user.Role = input.Role

	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	// TODO: Trigger Agent to delete system user
	return config.DB.Delete(&models.User{}, id).Error
}

func (s *userService) ChangeStatus(id uint, status string) error {
	return config.DB.Model(&models.User{}).Where("id = ?", id).Update("status", status).Error
}
