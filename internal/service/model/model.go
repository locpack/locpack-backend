package model

import "placelists-back/internal/service/types"

type Place struct {
	ID      string `copier:"PublicID"`
	Name    string
	Address string
	Visited bool
}

type PlaceCreate struct {
	Name    string
	Address string
	Visited bool
}

type PlaceUpdate struct {
	Name    string
	Address string
	Visited bool
}

type Placelist struct {
	ID             string `copier:"PublicID"`
	Name           string
	AuthorID       string
	AuthorUsername string
	Status         types.PlacelistStatus
}

type PlacelistCreate struct {
	Name string
}

type PlacelistUpdate struct {
	Name      string
	Status    types.PlacelistStatus
	PlacesIDs []string
}

type User struct {
	ID       string `copier:"PublicID"`
	Username string
}

type UserUpdate struct {
	Username string
}
