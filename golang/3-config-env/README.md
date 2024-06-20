# 3-config-env

install godotenv (aka dotenv for go) package
```bash
go get github.com/joho/godotenv
```
after installing, there will be a new file `go.sum`, which is similar to `package-lock.json` in JS.

## New files
### /config/config.go
App config, which reads the environment variables.

### /.env
Environment variables file. It is used to store sensitive information like API keys, database passwords, etc.

### /.env.template
Template for the `.env` file. It is used to show the required environment variables. You commit this to remote repositories. In this tutorial, we will commit the .env file so you can see, but please don't commit it in real projects. Put it in `.gitignore`.

### /.gitignore
choose files and directories from being committed to the repository.


```go
    conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}
```
In golang, you will see the `if err != nil` so often. It is a common pattern in Go to return an error as the last return value. If the function is successful, it will return conf as config and error as `nil`. If it fails, it will return conf as `nil` and error as `error`. You should always check the error value.