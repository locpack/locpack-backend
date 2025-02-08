package entities

type Place struct {
	PublicEntity

	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`

	Placelists *[]PlacelistPlace
	Users      *[]UserPlace
}
