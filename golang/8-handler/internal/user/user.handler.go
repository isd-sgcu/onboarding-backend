package user

import (
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/dto"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/router"
)

type Handler interface {
	// handler's methods return nothing because it's job is to respond the HTTP requests via the router.Context
	Create(c router.Context)
	FindOne(c router.Context)
	Delete(c router.Context)
}

type handlerImpl struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{
		svc: svc,
	}
}

func (h *handlerImpl) Create(c router.Context) {
	var createUserDto dto.CreaterUserRequest
	if err := c.Bind(&createUserDto); err != nil {
		c.JSON(400, err)
		return
	}

	createdUser, err := h.svc.Create(&createUserDto)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, createdUser)
}

func (h *handlerImpl) FindOne(c router.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, "id is required in url param")
		return
	}

	total, apperr := h.svc.FindOne(&dto.FindOneUserRequest{Id: id})
	if apperr != nil {
		c.JSON(500, apperr)
		return
	}

	c.JSON(200, total)
}

func (h *handlerImpl) Delete(c router.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, "id is required in url param")
		return
	}

	total, apperr := h.svc.Delete(&dto.DeleteUserRequest{Id: id})
	if apperr != nil {
		c.JSON(500, apperr)
		return
	}

	c.JSON(200, total)
}
