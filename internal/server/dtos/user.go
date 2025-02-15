package dtos

type User struct {
	ID       string
	Username string `json:"username"`
}

type UserUpdate struct {
	Username string `json:"username"`
}
