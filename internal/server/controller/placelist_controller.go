package controller

import (
	"net/http"

	"github.com/jinzhu/copier"
	"placelists-back/internal/server"
	"placelists-back/internal/server/dto"
	"placelists-back/internal/service"
	"placelists-back/internal/service/model"
	"placelists-back/pkg/adapter"
)

type placelistControllerImpl struct {
	service service.PlacelistService
}

func NewPlacelistController(service service.PlacelistService) server.PlacelistController {
	return &placelistControllerImpl{service}
}

// GetPlacelistsByQuery
// @Summary Search placelists by query
// @Description Get placelists matching name or author
// @Tags placelists
// @Security BearerAuth
// @Param query query string true "Search query"
// @Success 200 {object} dto.ResponseWrapper{data=[]dto.Placelist}
// @Failure 400 {object} dto.ResponseWrapper{data=[]dto.Placelist}
// @Router /api/v1/placelists [get]
func (c *placelistControllerImpl) GetPlacelistsByQuery(ctx adapter.APIContext) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

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

	placelists, err := c.service.GetByNameOrAuthor(query, userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var placelistsDTOs []dto.Placelist
	err = copier.Copy(&placelists, &placelistsDTOs)
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
		Data:   placelistsDTOs,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// PostPlacelist
// @Summary Create a new placelist
// @Description Add a new placelist to the database
// @Tags placelists
// @Security BearerAuth
// @Param placelist body dto.PlacelistCreate true "Placelist data"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Placelist}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Placelist}
// @Router /api/v1/placelists [post]
func (c *placelistControllerImpl) PostPlacelist(ctx adapter.APIContext) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var placelistCreateDTO dto.PlacelistCreate
	err := ctx.ShouldBindJSON(&placelistCreateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistCreate := model.PlacelistCreate{}
	err = copier.Copy(&placelistCreateDTO, &placelistCreate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelist, err := c.service.Create(userID, placelistCreate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistDTO := dto.Placelist{}
	err = copier.Copy(&placelist, &placelistDTO)
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
		Data:   placelistDTO,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// GetPlacelistsFollowed
// @Summary Get followed placelists
// @Description Get placelists followed by the current user
// @Tags placelists
// @Security BearerAuth
// @Success 200 {object} dto.ResponseWrapper{data=[]dto.Placelist}
// @Failure 400 {object} dto.ResponseWrapper{data=[]dto.Placelist}
// @Router /api/v1/placelists/followed [get]
func (c *placelistControllerImpl) GetPlacelistsFollowed(ctx adapter.APIContext) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelists, err := c.service.GetFollowedByUserID(userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistsDTOs := dto.Placelist{}
	err = copier.Copy(&placelists, &placelistsDTOs)
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
		Data:   placelistsDTOs,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// GetPlacelistsCreated
// @Summary Get created placelists
// @Description Get placelists created by the current user
// @Tags placelists
// @Security BearerAuth
// @Success 200 {object} dto.ResponseWrapper{data=[]dto.Placelist}
// @Failure 400 {object} dto.ResponseWrapper{data=[]dto.Placelist}
// @Router /api/v1/placelists/created [get]
func (c *placelistControllerImpl) GetPlacelistsCreated(ctx adapter.APIContext) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelists, err := c.service.GetCreatedByUserID(userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var placelistsDTOs []dto.Placelist
	err = copier.Copy(&placelists, &placelistsDTOs)
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
		Data:   placelistsDTOs,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// GetPlacelistByID
// @Summary Get placelist by ID
// @Description Get a specific placelist by its ID
// @Tags placelists
// @Security BearerAuth
// @Param id path string true "Placelist ID"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Placelist}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Placelist}
// @Router /api/v1/placelists/{id} [get]
func (c *placelistControllerImpl) GetPlacelistByID(ctx adapter.APIContext) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistID := ctx.Param("id")
	if len(placelistID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelist, err := c.service.GetByID(placelistID, userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistDTO := dto.Placelist{}
	err = copier.Copy(&placelist, &placelistDTO)
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
		Data:   placelistDTO,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// PutPlacelistByID
// @Summary Update placelist by ID
// @Description Update a specific placelist by its ID
// @Tags placelists
// @Security BearerAuth
// @Param id path string true "Placelist ID"
// @Param placelist body dto.PlacelistUpdate true "Placelist data"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Placelist}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Placelist}
// @Router /api/v1/placelists/{id} [put]
func (c *placelistControllerImpl) PutPlacelistByID(ctx adapter.APIContext) {
	userID := ctx.GetString("userID")
	if len(userID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistID := ctx.Param("id")
	if len(placelistID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var placelistUpdateDTO dto.PlacelistUpdate
	err := ctx.ShouldBindJSON(&placelistUpdateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistUpdate := model.PlacelistUpdate{}
	err = copier.Copy(&placelistUpdateDTO, &placelistUpdate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelist, err := c.service.UpdateByID(placelistID, userID, placelistUpdate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	placelistDTO := dto.Placelist{}
	err = copier.Copy(&placelist, &placelistDTO)
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
		Data:   placelistDTO,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}
