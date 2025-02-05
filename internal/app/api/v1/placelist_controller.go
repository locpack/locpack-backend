package apiV1

import (
	"placelists/internal/app/api"
	"placelists/internal/app/api/dtos"
	"placelists/internal/core/services"

	"github.com/gin-gonic/gin"
)

type PlacelistController struct {
	placelistService services.PlacelistService
}

func NewPlacelistController(ps services.PlacelistService) *PlacelistController {
	return &PlacelistController{placelistService: ps}
}

func (pc *PlacelistController) RegisterRoutes(r *gin.RouterGroup) {
	placelists := r.Group("/placelists")
	{
		placelists.GET("/", pc.getPlacelistsByQuery)
		placelists.POST("/", pc.postPlacelist)
		placelists.GET("/followed", pc.getPlacelistsFollowed)
		placelists.GET("/created", pc.getPlacelistsCreated)
		placelists.GET("/:id", pc.getPlacelistByID)
		placelists.PUT("/:id", pc.putPlacelistByID)
		placelists.DELETE("/:id", pc.deletePlacelistByID)
		placelists.GET("/:id/places", pc.getPlacelistPlacesByID)
		placelists.PUT("/:id/places", pc.putPlacelistPlacesByID)
	}
}

func (pc *PlacelistController) getPlacelistsByQuery(c *gin.Context) {
	query := c.Query("query")
	if len(query) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Placelist{}, errors)
		return
	}

	placelists, err := pc.placelistService.GetPlacelistsByNameOrAuthor(query)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, placelists, []dtos.Error{})
}

func (pc *PlacelistController) getPlacelistsFollowed(c *gin.Context) {
	username := c.GetString("username")
	if len(username) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Placelist{}, errors)
		return
	}

	placelists, err := pc.placelistService.GetPlacelistsFollowedByUsername(username)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, placelists, []dtos.Error{})
}

func (pc *PlacelistController) getPlacelistsCreated(c *gin.Context) {
	username := c.GetString("username")
	if len(username) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Placelist{}, errors)
		return
	}

	placelists, err := pc.placelistService.GetPlacelistsCreatedByUsername(username)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, placelists, []dtos.Error{})
}

func (pc *PlacelistController) getPlacelistByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	placelist, err := pc.placelistService.GetPlacelistByID(id)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, placelist, []dtos.Error{})
}

func (pc *PlacelistController) postPlacelist(c *gin.Context) {
	var placelistCreate dtos.PlacelistCreate
	err := c.ShouldBindJSON(&placelistCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	placelist, err := pc.placelistService.CreatePlacelist(placelistCreate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, placelist, []dtos.Error{})
}

func (pc *PlacelistController) getPlacelistPlacesByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Place{}, errors)
		return
	}

	places, err := pc.placelistService.GetPlacelistPlacesByID(id)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, places, []dtos.Error{})
}

func (pc *PlacelistController) putPlacelistByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	var placelistUpdate dtos.PlacelistUpdate
	err := c.ShouldBindJSON(&placelistUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	placelist, err := pc.placelistService.UpdatePlacelistByID(id, placelistUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, placelist, []dtos.Error{})
}

func (pc *PlacelistController) putPlacelistPlacesByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Place{}, errors)
		return
	}

	var placesUpdate []dtos.Place
	err := c.ShouldBindJSON(&placesUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Place{}, errors)
		return
	}

	places, err := pc.placelistService.UpdatePlacelistPlacesByID(id, placesUpdate)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, []dtos.Place{}, errors)
		return
	}

	api.Response(c, 200, places, []dtos.Error{})
}

func (pc *PlacelistController) deletePlacelistByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	placelist, err := pc.placelistService.RemovePlacelistByID(id)
	if err != nil {
		errors := []dtos.Error{{Message: "Some error", Code: "000"}}
		api.Response(c, 400, nil, errors)
		return
	}

	api.Response(c, 200, placelist, []dtos.Error{})
}
