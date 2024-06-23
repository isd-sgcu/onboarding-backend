# 9-mocks

# Running the project
- copy `.env.template` to `.env`
- run `docker-compose up -d` or `make docker` to start the database
- run `go run cmd/main.go` or `make server` for normal run, `air` for hot reload

# Why mocks?
In unit testing, for example, the user service has user repository as a dependency. We don't want to test the user repository in the user service test. We want to test the user service only. So, we create a mock of the user repository and use it in the user service test.

User repo mock is:
- a fake user repo that implements the user repo interface.
- we tell them what to return when what repo's method is called.

By doing this, we can test the user service without worrying about the user repository. We only care about the logic of the user service in unit testing of user service.

# Creating mocks
There are 2 ways to create mocks:
1. Manually creating the mocks
2. Using `mockgen` from `github.com/golang/mock/mockgen`

## Manually creating the mocks
see `/internal/user/user.service_manual.go` to get an idea of how mocks actually work.

## Using `mockgen` (preferred)
1. Install `mockgen`
```bash
# install mockgen
go get github.com/golang/mock/mockgen@v1.4.4

# generate mocks
make mockgen
```
mocks: `/mocks/user/user.repository.go` and `/mocks/user/user.service.go`