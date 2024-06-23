package mock_user

import (
	apperror "github.com/isd-sgcu/onboarding-backend/golang/6-router/apperror"
	dto "github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/dto"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (c *ServiceMock) Create(req *dto.CreaterUserRequest) (res *dto.CreaterUserResponse, err *apperror.AppError) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*dto.CreaterUserResponse)
	}

	return res, &apperror.AppError{
		Id:       args.Error(1).Error(),
		HttpCode: args.Int(2),
	}
}

func (c *ServiceMock) FindOne(req *dto.FindOneUserRequest) (res *dto.FindOneUserResponse, err *apperror.AppError) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*dto.FindOneUserResponse)
	}

	return res, &apperror.AppError{
		Id:       args.Error(1).Error(),
		HttpCode: args.Int(2),
	}
}

func (c *ServiceMock) Delete(req *dto.DeleteUserRequest) (res *dto.DeleteUserResponse, err *apperror.AppError) {
	args := c.Called(req)

	if args.Get(0) != nil {
		return args.Get(0).(*dto.DeleteUserResponse), nil
	}

	return res, &apperror.AppError{
		Id:       args.Error(1).Error(),
		HttpCode: args.Int(2),
	}
}
