package models

type Place struct {
	ID      string
	Name    string
	Address string
	Visited bool
}

type PlaceCreate struct {
	Name    string
	Address string
}

type PlaceUpdate struct {
	Name    string
	Address string
	Visited bool
}
