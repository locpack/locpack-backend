package dto

type User struct {
	ID       string `json:"id" copier:"PublicID"`
	Username string `json:"username"`
}

type UserUpdate struct {
	Username string `json:"username"`
}
