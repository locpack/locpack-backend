package entities

type Place struct {
	ID      ID     `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
