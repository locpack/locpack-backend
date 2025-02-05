package dtos

type Place struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Visited bool   `json:"visited"`
}

type PlaceCreate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type PlaceUpdate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Visited bool   `json:"visited"`
}
