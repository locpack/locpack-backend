package dto

import (
	"placelists-back/internal/server/types"
)

type Meta struct {
	Success bool `json:"success"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type ResponseWrapper struct {
	Data   any     `json:"data"`
	Meta   Meta    `json:"meta"`
	Errors []Error `json:"errors"`
}

type Place struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Visited bool   `json:"visited"`
}

type PlaceCreate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Visited string `json:"visited"`
}

type PlaceUpdate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Visited bool   `json:"visited"`
}

type Placelist struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	AuthorID       string                `json:"author_id"`
	AuthorUsername string                `json:"author_username"`
	Status         types.PlacelistStatus `json:"status"`
}

type PlacelistCreate struct {
	Name string `json:"name"`
}

type PlacelistUpdate struct {
	Name      string                `json:"name"`
	Status    types.PlacelistStatus `json:"status"`
	PlacesIDs []string              `json:"places_ids"`
}

type User struct {
	ID       string
	Username string `json:"username"`
}

type UserUpdate struct {
	Username string `json:"username"`
}
