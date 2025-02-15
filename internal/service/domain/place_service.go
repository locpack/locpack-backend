package domain

import (
	"placelists/internal/service"
	"placelists/internal/service/models"
	"placelists/internal/storage"
	"placelists/internal/storage/entities"
	"placelists/pkg/rdg"
)

type placeService struct {
	placeRepository storage.PlaceRepository
	userRepository  storage.UserRepository
}

func NewPlaceService(
	placeRepository storage.PlaceRepository,
	userRepository storage.UserRepository,
) service.PlaceService {
	return &placeService{placeRepository, userRepository}
}

func (s *placeService) GetByPublicID(placeID string, userID string) (models.Place, error) {
	place, err := s.placeRepository.GetByPublicIDFull(placeID)
	if err != nil {
		return models.Place{}, err
	}

	visited := false
	for _, visitor := range place.Visitors {
		if visitor.PublicID == userID {
			visited = true
			break
		}
	}

	foundPlace := models.Place{
		ID:      place.PublicID,
		Name:    place.Name,
		Address: place.Address,
		Visited: visited,
	}

	return foundPlace, nil
}

func (s *placeService) GetByNameOrAddress(query string, userID string) ([]models.Place, error) {
	places, err := s.placeRepository.GetByNameOrAddressFull(query)
	if err != nil {
		return []models.Place{}, err
	}

	foundPlaces := []models.Place{}

	for _, place := range places {
		visited := false
		for _, visitor := range place.Visitors {
			if visitor.PublicID == userID {
				visited = true
				break
			}
		}
		newPlace := models.Place{
			ID:      place.PublicID,
			Name:    place.Name,
			Address: place.Address,
			Visited: visited,
		}
		foundPlaces = append(foundPlaces, newPlace)
	}

	return foundPlaces, nil
}

func (s *placeService) Create(userID string, pc models.PlaceCreate) error {
	user, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return err
	}

	visitors := []entities.User{}
	if pc.Visited {
		visitors = append(visitors, user)
	}

	place := entities.Place{
		ID:       rdg.GenerateID(),
		PublicID: rdg.GeneratePublicID(),
		Name:     pc.Name,
		Address:  pc.Address,
		AuthorID: user.ID,
		Visitors: visitors,
	}

	err = s.placeRepository.Create(place)

	return err
}

func (s *placeService) UpdateByPublicID(placeID string, userID string, pu models.PlaceUpdate) error {
	user, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return err
	}

	place, err := s.placeRepository.GetByPublicIDFull(placeID)
	if err != nil {
		return err
	}

	place.Name = pu.Name
	place.Address = pu.Address

	if pu.Visited {
		place.Visitors = append(place.Visitors, user)
	} else {
		for i, visitor := range place.Visitors {
			if visitor.PublicID == userID {
				place.Visitors = append(place.Visitors[:i], place.Visitors[i+1:]...)
			}
		}
	}

	err = s.placeRepository.Update(place)

	return err
}
