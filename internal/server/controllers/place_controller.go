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
		api.Response(ctx, http.StatusBadRequest, []dtos.Place{}, errors)
		return
	}

	query := ctx.Query("query")
	if len(query) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Place{}, errors)
		return
	}

	places, err := c.service.GetByNameOrAddress(query, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, []dtos.Place{}, errors)
		return
	}

	placesDTOs := []dtos.Place{}
	copier.Copy(&places, &placesDTOs)

	api.Response(ctx, http.StatusOK, placesDTOs, []dtos.Error{})
}

func (c *placeControllerImpl) PostPlace(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	var placeCreateDTO dtos.PlaceCreate
	err := ctx.ShouldBindJSON(&placeCreateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placeCreate := models.PlaceCreate{}
	copier.Copy(&placeCreateDTO, &placeCreate)

	place, err := c.service.Create(userID, placeCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placeDTO := dtos.Place{}
	copier.Copy(&place, &placeDTO)

	api.Response(ctx, http.StatusOK, placeDTO, []dtos.Error{})
}

func (c *placeControllerImpl) GetPlaceByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placeID := ctx.Param("id")
	if len(placeID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	place, err := c.service.GetByID(placeID, userID)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placeDTO := dtos.Place{}
	copier.Copy(&place, &placeDTO)

	api.Response(ctx, http.StatusOK, placeDTO, []dtos.Error{})
}

func (c *placeControllerImpl) PutPlaceByID(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placeID := ctx.Param("id")
	if len(placeID) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	var placeUpdateDTO dtos.PlaceUpdate
	err := ctx.ShouldBindJSON(&placeUpdateDTO)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placeUpdate := models.PlaceUpdate{}
	copier.Copy(&placeUpdateDTO, &placeUpdate)

	place, err := c.service.UpdateByID(placeID, userID, placeUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(ctx, http.StatusBadRequest, nil, errors)
		return
	}

	placeDTO := dtos.Place{}
	copier.Copy(&place, &placeDTO)

	api.Response(ctx, http.StatusOK, placeDTO, []dtos.Error{})
}
