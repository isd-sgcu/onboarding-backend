package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/apperror"
)

type Context interface {
	JSON(statusCode int, obj interface{})
	ResponseError(err *apperror.AppError)
	BadRequestError(err string)
	InternalServerError(err string)
	NewUUID() uuid.UUID
	Bind(obj interface{}) error
	Param(key string) string
	Query(key string) string
	PostForm(key string) string
}

type contextImpl struct {
	*gin.Context
}

func NewContext(c *gin.Context) Context {
	return &contextImpl{c}
}

func (c *contextImpl) JSON(statusCode int, obj interface{}) {
	c.Context.JSON(statusCode, obj)
}

func (c *contextImpl) ResponseError(err *apperror.AppError) {
	c.JSON(err.HttpCode, gin.H{"error": err.Error()})
}

func (c *contextImpl) InternalServerError(err string) {
	c.JSON(apperror.InternalServer.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) BadRequestError(err string) {
	c.JSON(apperror.BadRequest.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) NewUUID() uuid.UUID {
	return uuid.New()
}

func (c *contextImpl) Bind(obj interface{}) error {
	return c.Context.Bind(obj)
}

func (c *contextImpl) Param(key string) string {
	return c.Context.Param(key)
}

func (c *contextImpl) Query(key string) string {
	return c.Context.Query(key)
}

func (c *contextImpl) PostForm(key string) string {
	return c.Context.PostForm(key)
}
