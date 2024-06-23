package test

import (
	"errors"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/apperror"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/dto"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/model"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/user"
	mock_user "github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/mocks/user"
	"github.com/stretchr/testify/suite"
)

type UserServiceTest struct {
	suite.Suite
	controller     *gomock.Controller
	CreateUserReq  *dto.CreaterUserRequest
	FindOneUserReq *dto.FindOneUserRequest
	User           *model.User
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.User = &model.User{
		Email:    faker.Email(),
		Password: faker.Password(),
	}
	t.CreateUserReq = &dto.CreaterUserRequest{
		Email:    t.User.Email,
		Password: t.User.Password,
	}
	t.FindOneUserReq = &dto.FindOneUserRequest{
		Id: faker.UUIDDigit(),
	}
}

// there are 4 parts in a unit test
func (t *UserServiceTest) TestCreateSuccess() {
	// 1. Create mock object + dependencies injection
	repo := mock_user.NewMockRepository(t.controller)
	svc := user.NewService(repo)

	// 2. Define expected result, input, output
	// for inputs we used t.CreateUserReq and t.User (declared in SetupTest)
	expectedResp := &dto.CreaterUserResponse{
		User: *t.User,
	}

	// 3. Define mock behavior
	// must be the same flow as the actual code i.e. user.service.go
	repo.EXPECT().Create(t.User).Return(nil)

	// 4. Call the function + check the result
	res, err := svc.Create(t.CreateUserReq)
	t.Equal(res, expectedResp) // compare the expected result with the actual result
	t.Nil(err)                 // check that the error is nil
}

func (t *UserServiceTest) TestCreateRepoError() {
	repo := mock_user.NewMockRepository(t.controller)
	svc := user.NewService(repo)

	repo.EXPECT().Create(t.User).Return(errors.New("error 123"))

	res, err := svc.Create(t.CreateUserReq)
	t.Nil(res)                                              // check that the result is nil
	t.Equal(err, apperror.InternalServerError("error 123")) // check that the error is the same as the expected error
}

func (t *UserServiceTest) TestFindOneSuccess() {
	repo := mock_user.NewMockRepository(t.controller)
	svc := user.NewService(repo)

	expectedResp := &dto.FindOneUserResponse{
		User: *t.User,
	}

	repo.EXPECT().FindOne(t.FindOneUserReq.Id, &model.User{}).SetArg(1, *t.User)
	// user.service.go:
	// user := &model.User{}
	// err := c.repo.FindOne(req.Id, user)
	// it c.repo.FindOne updates the user value (pass by reference)
	// user is originally an empty struct, but after the call, it will be updated with the actual value
	// so we used SetArg(1, *t.User) to update the user value
	// c.repo.FindOne on success case reutrns no error (nil), so we don't care about the return value

	res, err := svc.FindOne(t.FindOneUserReq)
	t.Equal(res, expectedResp)
	t.Nil(err)
}

func (t *UserServiceTest) TestFindOneRepoError() {
	repo := mock_user.NewMockRepository(t.controller)
	svc := user.NewService(repo)

	repo.EXPECT().FindOne(t.FindOneUserReq.Id, &model.User{}).Return(errors.New("error 321"))

	res, err := svc.FindOne(t.FindOneUserReq)
	t.Equal(err, apperror.InternalServerError("error 321"))
	t.Nil(res)
}
