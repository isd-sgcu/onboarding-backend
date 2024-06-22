# 11-logging-error

# Logging
Every time you see a `if err != nil` or other scenarios deemed appropriate, log it (it will save a LOT of time debugging). Only `service` and `handler` should log. Other components like `repository` should not log because we're already passing the errors up to the service or handler layer. We should have a universal standard for logging. We will use `zap` for logging example.
```go
// main.go, we make a authSvc logger to distinguish this service's logs from other services
authSvc := auth.NewService(userSvc, logger.Named("authSvc"))
```
```go
type Service interface {
	proto.AuthServiceServer
}

type serviceImpl struct {
	proto.UnimplementedAuthServiceServer
	userSvc  user.Service
	log      *zap.Logger
}

// we pass the logger from the main.go as dependency of the service
func NewService(userSvc user.Service, log *zap.Logger) Service {
	return &serviceImpl{
		userSvc:  userSvc,
		log:      log,
	}
}
```

```go
// s.log.Named(<method_name>).Error("<dependency_method>: ", zap.Error(err))
// We should pass the full error message to the log so that we can easily debug the error
func (s *serviceImpl) SignIn(_ context.Context, in *proto.SignInRequest) (res *proto.SignInResponse, err error) {
	user, err := s.userSvc.FindByEmail(context.Background(), &userProto.FindByEmailRequest{Email: in.Email})
	if err != nil {
		s.log.Named("SignIn").Error("FindByEmail: ", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
    // other logic
}
```

# Error Handling
## Gateway Service
It it the service that talks to frontend, so we should make it user-friendly.
- Service: it receives res/err from other services, so we should do a `switch-case` to check for each error status code and return the appropriate response. For errors you don't know the reason, put the full error message in the response.
```go
// omitted logging for brevity
func (s *serviceImpl) FindOneBaan(req *dto.FindOneBaanRequest) (*dto.FindOneBaanResponse, *apperrors.AppError) {
	res, err := s.client.FindOneBaan(ctx, &baanProto.FindOneBaanRequest{
		Id: req.Id,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok { // this has no err to deduce the status code, so we return InternalServer with default message "Internal Server Error"
			return nil, apperrors.InternalServer 
		}
		switch st.Code() {
		case codes.NotFound: // put appropriate error message so that the frontend team can understand
			return nil, apperrors.NotFoundError("Baan not found")
		case codes.Internal: // for errors you don't know the reason, put the full error message
			return nil, apperrors.InternalServerError(err.Error())
		default: // for default case, put ServiceUnavailable
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.FindOneBaanResponse{
		Baan: utils.ProtoToDto(res.Baan),
	}, nil
}
```

## Other Services
- Non-service components e.g. Repository, should return the error as is. The service will handle the error.
- Service: put the full error message in the response so that we can easily debug the error. For status code, we should return the appropriate status code based on the error we receive.
```go
func (s *serviceImpl) SignUp(_ context.Context, in *proto.SignUpRequest) (res *proto.SignUpResponse, err error) {
    hashedPassword, err := s.bcrypt.GenerateHashedPassword(in.Password)
	if err != nil {
		s.log.Named("SignUp").Error("GenerateHashedPassword: ", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error()) // we don't know the reason, so we put it as internal server error with full error message
	}
    // other logic
}
```
```go
func (s *serviceImpl) SignIn(_ context.Context, in *proto.SignInRequest) (res *proto.SignInResponse, err error) {
	user, err := s.userSvc.FindByEmail(context.Background(), &userProto.FindByEmailRequest{Email: in.Email})
	if err != nil {
		s.log.Named("SignIn").Error("FindByEmail: ", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error()) // cannot find user with this email when logging in, so we know it's unauthenticated
	}
    // other logic
}
```