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

type placelistControllerImpl struct {
	service service.PlacelistService
}

func NewPlacelistController(service service.PlacelistService) server.PlacelistController {
	return &placelistControllerImpl{service}
}

func (c *placelistControllerImpl) GetPlacelistsByQuery(ctx *gin.Context) {
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

	placelists, err := c.service.GetByNameOrAuthor(query, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistsDTOs := []dtos.Placelist{}
	err = copier.Copy(&placelists, &placelistsDTOs)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placelistsDTOs)
}

func (c *placelistControllerImpl) PostPlacelist(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	var placelistCreateDTO dtos.PlacelistCreate
	err := ctx.ShouldBindJSON(&placelistCreateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistCreate := models.PlacelistCreate{}
	err = copier.Copy(&placelistCreateDTO, &placelistCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelist, err := c.service.Create(userID, placelistCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistDTO := dtos.Placelist{}
	err = copier.Copy(&placelist, &placelistDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placelistDTO)
}

func (c *placelistControllerImpl) GetPlacelistsFollowed(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelists, err := c.service.GetFollowedByUserID(userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistsDTOs := dtos.Placelist{}
	err = copier.Copy(&placelists, &placelistsDTOs)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placelistsDTOs)
}

func (c *placelistControllerImpl) GetPlacelistsCreated(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelists, err := c.service.GetCreatedByUserID(userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistsDTOs := []dtos.Placelist{}
	err = copier.Copy(&placelists, &placelistsDTOs)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placelistsDTOs)
}

func (c *placelistControllerImpl) GetPlacelistByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistID := ctx.Param("id")
	if len(placelistID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelist, err := c.service.GetByID(placelistID, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistDTO := dtos.Placelist{}
	err = copier.Copy(&placelist, &placelistDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placelistDTO)
}

func (c *placelistControllerImpl) PutPlacelistByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistID := ctx.Param("id")
	if len(placelistID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	var placelistUpdateDTO dtos.PlacelistUpdate
	err := ctx.ShouldBindJSON(&placelistUpdateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistUpdate := models.PlacelistUpdate{}
	err = copier.Copy(&placelistUpdateDTO, &placelistUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelist, err := c.service.UpdateByID(placelistID, userID, placelistUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	placelistDTO := dtos.Placelist{}
	err = copier.Copy(&placelist, &placelistDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.ErrorResponse(ctx, http.StatusBadRequest, errors)
		return
	}

	api.SuccessResponse(ctx, http.StatusOK, placelistDTO)
}
