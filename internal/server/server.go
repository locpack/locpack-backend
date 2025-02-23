package server

import "github.com/gin-gonic/gin"

type Controller struct {
	Place     PlaceController
	Placelist PlacelistController
	User      UserController
}

type PlaceController interface {
	GetPlacesByQuery(ctx *gin.Context)
	PostPlace(ctx *gin.Context)
	GetPlaceByID(ctx *gin.Context)
	PutPlaceByID(ctx *gin.Context)
}

type PlacelistController interface {
	GetPlacelistsByQuery(ctx *gin.Context)
	PostPlacelist(ctx *gin.Context)
	GetPlacelistsFollowed(ctx *gin.Context)
	GetPlacelistsCreated(ctx *gin.Context)
	GetPlacelistByID(ctx *gin.Context)
	PutPlacelistByID(ctx *gin.Context)
}

type UserController interface {
	GetUserMy(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	PutUserByID(ctx *gin.Context)
}
