package domain

import (
	"errors"
	"placelists/internal/service"
	"placelists/internal/service/models"
	"placelists/internal/storage"
	"placelists/internal/storage/entities"
	"placelists/pkg/rdg"

	"github.com/jinzhu/copier"
)

type placelistServiceImpl struct {
	placelistRepository storage.PlacelistRepository
	placeRepository     storage.PlaceRepository
	userRepository      storage.UserRepository
}

func NewPlacelistService(
	placelistRepository storage.PlacelistRepository,
	placeRepository storage.PlaceRepository,
	userRepository storage.UserRepository,
) service.PlacelistService {
	return &placelistServiceImpl{placelistRepository, placeRepository, userRepository}
}

func (s *placelistServiceImpl) GetByID(placelistID string, userID string) (models.Placelist, error) {
	placelistEntity, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return models.Placelist{}, err
	}

	followed := false
	for _, follower := range placelistEntity.FollowedUsers {
		if follower.PublicID == userID {
			followed = true
			break
		}
	}

	status := models.PlacelistNone
	if placelistEntity.Author.PublicID == userID {
		status = models.PlacelistCreated
	} else if followed {
		status = models.PlacelistFollowed
	}

	foundPlacelist := models.Placelist{
		ID:             placelistEntity.PublicID,
		Name:           placelistEntity.Name,
		AuthorID:       placelistEntity.Author.PublicID,
		AuthorUsername: placelistEntity.Author.Username,
		Status:         status,
	}

	return foundPlacelist, nil
}

func (s *placelistServiceImpl) GetByNameOrAuthor(query string, userID string) ([]models.Placelist, error) {
	placelistsEntities, err := s.placelistRepository.GetByNameOrAuthorFull(query)
	if err != nil {
		return []models.Placelist{}, err
	}

	foundPlacelists := []models.Placelist{}

	for _, placelistEntity := range placelistsEntities {
		followed := false
		for _, follower := range placelistEntity.FollowedUsers {
			if follower.PublicID == userID {
				followed = true
				break
			}
		}

		status := models.PlacelistNone
		if placelistEntity.Author.PublicID == userID {
			status = models.PlacelistCreated
		} else if followed {
			status = models.PlacelistFollowed
		}

		placelist := models.Placelist{
			ID:             placelistEntity.PublicID,
			Name:           placelistEntity.Name,
			AuthorID:       placelistEntity.Author.PublicID,
			AuthorUsername: placelistEntity.Author.Username,
			Status:         status,
		}

		foundPlacelists = append(foundPlacelists, placelist)
	}

	return foundPlacelists, nil
}

func (s *placelistServiceImpl) GetFollowedByUserID(userID string) ([]models.Placelist, error) {
	userEntity, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []models.Placelist{}, err
	}

	foundPlacelists := []models.Placelist{}

	for _, placelistEntity := range userEntity.FollwedPlacelists {
		placelist := models.Placelist{
			ID:             placelistEntity.PublicID,
			Name:           placelistEntity.Name,
			AuthorID:       placelistEntity.Author.PublicID,
			AuthorUsername: placelistEntity.Author.Username,
			Status:         models.PlacelistFollowed,
		}
		foundPlacelists = append(foundPlacelists, placelist)
	}

	return foundPlacelists, nil
}

func (s *placelistServiceImpl) GetCreatedByUserID(userID string) ([]models.Placelist, error) {
	userEntity, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []models.Placelist{}, err
	}

	foundPlacelists := []models.Placelist{}

	for _, placelistEntity := range userEntity.FollwedPlacelists {
		placelist := models.Placelist{
			ID:             placelistEntity.PublicID,
			Name:           placelistEntity.Name,
			AuthorID:       placelistEntity.Author.PublicID,
			AuthorUsername: placelistEntity.Author.Username,
			Status:         models.PlacelistCreated,
		}
		foundPlacelists = append(foundPlacelists, placelist)
	}

	return foundPlacelists, nil
}

func (s *placelistServiceImpl) GetPlacesByID(placelistID string, userID string) ([]models.Place, error) {
	placelistEntity, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return []models.Place{}, err
	}

	foundPlaces := []models.Place{}

	for _, placeEntity := range placelistEntity.Places {
		visited := false
		for _, visitor := range placeEntity.Visitors {
			if visitor.PublicID == userID {
				visited = true
				break
			}
		}

		place := models.Place{}
		copier.Copy(&placeEntity, &place)
		place.Visited = visited

		foundPlaces = append(foundPlaces, place)
	}

	return foundPlaces, nil
}

func (s *placelistServiceImpl) Create(userID string, pc models.PlacelistCreate) (models.Placelist, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return models.Placelist{}, err
	}

	placelistEntity := entities.Placelist{
		ID:       rdg.GenerateID(),
		PublicID: rdg.GeneratePublicID(),
		Name:     pc.Name,
		AuthorID: userEntity.ID,
	}

	err = s.placelistRepository.Create(placelistEntity)
	if err != nil {
		return models.Placelist{}, err
	}

	placelist := models.Placelist{}
	copier.Copy(&placelistEntity, &placelist)

	return placelist, err
}

func (s *placelistServiceImpl) UpdateByID(placelistID string, userID string, pu models.PlacelistUpdate) (models.Placelist, error) {
	if pu.Status == models.PlacelistCreated {
		return models.Placelist{}, errors.New("impossible to create placelist with update function")
	}

	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return models.Placelist{}, err
	}

	placelistEntity, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return models.Placelist{}, err
	}

	placelistEntity.Name = pu.Name

	if pu.Status == models.PlacelistFollowed {
		placelistEntity.FollowedUsers = append(placelistEntity.FollowedUsers, userEntity)
	} else if pu.Status == models.PlacelistNone {
		for i, follower := range placelistEntity.FollowedUsers {
			if follower.PublicID == userID {
				placelistEntity.FollowedUsers = append(placelistEntity.FollowedUsers[:i], placelistEntity.FollowedUsers[i+1:]...)
			}
		}
	}

	placesEntities := []entities.Place{}

	for _, placeID := range pu.PlacesIDs {
		placeEntity, err := s.placeRepository.GetByPublicID(placeID)
		if err == nil {
			placesEntities = append(placesEntities, placeEntity)
		}
	}

	placelistEntity.Places = placesEntities

	err = s.placelistRepository.Update(placelistEntity)
	if err != nil {
		return models.Placelist{}, err
	}

	placelist := models.Placelist{}
	copier.Copy(&placelistEntity, &placelist)

	return placelist, err
}
