package user

import (
	"net/http"

	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/dto"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/router"
)

type Handler interface {
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
		c.BadRequestError(err.Error())
		return
	}

	createdUser, apperr := h.svc.Create(&createUserDto)
	if apperr != nil {
		c.ResponseError(apperr)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *handlerImpl) FindOne(c router.Context) {
	id := c.Param("id")
	if id == "" {
		c.BadRequestError("id is required in url param")
		return
	}

	total, apperr := h.svc.FindOne(&dto.FindOneUserRequest{Id: id})
	if apperr != nil {
		c.ResponseError(apperr)
		return
	}

	c.JSON(http.StatusOK, total)
}

func (h *handlerImpl) Delete(c router.Context) {
	id := c.Param("id")
	if id == "" {
		c.BadRequestError("id is required in url param")
		return
	}

	res, apperr := h.svc.Delete(&dto.DeleteUserRequest{Id: id})
	if apperr != nil {
		c.ResponseError(apperr)
		return
	}

	c.JSON(http.StatusOK, res)
}
