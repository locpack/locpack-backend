package domain

import (
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
) *placeService {
	return &placeService{placeRepository, userRepository}
}

func (s *placeService) GetByPublicIDWithUser(publicID string, userPublicID string) (*models.Place, error) {
	u, err := s.userRepository.GetByPublicID(userPublicID)
	if err != nil {
		return nil, err
	}

	p, err := s.placeRepository.GetByPublicIDWithUser(publicID, u.ID)
	if err != nil {
		return nil, err
	}

	foundPlace := &models.Place{
		ID:      p.PublicID,
		Name:    p.Name,
		Address: p.Address,
		Visited: p.Users[0].Visited,
	}

	return foundPlace, nil
}

func (s *placeService) GetByNameOrAddressWithUser(query string, userPublicID string) (*[]models.Place, error) {
	u, err := s.userRepository.GetByPublicID(userPublicID)
	if err != nil {
		return nil, err
	}

	places, err := s.placeRepository.GetByNameOrAddressWithUser(query, u.ID)
	if err != nil {
		return nil, err
	}

	foundPlaces := []models.Place{}

	for _, p := range *places {
		newPlace := models.Place{
			ID:      p.PublicID,
			Name:    p.Name,
			Address: p.Address,
			Visited: p.Users[0].Visited,
		}
		foundPlaces = append(foundPlaces, newPlace)
	}

	return &foundPlaces, nil
}

func (s *placeService) Create(userPublicID string, pc *models.PlaceCreate) error {
	u, err := s.userRepository.GetByPublicID(userPublicID)
	if err != nil {
		return err
	}

	placeID := rdg.GenerateID()
	place := &entities.Place{
		ID:       placeID,
		Name:     pc.Name,
		PublicID: rdg.GeneratePublicID(),
		Address:  pc.Address,
		Users: []entities.UserPlace{
			{
				ID:      rdg.GenerateID(),
				PlaceID: placeID,
				UserID:  u.ID,
				Visited: pc.Visited,
				Created: true,
			},
		},
	}

	err = s.placeRepository.Create(place)

	return err
}

func (s *placeService) UpdateByPublicIDWithUser(publicID string, userPublicID string, pu *models.PlaceUpdate) error {
	u, err := s.userRepository.GetByPublicID(userPublicID)
	if err != nil {
		return err
	}

	p, err := s.placeRepository.GetByPublicIDWithUser(publicID, u.ID)
	if err != nil {
		return err
	}

	p.Name = pu.Name
	p.Address = pu.Address
	p.Users[0].Visited = pu.Visited

	err = s.placeRepository.Update(p)

	return err
}
