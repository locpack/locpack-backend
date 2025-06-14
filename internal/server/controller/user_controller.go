package controller

import (
	"net/http"

	"github.com/jinzhu/copier"
	"locpack-backend/internal/server"
	"locpack-backend/internal/server/dto"
	"locpack-backend/internal/service"
	"locpack-backend/pkg/adapter"
)

type userControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) server.UserController {
	return &userControllerImpl{service}
}

// GetUserMy
// @Summary Get current user info
// @Description Get information about the currently authenticated user
// @Tags Users
// @Security BearerAuth
// @Success 200 {object} dto.ResponseWrapper{data=dto.User}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.User}
// @Router /api/v1/users/my [get]
func (c *userControllerImpl) GetUserMy(ctx adapter.APIContext) {
	myUserID := ctx.GetString("myUserID")
	if len(myUserID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	user, err := c.service.GetByID(myUserID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	userDTO := dto.User{}
	err = copier.Copy(&userDTO, &user)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data: userDTO,
		Meta: dto.Meta{Success: true},
	})
}

// GetUserByID
// @Summary Get user by ID
// @Description Get information about any user by their ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} dto.ResponseWrapper{data=dto.User}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.User}
// @Router /api/v1/users/{id} [get]
func (c *userControllerImpl) GetUserByID(ctx adapter.APIContext) {
	userID := ctx.Param("id")

	user, err := c.service.GetByID(userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	userDTO := dto.User{}
	err = copier.Copy(&userDTO, &user)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data: userDTO,
		Meta: dto.Meta{Success: true},
	})
}
