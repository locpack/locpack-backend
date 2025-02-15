package controllers

import (
	"net/http"
	"placelists/internal/server"
	"placelists/internal/server/api"
	"placelists/internal/server/dtos"
	"placelists/internal/service"
	"placelists/internal/service/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type userControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) server.UserController {
	return &userControllerImpl{service}
}

// GetUsers return list of all users from the database
// @Summary return list of all
// @Description return list of all users from the database
// @Tags Users
// @Router /users [get]
func (c *userControllerImpl) GetUserMy(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	user, err := c.service.GetByID(userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	userDTO := dtos.User{}
	copier.Copy(&user, &userDTO)

	api.Response(ctx, http.StatusOK, userDTO, []dtos.Error{})
}

func (c *userControllerImpl) GetUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := c.service.GetByID(userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	userDTO := dtos.User{}
	copier.Copy(&user, &userDTO)

	api.Response(ctx, http.StatusOK, userDTO, []dtos.Error{})
}

func (c *userControllerImpl) PutUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	var userUpdateDTO dtos.UserUpdate
	err := ctx.ShouldBindJSON(&userUpdateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	userUpdate := models.UserUpdate{}
	copier.Copy(&userUpdateDTO, &userUpdate)

	user, err := c.service.UpdateByID(userID, userUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	userDTO := dtos.User{}
	copier.Copy(&user, &userDTO)

	api.Response(ctx, http.StatusOK, userDTO, []dtos.Error{})
}
