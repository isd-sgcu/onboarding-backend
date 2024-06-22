# 5-architecture
We are following (not strictly) the [standard Go project layout](https://github.com/golang-standards/project-layout). The standard Go project layout is a project layout designed for Go applications. It is a way to organize your code in a way that is easy to understand and maintain. Moreover, It makes your code more modular, testable, and maintainable.

From this part onwards, we will implement the user model part of RPKM66.

# Components
## Handler
- Responsible for handling the request and returning the response.
- It should not contain any business logic.
- It should only call the service layer to perform the business logic.
- Extracts url params, query params, request body, multipart-forms, etc. into a `Data Transfer Object (DTO)`.

## Service
- Contains the business logic.
- It should not contain any database logic.
- It should only call the repository layer to perform the database logic.
- It should return the result to the handler.

## Repository
- Contains the database logic.
- It should not contain any business logic.
- It should only perform the database operations.
- It should return the result to the service.


## Request Flow
```bash
HTTP request from frontend <-> router <-> handler <-> service <-> repository <-> database/store
```
### Data being passed
- `router` <-> `handler`: HTTP request
- `handler` <-> `service`: DTO
- `service` <-> `repository`: models we defined or other parameters

### Request Flow if we're doing microservices e.g. gRPC
#### Gateway Service
```bash
HTTP request from frontend <-> router <-> handler <-> gRPC client (service of Microservice A) <-> gRPC server
```
#### Microservice A
```bash
gRPC server <-> service <-> repository <-> database/store
```

## Error flow
`Handler` is responsible for responding to the request with a response (data e.g. json + 200, 401, 400, 500, etc.). Since it is `handler <-> service <-> repository`, we should handle errors like this:
- `repository` return type pure golang `error`
- `service` receive `error` from `repository` and return to handler as a custom error equipped with HTTP status codes `apperror`
- `handler` only responses to the request with the return from `service`