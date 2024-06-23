package dto

import "github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/model"

// dtos are used for sending data in and out of the service layer
type CreaterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreaterUserResponse struct {
	User model.User `json:"user"`
}

type FindOneUserRequest struct {
	Id string `json:"id"`
}

type FindOneUserResponse struct {
	User model.User `json:"user"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}

type DeleteUserResponse struct {
	Success bool `json:"success"`
}
