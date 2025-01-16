# Authors [@nonydev](https://github.com/ITlqad1699)

Today we will deep dive into GO language as well as building API of a standard GO project.

# Go-Ecommerce-API
## A. Folder Structure
```
├───cmd			    # CMD package: contains cli, job, server (main)  
│   ├───cli                 # 
│   ├───job                 #
│   └───server              #
├───configs                 # Configs: Contains config file such as common
├───docs                    # Docs for DEV
├───global                  # 
├───internal                # Main business of your application
│   ├───controller          #
│   ├───initialize          #
│   ├───middleware          #
│   ├───models              #
│   ├───repo                #
│   ├───router              # 
│   └───service             # 
├───migrations              # Database: init, access layer
├───pkg                     # Package of system
│   ├───logger              #
│   ├───setting             #
│   └───utils               #
├───response                #
├───scripts                 #
├───tests                   #
├───third-party             #
└───go.mod                  #
```
## B. GIN
### 1. GIN - Definition:
- Gin is a web framework written in Go. It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
### 2. GIN - Installation:
```shell
<!-- import GIN -->
import "github.com/gin-gonic/gin"
<!-- Get GIN engine -->
go get -u github.com/gin-gonic/gin
```
### 3. GIN - Router:
**Routers:**
```go
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// /v1/2025
	v1 := r.Group("/v1/2025")
	{

		v1.GET("/user/1", c.NewUserController().GetUserByID)
		v1.GET("/ping", c.NewPongController().Pong)
	}

	// /v2/2025
	v2 := r.Group("/v2/2025")
	{
		v2.GET("/ping", temp_api)
		v2.PUT("/ping", temp_api)
		v2.PATCH("/ping", temp_api)
		v2.DELETE("/ping", temp_api)
		v2.POST("/ping", temp_api)
		v2.HEAD("/ping", temp_api)
		v2.OPTIONS("/ping", temp_api)
	}
	return r
}
```

**Controllers:**
```go
type PongController struct{}

// Return pointer to PongController
func NewPongController() *PongController {
	return &PongController{}
}

// uc user controller
// us user services
func (uc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	nameStart := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...`pong" + nameStart,
		"default": name,
		"uid":     uid,
		"user":    []string{"@nonydev", "@nony", "daLQA"},
	})
}
```
### 4. GIN - Error Handler:
**httpStatusCode.go:**
```go
package response

const (
	ErrorCodeSuccess      = 2001 // Success
	ErrorCodeParamInvalid = 2003 // Email is invalid
	ErrorInvalidToken     = 3001 // invalid token
)

// messages
var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorInvalidToken:     "invalid token",
}
```
**response.go:**
 ```go
type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // message
	Data    interface{} `json:"data"`    // data returned
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
```
### 5. GIN - Log Handler:
#### 5.1 Installation:
```shell
go get -u go.uber.org/zap
```
#### 5.2 Implementation:
**Example:**
```go
import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Sugar is a logger that wraps the Logger to provide a more ergonomic, but slightly slower, API.
	// sugar := zap.NewExample().Sugar()
	// sugar.Info("Hello name: %s, age: %d", "@nonydev", 25)
	// // Logger is a fast, structured, leveled logger.
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "@nonydev"), zap.Int("age", 25))

	// logger := zap.NewExample()
	// logger.Info("hello")

	// // Development and Production are two pre-configured loggers that are optimized for different environments.
	// logger, _ = zap.NewDevelopment()
	// logger.Info("hello new Developement")

	// logger, _ = zap.NewProduction()
	// logger.Info("hello new Production")

	// 3. Custom Configuration
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// Format logger
func getEncoderLog() zapcore.Encoder {
	// 1736170682.0589921 -> 2025-01-06T20:38:02.058+0700
	encodeConfiger := zap.NewProductionEncoderConfig()
	encodeConfiger.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfiger.TimeKey = "time"
	// lvl -> level
	encodeConfiger.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller": "main.log.go:line-number"
	encodeConfiger.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfiger)
}

func getWriterSync() zapcore.WriteSyncer {
	// This syntax is a shorthand for openning a file with the default permissions.
	// name: file name
	// flag: os.O_CREATE|os.O_APPEND|os.O_WRONLY|...
	// perm: 0666|0755
	os.MkdirAll("log", os.ModePerm)
	file, _ := os.OpenFile("log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

```
### 6. GIN - Viper Cinfig:
#### 6.1 Installation:
- Github repo:  [viper](https://github.com/spf13/viper)
```shell
go get github.com/spf13/viper
```
#### 6.2 Implementation: 
```go
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to look for the config file in
	viper.SetConfigName("local")      // file name
	viper.SetConfigType("yaml")       // file type

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err) // Terminate the program
	}

	fmt.Println("port:", viper.GetInt("server.port"))
	fmt.Println("jwt key:", viper.GetInt("security.jwt.key"))

	// configuration structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	fmt.Println("port:", config.Server.Port)

	// Syntax _, db -> ignore the index
	for _, db := range config.Databases {
		fmt.Println("user:", db.User)
		fmt.Println("password:", db.Password)
		fmt.Println("host:", db.Host)
	}
}

```
### 7. GIN - Middleware:
When building an application, a middleware is a code that hooks on the server-based request/response lifecycle, which will then chain the request from the client to the next middleware function and eventually the last function.
This article thoroughly explores basic middleware chaining/multiple middlewares.

#### What is Middleware?
A middleware is an http.handler that wraps another http.handler in a server request/response processing. Middleware can be defined in many components; It sits between the web server and the actual handler. Whenever a handler is defined for a URL pattern, the request hits the handler and executes the business logic. All middleware process these functions:
- Process the request from the client before hitting the handler(authentication) - interceptor
- Process the handler function
- Process the response for the client
- logging.. and so on

In an application with no middleware, if a client sends a request, the request reaches the server and is handled by some function handler, and it is sent back immediately from the server to the client. But in an application with middleware, the request made by the client passes through stages like logging, authenticating, session validation, and so on, then process the business logic. It filters wrong requests from interacting with the business logic.

#### Example
**auth.middleware.go**
```go
package middleware

import (...)

// gin.HandlerFunc => Presents a closure functio 
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Todo: Implement the middleware
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorInvalidToken, "invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
```
**router.go**
```go
func NewRouter() *gin.Engine {
	r := gin.Default()
	// Call middleware authentication
    r.Use(middleware.AuthMiddleware())

	// /v1/2025
	v1 := r.Group("/v1/2025")
	{
		v1.GET("/user/1", c.NewUserController().GetUserByID)
		v1.GET("/ping", c.NewPongController().Pong)
	}
	return r
}

func temp_api(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "temp api",
	})
}
```

### 8. GIN - Unit Test:
#### Fundamental:
- Define test function in package test (module test)
- Function has to start with Test***() Prefix
- Import "testing" package of Go
- Repository: [testify](https://github.com/stretchr/testify)

#### Implementation:
```go
package basic

import (...)

// Test Assert
func TestAssert(t *testing.T) {
	assert.Equal(t, AddOne(1), 2, "Expected 2 but got 2")
    fmt.Println("executing")
}

// Test Require
func TestRequire(t *testing.T) {
	require.Equal(t, AddOne(1), 2, "Expected 2 but got 2")
    fmt.Println("not executing")
}

func AddOne(x int) int {
	return x + 1
}

```
#### Execution:
**Run command:**
```shell
lqad1@DaLQA MINGW64 /d/Projects/Go/github.com/anonydev/e-commerce-api/tests/basic (main)
$ go test -v
=== RUN   TestAssert
executing
--- PASS: TestAssert (0.00s)
=== RUN   TestRequire
not executing
--- PASS: TestRequire (0.00s)
PASS
ok      github.com/anonydev/e-commerce-api/tests/basic  0.419s
```

**Code Coverage:**
```shell
// Gen coverage to file coverage.out
go test -coverprofile=coverage.out

// View in HTML
go tool cover -html=coverage.out -o coverage.html

// Open file HTML
start coverage.html
start chrome <absolute path file>
start microsoft-edge:<absolute path file>
start firefox <absolute path file>
```