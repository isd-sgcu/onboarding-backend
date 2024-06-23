# 8-handler

# Running the project
- copy `.env.template` to `.env`
- run `docker-compose up -d` or `make docker` to start the database
- run `go run cmd/main.go` or `make server` for normal run, `air` for hot reload

# Handler
`main.go` is a bit messy and has a lot of imperative code. We will move the imperative code to the handler to make `main.go` more declarative.