# 2-dependency-injection

## Pointers
`*` is added to the TYPE to indicate that the variable is a pointer to the type.
`&` is used to get the memory address of the VARIABLE.

To sum up, use `*` with TYPES and `&` with VARIABLES to encourage pass by reference, which is more efficient and faster.

## Dependency Injection
Dependency Injection is a technique in which an object receives other objects that it DEPENDS on. These other objects are called dependencies. Instead of creating the dependencies INSIDE THE OBJECT, the object is given the dependencies by an external entity that configures the dependencies.
```go
type Titan struct {
    height int
}

// we cannot change the height of the Titan
func NewTitan() *Titan {
    return &Titan{height: 10}
}

// With DI, we can change the height of the Titan
func NewTitanDependencyInjection(height int) *Titan {
    return &Titan{height: height}
}
```

Dependency Injection is a technique that allows us to write code that is more modular, testable, and maintainable. It is a way to remove hard-coded dependencies and make it possible to change them easily.

We don't want to store orders as a map because when the program is restarted, the orders will be lost. We want to store the orders in a database. We can use an interface to abstract the database operations and use dependency injection to inject that interface into the service. We call interfaces that are used for database/stores operations "repositories".

## New files
### ./repository/cart.repository.go
Repository file for cart. It is the interface that abstracts the database operations.