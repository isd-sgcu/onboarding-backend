# 1-struct-interface

```bash
# Run the program
go run main.go
```

In Go, struct, interface, or method with names that begin with a lowercase letter are private to the package. If you want to make them public, you need to start the name with an uppercase letter.

Packages in Go are like folders in JS. They are used to organize the code. The package name should be the same as the folder name.

## New files
### main.go
Entry point of the program.

### go.mod
Go module file. It is like package.json in JS

### ./service/cart.service.go
Service file for cart.