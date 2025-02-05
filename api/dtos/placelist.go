package dtos

type Placelist struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	AuthorUsername string `json:"author_username"`
}

type PlacelistCreate struct {
	Name string `json:"name"`
}

type PlacelistUpdate struct {
	Name string `json:"name"`
}
