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
		api.Response(ctx, http.StatusBadRequest, []dtos.Placelist{}, errors)
		return
	}

	query := ctx.Query("query")
	if len(query) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Placelist{}, errors)
		return
	}

	placelists, err := c.service.GetByNameOrAuthor(query, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Placelist{}, errors)
		return
	}

	placelistsDTOs := []dtos.Placelist{}
	copier.Copy(&placelists, &placelistsDTOs)

	api.Response(ctx, http.StatusOK, placelistsDTOs, []dtos.Error{})
}

func (c *placelistControllerImpl) PostPlacelist(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	var placelistCreateDTO dtos.PlacelistCreate
	err := ctx.ShouldBindJSON(&placelistCreateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelistCreate := models.PlacelistCreate{}
	copier.Copy(&placelistCreateDTO, &placelistCreate)

	placelist, err := c.service.Create(userID, placelistCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelistDTO := dtos.Placelist{}
	copier.Copy(&placelist, &placelistDTO)

	api.Response(ctx, http.StatusOK, placelistDTO, []dtos.Error{})
}

func (c *placelistControllerImpl) GetPlacelistsFollowed(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Placelist{}, errors)
		return
	}

	placelists, err := c.service.GetFollowedByUserID(userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Placelist{}, errors)
		return
	}

	placelistsDTOs := dtos.Placelist{}
	copier.Copy(&placelists, &placelistsDTOs)

	api.Response(ctx, http.StatusOK, placelistsDTOs, []dtos.Error{})
}

func (c *placelistControllerImpl) GetPlacelistsCreated(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Placelist{}, errors)
		return
	}

	placelists, err := c.service.GetCreatedByUserID(userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Placelist{}, errors)
		return
	}

	placelistsDTOs := []dtos.Placelist{}
	copier.Copy(&placelists, &placelistsDTOs)

	api.Response(ctx, http.StatusOK, placelistsDTOs, []dtos.Error{})
}

func (c *placelistControllerImpl) GetPlacelistByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelistID := ctx.Param("id")
	if len(placelistID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelist, err := c.service.GetByID(placelistID, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelistDTO := dtos.Placelist{}
	copier.Copy(&placelist, &placelistDTO)

	api.Response(ctx, http.StatusOK, placelistDTO, []dtos.Error{})
}

func (c *placelistControllerImpl) PutPlacelistByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelistID := ctx.Param("id")
	if len(placelistID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	var placelistUpdateDTO dtos.PlacelistUpdate
	err := ctx.ShouldBindJSON(&placelistUpdateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelistUpdate := models.PlacelistUpdate{}
	copier.Copy(&placelistUpdateDTO, &placelistUpdate)

	placelist, err := c.service.UpdateByID(placelistID, userID, placelistUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placelistDTO := dtos.Placelist{}
	copier.Copy(&placelist, &placelistDTO)

	api.Response(ctx, http.StatusOK, placelistDTO, []dtos.Error{})
}
