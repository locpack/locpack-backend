package entities

type Placelist struct {
	PublicEntity

	Name string `gorm:"not null"`

	Places *[]PlacelistPlace
	Users  *[]UserPlacelist
}
