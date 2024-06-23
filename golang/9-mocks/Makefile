server:
	go run cmd/main.go

docker:
	docker-compose up -d 

mock-gen:
	mockgen -source ./internal/user/user.repository.go -destination ./mocks/user/user.repository.go
	mockgen -source ./internal/user/user.service.go -destination ./mocks/user/user.service.go