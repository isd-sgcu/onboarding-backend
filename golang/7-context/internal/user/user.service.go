package user

import (
	"github.com/isd-sgcu/onboarding-backend/golang/7-context/apperror"
	"github.com/isd-sgcu/onboarding-backend/golang/7-context/internal/dto"
	"github.com/isd-sgcu/onboarding-backend/golang/7-context/internal/model"
)

type Service interface {
	Create(req *dto.CreaterUserRequest) (*dto.CreaterUserResponse, *apperror.AppError)
	FindOne(req *dto.FindOneUserRequest) (*dto.FindOneUserResponse, *apperror.AppError)
	Delete(req *dto.DeleteUserRequest) (*dto.DeleteUserResponse, *apperror.AppError)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (c *serviceImpl) Create(req *dto.CreaterUserRequest) (*dto.CreaterUserResponse, *apperror.AppError) {
	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := c.repo.Create(user)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	return &dto.CreaterUserResponse{
		User: *user,
	}, nil
}

func (c *serviceImpl) FindOne(req *dto.FindOneUserRequest) (*dto.FindOneUserResponse, *apperror.AppError) {
	user := &model.User{}
	err := c.repo.FindOne(req.Id, user)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	return &dto.FindOneUserResponse{
		User: *user,
	}, nil
}

func (c *serviceImpl) Delete(req *dto.DeleteUserRequest) (*dto.DeleteUserResponse, *apperror.AppError) {
	err := c.repo.Delete(req.Id)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	return &dto.DeleteUserResponse{
		Success: true,
	}, nil
}
