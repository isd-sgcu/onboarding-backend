# 10-unit-test

# Running the project
- copy `.env.template` to `.env`
- run `docker-compose up -d` or `make docker` to start the database
- run `go run cmd/main.go` or `make server` for normal run, `air` for hot reload

# Why unit testing?
Unit testing is important because:
- it helps us to find bugs early in the development cycle
- it gives us confidence in our code
- it helps us to refactor the code with confidence
- it helps us to write better code

# Unit testing
In this part, we will show how to write unit tests for the user handler.

see `/internal/user/test/user.handler_test.go` for the tests.