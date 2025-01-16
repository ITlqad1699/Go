
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
#### 2.1 GIN - Get engine:
```shell
<!-- import GIN -->
import "github.com/gin-gonic/gin"
<!-- Get GIN engine -->
go get -u github.com/gin-gonic/gin
```

### 2. GIN - Installation:
