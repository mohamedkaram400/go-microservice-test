package service

import (
	"context"

	"github.com/mohamed/microservices/user-service/dto"
	"github.com/mohamed/microservices/user-service/entity"
	"github.com/mohamed/microservices/user-service/repository"
)

type UserService struct {
    userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*entity.User, error) {

	user := &entity.User{
		Name:  req.Name,
        Email: req.Email,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, id int) (*entity.User, error) {
    return s.userRepository.GetUserByID(ctx, id)
}