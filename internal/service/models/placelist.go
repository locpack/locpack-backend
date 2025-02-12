package models

type PlacelistStatus string

const (
	PlacelistFollowed PlacelistStatus = "FOLLOWED"
	PlacelistCreated  PlacelistStatus = "CREATED"
	PlacelistNone     PlacelistStatus = "NONE"
)

type Placelist struct {
	ID             string
	Name           string
	AuthorID       string
	AuthorUsername string
	Status         PlacelistStatus
}

type PlacelistCreate struct {
	Name string
}

type PlacelistUpdate struct {
	Name      string
	Status    PlacelistStatus
	PlacesIDs []string
}
