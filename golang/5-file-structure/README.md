# 5-file-structure
We are following (not strictly) the [standard Go project layout](https://github.com/golang-standards/project-layout). The standard Go project layout is a project layout designed for Go applications. It is a way to organize your code in a way that is easy to understand and maintain. Moreover, It makes your code more modular, testable, and maintainable.

From this part onwards, we will implement the user model part of RPKM66.

## Request Flow
HTTP request from frontend -> router -> handler -> service -> repository -> database/store

If we're doing microservices e.g. gRPC
### Gateway Service
HTTP request from frontend -> router -> handler -> gRPC client (service of Microservice A) -> gRPC server
### Microservice A
gRPC server -> service -> repository -> database/store