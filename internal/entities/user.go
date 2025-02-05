package entities

type User struct {
	ID       ID     `json:"id"`
	Username string `json:"username"`
}

type UserPlacelist struct {
	ID          ID              `json:"id"`
	Status      PlacelistStatus `json:"status"`
	UserID      ID              `json:"user_id"`
	PlacelistID ID              `json:"placelist_id"`
}

type UserCreatedPlace struct {
	ID      ID   `json:"id"`
	Visited bool `json:"visited"`
	UserID  ID   `json:"user_id"`
	PlaceID ID   `json:"place_id"`
}
