package controller

import (
	"net/http"

	"github.com/jinzhu/copier"
	"locpack-backend/internal/server"
	"locpack-backend/internal/server/dto"
	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
	"locpack-backend/pkg/adapter"
)

type authControllerImpl struct {
	service service.AuthService
}

func NewAuthController(authService service.AuthService) server.AuthController {
	return &authControllerImpl{authService}
}

// Register
// @Summary User registration
// @Description Register new user account
// @Tags Auth
// @Param register body dto.Register true "Registration details"
// @Success 200 {object} dto.ResponseWrapper{data=dto.AccessToken}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.AccessToken}
// @Router /api/v1/auth/register [post]
func (c *authControllerImpl) Register(ctx adapter.APIContext) {
	var registerDTO dto.Register
	err := ctx.ShouldBindJSON(&registerDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	register := model.Register{}
	err = copier.Copy(&register, &registerDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	accessToken, err := c.service.Register(register)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	accessTokenDTO := dto.AccessToken{}
	err = copier.Copy(&accessTokenDTO, &accessToken)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data: accessTokenDTO,
		Meta: dto.Meta{Success: true},
	})
}

// Login
// @Summary User login
// @Description Authenticate user and return token
// @Tags Auth
// @Param login body dto.Login true "Login details"
// @Success 200 {object} dto.ResponseWrapper{data=dto.AccessToken}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.AccessToken}
// @Router /api/v1/auth/login [post]
func (c *authControllerImpl) Login(ctx adapter.APIContext) {
	var loginDTO dto.Login
	err := ctx.ShouldBindJSON(&loginDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	login := model.Login{}
	err = copier.Copy(&login, &loginDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	accessToken, err := c.service.Login(login)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	accessTokenDTO := dto.AccessToken{}
	err = copier.Copy(&accessTokenDTO, &accessToken)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data: accessTokenDTO,
		Meta: dto.Meta{Success: true},
	})
}

// Refresh
// @Summary User refresh token
// @Description Refresh user token and return it
// @Tags Auth
// @Param refresh body dto.Refresh true "Refresh details"
// @Success 200 {object} dto.ResponseWrapper{data=dto.AccessToken}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.AccessToken}
// @Router /api/v1/auth/refresh [post]
func (c *authControllerImpl) Refresh(ctx adapter.APIContext) {
	var refreshDTO dto.Refresh
	err := ctx.ShouldBindJSON(&refreshDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	refresh := model.Refresh{}
	err = copier.Copy(&refresh, &refreshDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	accessToken, err := c.service.Refresh(refresh)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	accessTokenDTO := dto.AccessToken{}
	err = copier.Copy(&accessTokenDTO, &accessToken)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data: accessTokenDTO,
		Meta: dto.Meta{Success: true},
	})
}
