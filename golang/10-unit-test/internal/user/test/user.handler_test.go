package test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/apperror"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/dto"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/model"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/user"
	mock_router "github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/mocks/router"
	mock_user "github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/mocks/user"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTest struct {
	suite.Suite
	controller    *gomock.Controller
	CreateUserReq *dto.CreaterUserRequest
	User          *model.User
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTest))
}

func (t *UserHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.User = &model.User{
		Email:    faker.Email(),
		Password: faker.Password(),
	}
	t.CreateUserReq = &dto.CreaterUserRequest{
		Email:    t.User.Email,
		Password: t.User.Password,
	}
}

// there are 4 parts in a unit test
func (t *UserHandlerTest) TestCreateSuccess() {
	// 1. Create mock object + dependencies injection
	svc := mock_user.NewMockService(t.controller)
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(svc)

	// 2. Define expected result, input, output
	expectedResp := &dto.CreaterUserResponse{
		User: *t.User,
	}

	// 3. Define mock behavior
	// must be the same flow as the actual code i.e. bind -> service -> response in user.handler.go
	// Bind(&dto.CreaterUserRequest{}) because when we bind the request in the handler, the variable is also empty struct
	ctx.EXPECT().Bind(&dto.CreaterUserRequest{}).SetArg(0, *t.CreateUserReq)
	svc.EXPECT().Create(t.CreateUserReq).Return(expectedResp, nil)
	ctx.EXPECT().JSON(http.StatusCreated, expectedResp)

	// 4. Call the function
	hdr.Create(ctx)
}

func (t *UserHandlerTest) TestCreateBindError() {
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(nil)
	// we don't need to call user service because the error is in the binding (it is before the service call)
	// so we don't need to create a mock for user service, less code = better

	ctx.EXPECT().Bind(&dto.CreaterUserRequest{}).Return(errors.New("error"))
	ctx.EXPECT().BadRequestError("error")

	hdr.Create(ctx)
}

func (t *UserHandlerTest) TestCreateUserServiceError() {
	svc := mock_user.NewMockService(t.controller)
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(svc)
	// we don't need to call user service because the error is in the binding (it is before the service call)
	// so we don't need to create a mock for user service, less code = better

	ctx.EXPECT().Bind(&dto.CreaterUserRequest{}).SetArg(0, *t.CreateUserReq)
	svc.EXPECT().Create(t.CreateUserReq).Return(nil, apperror.InternalServer) // pick any apperror you want, we just want
	//check that the same error is passed to the ctx.EXPECT().ResponseError
	ctx.EXPECT().ResponseError(apperror.InternalServer)

	hdr.Create(ctx)
}

func (t *UserHandlerTest) TestFindOneSuccess() {
	svc := mock_user.NewMockService(t.controller)
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(svc)
	id := faker.UUIDDigit()

	expectedResp := &dto.FindOneUserResponse{
		User: *t.User,
	}

	ctx.EXPECT().Param("id").Return(id)
	svc.EXPECT().FindOne(&dto.FindOneUserRequest{Id: id}).Return(expectedResp, nil)
	ctx.EXPECT().JSON(http.StatusOK, expectedResp)

	hdr.FindOne(ctx)
}

func (t *UserHandlerTest) TestFindOneParamEmpty() {
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(nil)

	ctx.EXPECT().Param("id").Return("")
	ctx.EXPECT().BadRequestError("id is required in url param")

	hdr.FindOne(ctx)
}

func (t *UserHandlerTest) TestFindOneServiceError() {
	svc := mock_user.NewMockService(t.controller)
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(svc)
	id := faker.UUIDDigit()

	ctx.EXPECT().Param("id").Return(id)
	svc.EXPECT().FindOne(&dto.FindOneUserRequest{Id: id}).Return(nil, apperror.InternalServer)
	ctx.EXPECT().ResponseError(apperror.InternalServer)

	hdr.FindOne(ctx)
}

func (t *UserHandlerTest) TestDeleteSuccess() {
	svc := mock_user.NewMockService(t.controller)
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(svc)
	id := faker.UUIDDigit()

	expectedResp := &dto.DeleteUserResponse{
		Success: true,
	}

	ctx.EXPECT().Param("id").Return(id)
	svc.EXPECT().Delete(&dto.DeleteUserRequest{Id: id}).Return(expectedResp, nil)
	ctx.EXPECT().JSON(http.StatusOK, expectedResp)

	hdr.Delete(ctx)
}

func (t *UserHandlerTest) TestDeleteParamEmpty() {
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(nil)

	ctx.EXPECT().Param("id").Return("")
	ctx.EXPECT().BadRequestError("id is required in url param")

	hdr.Delete(ctx)
}

func (t *UserHandlerTest) TestDeleteServiceError() {
	svc := mock_user.NewMockService(t.controller)
	ctx := mock_router.NewMockContext(t.controller)
	hdr := user.NewHandler(svc)
	id := faker.UUIDDigit()

	ctx.EXPECT().Param("id").Return(id)
	svc.EXPECT().Delete(&dto.DeleteUserRequest{Id: id}).Return(nil, apperror.InternalServer)
	ctx.EXPECT().ResponseError(apperror.InternalServer)

	hdr.Delete(ctx)
}

func (t *UserHandlerTest) TearDownTest() {
	t.controller.Finish()
}
