package entities

type User struct {
	PublicEntity

	Username string `gorm:"unique;not null"`

	Placelists *[]UserPlacelist
	Places     *[]UserPlace
}
