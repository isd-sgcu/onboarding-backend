package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type UserHandlerTest struct {
	suite.Suite
	controller *gomock.Controller
	logger     *zap.Logger
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTest))
}

func (t *UserHandlerTest) SetupTest() {}

func (t *UserHandlerTest) TestSignUpSuccess() {

}
