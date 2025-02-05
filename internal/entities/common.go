package entities

type ID = string

type PlacelistStatus = string

const (
	Created  PlacelistStatus = "CREATED"
	Followed PlacelistStatus = "FOLLOWED"
)
