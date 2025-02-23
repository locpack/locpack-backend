package dtos

type PlacelistStatus string

const (
	PlacelistFollowed PlacelistStatus = "FOLLOWED"
	PlacelistCreated  PlacelistStatus = "CREATED"
	PlacelistNone     PlacelistStatus = "NONE"
)

type Placelist struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	AuthorID       string          `json:"author_id"`
	AuthorUsername string          `json:"author_username"`
	Status         PlacelistStatus `json:"status"`
}

type PlacelistCreate struct {
	Name string `json:"name"`
}

type PlacelistUpdate struct {
	Name      string          `json:"name"`
	Status    PlacelistStatus `json:"status"`
	PlacesIDs []string        `json:"places_ids"`
}
