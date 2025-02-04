package apiV1

import (
	"placelists/internal/app/api"
	"placelists/internal/app/api/dtos"
	"placelists/internal/core/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) GetUserMy(c *gin.Context) {
	username := c.GetString("username")
	if len(username) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	user, err := uc.userService.GetUserByUsername(username)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, user, []dtos.Error{})
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := uc.userService.GetUserByUsername(username)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, user, []dtos.Error{})
}

func (uc *UserController) PutUserByUsername(c *gin.Context) {
	username := c.Param("username")

	var userUpdate dtos.UserUpdate
	err := c.ShouldBindJSON(&userUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	user, err := uc.userService.UpdateUserByUsername(username, userUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, user, []dtos.Error{})
}
