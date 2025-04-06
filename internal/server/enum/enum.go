package enum

import "placelists-back/internal/server/types"

const (
	PlacelistFollowed types.PlacelistStatus = iota
	PlacelistCreated
	PlacelistNone
)
