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

type packControllerImpl struct {
	service service.PackService
}

func NewPackController(service service.PackService) server.PackController {
	return &packControllerImpl{service}
}

// GetPacksByQuery
// @Summary Search packs by query
// @Description Get packs matching name or author
// @Tags Packs
// @Security BearerAuth
// @Param query query string true "Search query"
// @Success 200 {object} dto.ResponseWrapper{data=[]dto.Pack}
// @Failure 400 {object} dto.ResponseWrapper{data=[]dto.Pack}
// @Router /api/v1/packs [get]
func (c *packControllerImpl) GetPacksByQuery(ctx adapter.APIContext) {
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

	packs, err := c.service.GetByNameOrAuthor(query, userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var packsDTOs []dto.Pack
	err = copier.Copy(&packsDTOs, &packs)
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
		Data:   packsDTOs,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// PostPack
// @Summary Create a new pack
// @Description Add a new pack to the database
// @Tags Packs
// @Security BearerAuth
// @Param pack body dto.PackCreate true "Pack data"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Pack}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Pack}
// @Router /api/v1/packs [post]
func (c *packControllerImpl) PostPack(ctx adapter.APIContext) {
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

	var packCreateDTO dto.PackCreate
	err := ctx.ShouldBindJSON(&packCreateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	packCreate := model.PackCreate{}
	err = copier.Copy(&packCreate, &packCreateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	pack, err := c.service.Create(userID, packCreate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	packDTO := dto.Pack{}
	err = copier.Copy(&packDTO, &pack)
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
		Data:   packDTO,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// GetPacksFollowed
// @Summary Get followed packs
// @Description Get packs followed by the current user
// @Tags Packs
// @Security BearerAuth
// @Success 200 {object} dto.ResponseWrapper{data=[]dto.Pack}
// @Failure 400 {object} dto.ResponseWrapper{data=[]dto.Pack}
// @Router /api/v1/packs/followed [get]
func (c *packControllerImpl) GetPacksFollowed(ctx adapter.APIContext) {
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

	packs, err := c.service.GetFollowedByUserID(userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var packsDTOs []dto.Pack
	err = copier.Copy(&packsDTOs, &packs)
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
		Data:   packsDTOs,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// GetPacksCreated
// @Summary Get created packs
// @Description Get packs created by the current user
// @Tags Packs
// @Security BearerAuth
// @Success 200 {object} dto.ResponseWrapper{data=[]dto.Pack}
// @Failure 400 {object} dto.ResponseWrapper{data=[]dto.Pack}
// @Router /api/v1/packs/created [get]
func (c *packControllerImpl) GetPacksCreated(ctx adapter.APIContext) {
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

	packs, err := c.service.GetCreatedByUserID(userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var packsDTOs []dto.Pack
	err = copier.Copy(&packsDTOs, &packs)
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
		Data:   packsDTOs,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// GetPackByID
// @Summary Get pack by ID
// @Description Get a specific pack by its ID
// @Tags Packs
// @Security BearerAuth
// @Param id path string true "Pack ID"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Pack}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Pack}
// @Router /api/v1/packs/{id} [get]
func (c *packControllerImpl) GetPackByID(ctx adapter.APIContext) {
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

	packID := ctx.Param("id")
	if len(packID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	pack, err := c.service.GetByID(packID, userID)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	packDTO := dto.Pack{}
	err = copier.Copy(&packDTO, &pack)
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
		Data:   packDTO,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}

// PutPackByID
// @Summary Update pack by ID
// @Description Update a specific pack by its ID
// @Tags Packs
// @Security BearerAuth
// @Param id path string true "Pack ID"
// @Param pack body dto.PackUpdate true "Pack data"
// @Success 200 {object} dto.ResponseWrapper{data=dto.Pack}
// @Failure 400 {object} dto.ResponseWrapper{data=dto.Pack}
// @Router /api/v1/packs/{id} [put]
func (c *packControllerImpl) PutPackByID(ctx adapter.APIContext) {
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

	packID := ctx.Param("id")
	if len(packID) == 0 {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	var packUpdateDTO dto.PackUpdate
	err := ctx.ShouldBindJSON(&packUpdateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	packUpdate := model.PackUpdate{}
	err = copier.Copy(&packUpdate, &packUpdateDTO)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	pack, err := c.service.UpdateByID(packID, userID, packUpdate)
	if err != nil {
		errors := []dto.Error{{Message: "Some error", Code: "000"}}
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Data:   nil,
			Meta:   dto.Meta{Success: false},
			Errors: errors,
		})
		return
	}

	packDTO := dto.Pack{}
	err = copier.Copy(&packDTO, &pack)
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
		Data:   packDTO,
		Meta:   dto.Meta{Success: true},
		Errors: nil,
	})
}
