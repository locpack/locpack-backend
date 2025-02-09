package models

type Placelist struct {
	ID     string
	Name   string
	Author string
}

type PlacelistCreate struct {
	Name string
}

type PlacelistUpdate struct {
	Name string
}
