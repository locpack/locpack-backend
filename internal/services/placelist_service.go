package services

import "placelists/api/dtos"

type PlacelistService struct {
}

func NewPlacelistService() PlacelistService {
	return PlacelistService{}
}

func (s *PlacelistService) GetPlacelistsByNameOrAuthor(query string) (*[]dtos.Placelist, error) {
	return nil, nil
}

func (s *PlacelistService) GetPlacelistsFollowedByUsername(username string) (*[]dtos.Placelist, error) {
	return nil, nil
}

func (s *PlacelistService) GetPlacelistsCreatedByUsername(username string) (*[]dtos.Placelist, error) {
	return nil, nil
}

func (s *PlacelistService) GetPlacelistByID(id string) (*dtos.Placelist, error) {
	return nil, nil
}

func (s *PlacelistService) CreatePlacelist(dtos.PlacelistCreate) (*dtos.Placelist, error) {
	return nil, nil
}

func (s *PlacelistService) GetPlacelistPlacesByID(id string) (*[]dtos.Place, error) {
	return nil, nil
}

func (s *PlacelistService) UpdatePlacelistByID(id string, pu dtos.PlacelistUpdate) (*dtos.Placelist, error) {
	return nil, nil
}

func (s *PlacelistService) UpdatePlacelistPlacesByID(id string, places []dtos.Place) (*[]dtos.Place, error) {
	return nil, nil
}

func (s *PlacelistService) RemovePlacelistByID(id string) (*dtos.Placelist, error) {
	return nil, nil
}
