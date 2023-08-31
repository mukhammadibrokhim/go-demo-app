package service

import (
	"go-demo-app/models"
	"go-demo-app/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {

	err := s.UserRepo.CreateUser(user)
	return err
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {

	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
