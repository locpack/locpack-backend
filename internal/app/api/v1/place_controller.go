package apiV1

import (
	"placelists/internal/app/api"
	"placelists/internal/app/api/dtos"
	"placelists/internal/core/services"

	"github.com/gin-gonic/gin"
)

type PlaceController struct {
	placeService services.PlaceService
}

func NewPlaceController(ps services.PlaceService) *PlaceController {
	return &PlaceController{placeService: ps}
}

func (pc *PlaceController) GetPlacesByQuery(c *gin.Context) {
	query := c.Query("query")
	if len(query) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Place{}, errors)
		return
	}

	places, err := pc.placeService.GetPlacesByNameOrAddress(query)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, places, []dtos.Error{})
}

func (pc *PlaceController) GetPlaceByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	place, err := pc.placeService.GetPlaceByID(id)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, place, []dtos.Error{})
}

func (pc *PlaceController) PostPlace(c *gin.Context) {
	var placeCreate dtos.PlaceCreate
	err := c.ShouldBindJSON(&placeCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	place, err := pc.placeService.CreatePlace(placeCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, place, []dtos.Error{})
}

func (pc *PlaceController) PutPlaceByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	var placeUpdate dtos.PlaceUpdate
	err := c.ShouldBindJSON(&placeUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	place, err := pc.placeService.UpdatePlaceByID(id, placeUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, place, []dtos.Error{})
}
