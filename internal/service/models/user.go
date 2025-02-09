package models

type User struct {
	ID       string
	Username string
}

type UserUpdate struct {
	Username string
}
