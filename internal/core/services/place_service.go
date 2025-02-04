package services

import "placelists/internal/app/api/dtos"

type PlaceService struct {
}

func NewPlaceService() PlaceService {
	return PlaceService{}
}

func (s *PlaceService) GetPlacesByNameOrAddress(query string) (*[]dtos.Place, error) {
	return &[]dtos.Place{}, nil
}

func (s *PlaceService) GetPlaceByID(id string) (*dtos.Place, error) {
	return nil, nil
}

func (s *PlaceService) CreatePlace(pc dtos.PlaceCreate) (*dtos.Place, error) {
	return nil, nil
}

func (s *PlaceService) UpdatePlaceByID(id string, pu dtos.PlaceUpdate) (*dtos.Place, error) {
	return nil, nil
}
