package fakes

import (
	"errors"
	"placelists/internal/storage/entities"
	"strings"
)

type PlaceRepositoryFakeImpl struct {
	Places []entities.Place
}

func NewFakePlaceRepository() *PlaceRepositoryFakeImpl {
	return &PlaceRepositoryFakeImpl{}
}

func (r *PlaceRepositoryFakeImpl) GetByPublicID(placeID string) (*entities.Place, error) {
	for _, place := range r.Places {
		if place.PublicID == placeID {
			return &place, nil
		}
	}
	return nil, errors.New("place not found")
}

func (r *PlaceRepositoryFakeImpl) GetByPublicIDFull(placeID string) (*entities.Place, error) {
	places, err := r.GetByPublicID(placeID)
	return places, err
}

func (r *PlaceRepositoryFakeImpl) GetByNameOrAddress(query string) (*[]entities.Place, error) {
	var results []entities.Place
	for _, place := range r.Places {
		if strings.Contains(strings.ToLower(place.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(place.Address), strings.ToLower(query)) {
			results = append(results, place)
		}
	}
	if len(results) == 0 {
		return nil, errors.New("no places found")
	}
	return &results, nil
}

func (r *PlaceRepositoryFakeImpl) GetByNameOrAddressFull(query string) (*[]entities.Place, error) {
	places, err := r.GetByNameOrAddress(query)
	return places, err
}

func (r *PlaceRepositoryFakeImpl) Create(p *entities.Place) error {
	for _, place := range r.Places {
		if place.ID == p.ID {
			return errors.New("place already exists")
		}
	}
	r.Places = append(r.Places, *p)
	return nil
}

func (r *PlaceRepositoryFakeImpl) Update(p *entities.Place) error {
	for i, place := range r.Places {
		if place.ID == p.ID {
			r.Places[i] = *p
			return nil
		}
	}
	return errors.New("place not found")
}
