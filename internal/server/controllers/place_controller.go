package controllers

import (
	"net/http"
	"placelists/internal/server"
	"placelists/internal/server/dtos"
	"placelists/internal/service"
	"placelists/internal/service/models"
	"placelists/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type placeControllerImpl struct {
	service service.PlaceService
}

func NewPlaceController(service service.PlaceService) server.PlaceController {
	return &placeControllerImpl{service}
}

func (c *placeControllerImpl) GetPlacesByQuery(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	query := ctx.Query("query")
	if len(query) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	places, err := c.service.GetByNameOrAddress(query, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placesDTOs := []dtos.Place{}
	err = copier.Copy(&places, &placesDTOs)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placesDTOs)
}

func (c *placeControllerImpl) PostPlace(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	var placeCreateDTO dtos.PlaceCreate
	err := ctx.ShouldBindJSON(&placeCreateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placeCreate := models.PlaceCreate{}
	err = copier.Copy(&placeCreateDTO, &placeCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	place, err := c.service.Create(userID, placeCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placeDTO := dtos.Place{}
	err = copier.Copy(&place, &placeDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placeDTO)
}

func (c *placeControllerImpl) GetPlaceByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placeID := ctx.Param("id")
	if len(placeID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	place, err := c.service.GetByID(placeID, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placeDTO := dtos.Place{}
	err = copier.Copy(&place, &placeDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placeDTO)
}

func (c *placeControllerImpl) PutPlaceByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placeID := ctx.Param("id")
	if len(placeID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	var placeUpdateDTO dtos.PlaceUpdate
	err := ctx.ShouldBindJSON(&placeUpdateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placeUpdate := models.PlaceUpdate{}
	err = copier.Copy(&placeUpdateDTO, &placeUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	place, err := c.service.UpdateByID(placeID, userID, placeUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placeDTO := dtos.Place{}
	err = copier.Copy(&place, &placeDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placeDTO)
}
