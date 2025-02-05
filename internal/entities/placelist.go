package entities

type Placelist struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

type PlacelistPlace struct {
	ID          ID `json:"id"`
	PlaceID     ID `json:"place_id"`
	PlacelistID ID `json:"placelist_id"`
}
