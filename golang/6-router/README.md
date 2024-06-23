# 6-router

# Running the project
- copy `.env.template` to `.env`
- run `docker-compose up -d` or `make docker` to start the database
- run `go run cmd/main.go` or `make server` for normal run, `air` for hot reload

# Changes made
- added gin `router` to make the app into API
- registered the services to the router