package api

import (
	"placelists/api/dtos"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, statusCode int, data any, errors []dtos.Error) {
	c.JSON(statusCode, dtos.ResponseWrapper[any]{
		Data:   data,
		Meta:   dtos.Meta{Success: len(errors) == 0},
		Errors: errors,
	})
}
