package model

type User struct {
	ID       string `copier:"PublicID"`
	Username string
}

type UserUpdate struct {
	Username string
}
