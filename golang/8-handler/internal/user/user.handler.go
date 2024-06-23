package user

import (
	"net/http"

	"github.com/isd-sgcu/onboarding-backend/golang/8-handler/internal/dto"
	"github.com/isd-sgcu/onboarding-backend/golang/8-handler/internal/router"
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
		// c.JSON(400, err)
		// use this custom BadRequestError method instead of c.JSON(400, err)
		c.BadRequestError(err.Error())
		return
	}

	createdUser, apperr := h.svc.Create(&createUserDto)
	if apperr != nil {
		// c.JSON(500, err)
		// use this custom ResponseError instead because if service returned a 401, it will return 401 (correct) instead of 500 (wrong)
		// use ResponseError when dealing with apperror.AppError, esp. from services
		c.ResponseError(apperr)
		return
	}

	// c.JSON(201, createdUser)
	// use http.StatusCreated instead of 201, it's less ambiguous
	c.JSON(http.StatusCreated, createdUser)
}

func (h *handlerImpl) FindOne(c router.Context) {
	id := c.Param("id")
	if id == "" {
		// c.JSON(400, "id is required in url param")
		c.BadRequestError("id is required in url param")
		return
	}

	res, apperr := h.svc.FindOne(&dto.FindOneUserRequest{Id: id})
	if apperr != nil {
		// c.JSON(500, apperr)
		c.ResponseError(apperr)
		return
	}

	// c.JSON(200, res)
	c.JSON(http.StatusOK, res)
}

func (h *handlerImpl) Delete(c router.Context) {
	id := c.Param("id")
	if id == "" {
		// c.JSON(400, "id is required in url param")
		c.BadRequestError("id is required in url param")
		return
	}

	res, apperr := h.svc.Delete(&dto.DeleteUserRequest{Id: id})
	if apperr != nil {
		// c.JSON(500, apperr)
		c.ResponseError(apperr)
		return
	}

	// c.JSON(200, res)
	c.JSON(http.StatusOK, res)
}
