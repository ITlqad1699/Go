package controller

import (
	"github.com/anonydev/e-commerce-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

// Return pointer to UserController
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user services
// router -> controller -> service -> repo -> models(struct) -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	return response.ErrorResponse(c, 2003, "Email is invalid")
	// }
	// return response.SuccessResponse(c, 2001, []string{"@nonydev", "user2"})

}
