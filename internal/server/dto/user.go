package dto

type User struct {
	ID       string `json:"id" copier:"PublicID"`
	Username string `json:"username"`
}
