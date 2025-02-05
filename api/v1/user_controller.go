package v1

import (
	"placelists/api"
	"placelists/api/dtos"
	"placelists/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/users")
	g.GET("/my", uc.getUserMy)
	g.GET("/:username", uc.getUserByUsername)
	g.PUT("/:username", uc.putUserByUsername)
}

func (uc *UserController) getUserMy(c *gin.Context) {
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

func (uc *UserController) getUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := uc.userService.GetUserByUsername(username)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, user, []dtos.Error{})
}

func (uc *UserController) putUserByUsername(c *gin.Context) {
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
