# 4-database

## Database
We will be using `docker-compose.yml` to run a PostgreSQL database. So, you should have Docker installed on your machine and have some basic docker-compose knowledge.

## New files
### docker-compose.yml
Docker Compose file. It is used to run services like databases, message brokers, etc. It is a YAML file that contains the configuration for the services.

### Makefile
Makefile is a file that contains a set of directives used to build the project. It is used to automate the process of building the project. (aka command shortcuts)

### ./database/postgres.connection.go
Database connection file for PostgreSQL. It is used to connect to the database.

### ./model/common.model.go
Common model file. It contains the common fields e.g. created_at, updated_at, id, etc.

### ./model/order.model.go
Order model file. It contains the order model.

## Running the program
```bash
    # Run the database
    make compose
    # or
    docker-compose up -d

    # Run the program
    go run main.go
```