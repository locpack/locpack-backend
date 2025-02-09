package models

type Placelist struct {
	ID             string
	Name           string
	AuthorUsername string
}

type PlacelistCreate struct {
	Name string
}

type PlacelistUpdate struct {
	Name string
}
