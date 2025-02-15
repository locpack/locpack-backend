package domain

import (
	"errors"
	"placelists/internal/service"
	"placelists/internal/service/models"
	"placelists/internal/storage"
	"placelists/internal/storage/entities"
	"placelists/pkg/rdg"
)

type placelistService struct {
	placelistRepository storage.PlacelistRepository
	placeRepository     storage.PlaceRepository
	userRepository      storage.UserRepository
}

func NewPlacelistService(
	placelistRepository storage.PlacelistRepository,
	placeRepository storage.PlaceRepository,
	userRepository storage.UserRepository,
) service.PlacelistService {
	return &placelistService{placelistRepository, placeRepository, userRepository}
}

func (s *placelistService) GetByPublicID(placelistID string, userID string) (models.Placelist, error) {
	placelist, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return models.Placelist{}, err
	}

	followed := false
	for _, follower := range placelist.FollowedUsers {
		if follower.PublicID == userID {
			followed = true
			break
		}
	}

	status := models.PlacelistNone
	if placelist.Author.PublicID == userID {
		status = models.PlacelistCreated
	} else if followed {
		status = models.PlacelistFollowed
	}

	foundPlacelist := models.Placelist{
		ID:             placelist.PublicID,
		Name:           placelist.Name,
		AuthorID:       placelist.Author.PublicID,
		AuthorUsername: placelist.Author.Username,
		Status:         status,
	}

	return foundPlacelist, nil
}

func (s *placelistService) GetByNameOrAuthor(query string, userID string) ([]models.Placelist, error) {
	placelists, err := s.placelistRepository.GetByNameOrAuthorFull(query)
	if err != nil {
		return []models.Placelist{}, err
	}

	foundPlacelists := []models.Placelist{}

	for _, placelist := range placelists {
		followed := false
		for _, follower := range placelist.FollowedUsers {
			if follower.PublicID == userID {
				followed = true
				break
			}
		}

		status := models.PlacelistNone
		if placelist.Author.PublicID == userID {
			status = models.PlacelistCreated
		} else if followed {
			status = models.PlacelistFollowed
		}

		newPlacelist := models.Placelist{
			ID:             placelist.PublicID,
			Name:           placelist.Name,
			AuthorID:       placelist.Author.PublicID,
			AuthorUsername: placelist.Author.Username,
			Status:         status,
		}
		foundPlacelists = append(foundPlacelists, newPlacelist)
	}

	return foundPlacelists, nil
}

func (s *placelistService) GetFollowedByUserID(userID string) ([]models.Placelist, error) {
	user, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []models.Placelist{}, err
	}

	foundPlacelists := []models.Placelist{}

	for _, placelist := range user.FollwedPlacelists {
		newPlacelist := models.Placelist{
			ID:             placelist.PublicID,
			Name:           placelist.Name,
			AuthorID:       placelist.Author.PublicID,
			AuthorUsername: placelist.Author.Username,
			Status:         models.PlacelistFollowed,
		}
		foundPlacelists = append(foundPlacelists, newPlacelist)
	}

	return foundPlacelists, nil
}

func (s *placelistService) GetCreatedByUserID(userID string) ([]models.Placelist, error) {
	user, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []models.Placelist{}, err
	}

	foundPlacelists := []models.Placelist{}

	for _, placelist := range user.FollwedPlacelists {
		newPlacelist := models.Placelist{
			ID:             placelist.PublicID,
			Name:           placelist.Name,
			AuthorID:       placelist.Author.PublicID,
			AuthorUsername: placelist.Author.Username,
			Status:         models.PlacelistCreated,
		}
		foundPlacelists = append(foundPlacelists, newPlacelist)
	}

	return foundPlacelists, nil
}

func (s *placelistService) GetPlacesByPublicID(placelistID string, userID string) ([]models.Place, error) {
	placelist, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return []models.Place{}, err
	}

	foundPlaces := []models.Place{}

	for _, place := range placelist.Places {
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

func (s *placelistService) Create(userID string, pc models.PlacelistCreate) error {
	user, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return err
	}

	placelist := entities.Placelist{
		ID:       rdg.GenerateID(),
		PublicID: rdg.GeneratePublicID(),
		Name:     pc.Name,
		AuthorID: user.ID,
	}

	err = s.placelistRepository.Create(placelist)

	return err
}

func (s *placelistService) UpdateByPublicID(placelistID string, userID string, pu models.PlacelistUpdate) error {
	if pu.Status == models.PlacelistCreated {
		return errors.New("impossible to create placelist with update function")
	}

	user, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return err
	}

	placelist, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return err
	}

	placelist.Name = pu.Name

	if pu.Status == models.PlacelistFollowed {
		placelist.FollowedUsers = append(placelist.FollowedUsers, user)
	} else if pu.Status == models.PlacelistNone {
		for i, follower := range placelist.FollowedUsers {
			if follower.PublicID == userID {
				placelist.FollowedUsers = append(placelist.FollowedUsers[:i], placelist.FollowedUsers[i+1:]...)
			}
		}
	}

	newPlaces := []entities.Place{}

	for _, placeID := range pu.PlacesIDs {
		place, err := s.placeRepository.GetByPublicID(placeID)
		if err == nil {
			newPlaces = append(newPlaces, place)
		}
	}

	placelist.Places = newPlaces

	err = s.placelistRepository.Update(placelist)

	return err
}
