package controller

import (
	"net/http"

	"locpack-backend/internal/server"
	"locpack-backend/internal/server/dto"
	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
	"locpack-backend/pkg/adapter"

	"github.com/jinzhu/copier"
)

type placeControllerImpl struct {
	service service.PlaceService
}

func NewPlaceController(service service.PlaceService) server.PlaceController {
	return &placeControllerImpl{service}
}

// GetPlacesByQuery
// @Summary Search places by query
// @Description Get places matching name or address
// @Tags Places
// @Param query query string true "Search query"
// @Success 200 {object} dto.ResponseWrapper{data=[]dto.Place}
// @Failure 400 {object} dto.ResponseWrapper{data=[]dto.Place}
// @Router /api/v1/places [get]
func (c *placeControllerImpl) GetPlacesByQuery(ctx adapter.APIContext) {
	myUserID := ctx.GetString("myUserID")

	query := ctx.Query("query")
	if len(query) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	places, err := c.service.GetByNameOrAddress(query, myUserID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var placesDTOs []dto.Place
	err = copier.Copy(&placesDTOs, &places)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data:   placesDTOs,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// PostPlace
// @Summary Register a new place
// @Description Add a new place to the database
// @Tags Places
// @Security BearerAuth
// @Param place body dto.PlaceCreate true "Place data"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Place}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Place}
// @Router /api/v1/places [post]
func (c *placeControllerImpl) PostPlace(ctx adapter.APIContext) {
	myUserID := ctx.GetString("userID")
	if len(myUserID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var placeCreateDTO dto.PlaceCreate
	err := ctx.ShouldBindJSON(&placeCreateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placeCreate := model.PlaceCreate{}
	err = copier.Copy(&placeCreate, &placeCreateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	place, err := c.service.Create(myUserID, placeCreate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placeDTO := dto.Place{}
	err = copier.Copy(&placeDTO, &place)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data:   placeDTO,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// GetPlaceByID
// @Summary Get place by ID
// @Description Get a specific place by its ID
// @Tags Places
// @Param id path string true "Place ID"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Place}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Place}
// @Router /api/v1/places/{id} [get]
func (c *placeControllerImpl) GetPlaceByID(ctx adapter.APIContext) {
	myUserID := ctx.GetString("userID")

	placeID := ctx.Param("id")
	if len(placeID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	place, err := c.service.GetByID(placeID, myUserID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placeDTO := dto.Place{}
	err = copier.Copy(&placeDTO, &place)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data: placeDTO,
		Meta: dto.Meta{Success: true},
	})
}

// PutPlaceByID
// @Summary Update place by ID
// @Description Update a specific place by its ID
// @Tags Places
// @Security BearerAuth
// @Param id path string true "Place ID"
// @Param place body dto.PlaceUpdate true "Place data"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Place}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Place}
// @Router /api/v1/places/{id} [put]
func (c *placeControllerImpl) PutPlaceByID(ctx adapter.APIContext) {
	myUserID := ctx.GetString("myUserID")
	if len(myUserID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placeID := ctx.Param("id")
	if len(placeID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var placeUpdateDTO dto.PlaceUpdate
	err := ctx.ShouldBindJSON(&placeUpdateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placeUpdate := model.PlaceUpdate{}
	err = copier.Copy(&placeUpdate, &placeUpdateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	place, err := c.service.UpdateByID(placeID, myUserID, placeUpdate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placeDTO := dto.Place{}
	err = copier.Copy(&placeDTO, &place)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrapper{
		Data: placeDTO,
		Meta: dto.Meta{Success: true},
	})
}
