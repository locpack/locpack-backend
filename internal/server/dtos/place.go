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
	Visited string `json:"visited"`
}

type PlaceUpdate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Visited bool   `json:"visited"`
}
