package server

import (
	"locpack-backend/pkg/adapter"
)

type PlaceController interface {
	GetPlacesByQuery(ctx adapter.APIContext)
	PostPlace(ctx adapter.APIContext)
	GetPlaceByID(ctx adapter.APIContext)
	PutPlaceByID(ctx adapter.APIContext)
}

type PackController interface {
	GetPacksByQuery(ctx adapter.APIContext)
	PostPack(ctx adapter.APIContext)
	GetPacksFollowed(ctx adapter.APIContext)
	GetPacksCreated(ctx adapter.APIContext)
	GetPackByID(ctx adapter.APIContext)
	PutPackByID(ctx adapter.APIContext)
}

type UserController interface {
	GetUserMy(ctx adapter.APIContext)
	GetUserByID(ctx adapter.APIContext)
	PutUserByID(ctx adapter.APIContext)
}
