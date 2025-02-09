package domain

import (
	"placelists/internal/service/models"
	"placelists/internal/storage"
	"placelists/internal/storage/entities"
	"placelists/pkg/rdg"
)

type placelistService struct {
	placelistRepository storage.PlacelistRepository
	userRepository      storage.UserRepository
}

func NewPlacelistService(
	placelistRepository storage.PlacelistRepository,
	userRepository storage.UserRepository,
) *placelistService {
	return &placelistService{placelistRepository, userRepository}
}

func (s *placelistService) GetByNameOrUsernameWithUser(query string, userPublicID string) (*[]models.Placelist, error) {
	u, err := s.userRepository.GetByPublicID(userPublicID)
	if err != nil {
		return nil, err
	}

	s.placelistRepository.GetByNameOrAuthorWithUser(query, u.ID)
	if err != nil {
		return nil, err
	}

	foundPlacelists := []models.Placelist{}

	// for _, p := range *placelists {
	// 	newPlacelist := models.Placelist{
	// 		ID:      p.PublicID,
	// 		Name:    p.Name,
	// 		Author: p.Users[0].,
	// 		Visited: p.Users[0].Visited,
	// 	}
	// 	foundPlacelists = append(foundPlacelists, newPlacelist)
	// }

	return &foundPlacelists, nil
}

func (s *placelistService) Create(userPublicID string, pc *models.PlacelistCreate) error {
	u, err := s.userRepository.GetByPublicID(userPublicID)
	if err != nil {
		return err
	}

	placelistID := rdg.GenerateID()
	placelist := &entities.Placelist{
		ID:       placelistID,
		Name:     pc.Name,
		PublicID: rdg.GeneratePublicID(),
		Users: []entities.UserPlacelist{
			{
				ID:          rdg.GenerateID(),
				PlacelistID: placelistID,
				UserID:      u.ID,
				Status:      entities.Created,
			},
		},
	}

	err = s.placelistRepository.Create(placelist)

	return err
}
