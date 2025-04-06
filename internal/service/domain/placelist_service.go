package domain

import (
	"errors"

	"placelists-back/internal/service"
	"placelists-back/internal/service/enum"
	"placelists-back/internal/service/model"
	"placelists-back/internal/storage"
	"placelists-back/internal/storage/entity"
	"placelists-back/pkg/utils"

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

func (s *placelistServiceImpl) GetByID(placelistID string, userID string) (model.Placelist, error) {
	placelistEntity, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return model.Placelist{}, err
	}

	followed := false
	for _, follower := range placelistEntity.FollowedUsers {
		if follower.PublicID == userID {
			followed = true
			break
		}
	}

	status := enum.PlacelistNone
	if placelistEntity.Author.PublicID == userID {
		status = enum.PlacelistCreated
	} else if followed {
		status = enum.PlacelistFollowed
	}

	foundPlacelist := model.Placelist{
		ID:             placelistEntity.PublicID,
		Name:           placelistEntity.Name,
		AuthorID:       placelistEntity.Author.PublicID,
		AuthorUsername: placelistEntity.Author.Username,
		Status:         status,
	}

	return foundPlacelist, nil
}

func (s *placelistServiceImpl) GetByNameOrAuthor(query string, userID string) ([]model.Placelist, error) {
	placelistsEntities, err := s.placelistRepository.GetByNameOrAuthorFull(query)
	if err != nil {
		return []model.Placelist{}, err
	}

	var foundPlacelists []model.Placelist

	for _, placelistEntity := range placelistsEntities {
		followed := false
		for _, follower := range placelistEntity.FollowedUsers {
			if follower.PublicID == userID {
				followed = true
				break
			}
		}

		status := enum.PlacelistNone
		if placelistEntity.Author.PublicID == userID {
			status = enum.PlacelistCreated
		} else if followed {
			status = enum.PlacelistFollowed
		}

		placelist := model.Placelist{
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

func (s *placelistServiceImpl) GetFollowedByUserID(userID string) ([]model.Placelist, error) {
	userEntity, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []model.Placelist{}, err
	}

	var foundPlacelists []model.Placelist

	for _, placelistEntity := range userEntity.FollowedPlacelists {
		placelist := model.Placelist{
			ID:             placelistEntity.PublicID,
			Name:           placelistEntity.Name,
			AuthorID:       placelistEntity.Author.PublicID,
			AuthorUsername: placelistEntity.Author.Username,
			Status:         enum.PlacelistFollowed,
		}
		foundPlacelists = append(foundPlacelists, placelist)
	}

	return foundPlacelists, nil
}

func (s *placelistServiceImpl) GetCreatedByUserID(userID string) ([]model.Placelist, error) {
	userEntity, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []model.Placelist{}, err
	}

	var foundPlacelists []model.Placelist

	for _, placelistEntity := range userEntity.FollowedPlacelists {
		placelist := model.Placelist{
			ID:             placelistEntity.PublicID,
			Name:           placelistEntity.Name,
			AuthorID:       placelistEntity.Author.PublicID,
			AuthorUsername: placelistEntity.Author.Username,
			Status:         enum.PlacelistCreated,
		}
		foundPlacelists = append(foundPlacelists, placelist)
	}

	return foundPlacelists, nil
}

func (s *placelistServiceImpl) GetPlacesByID(placelistID string, userID string) ([]model.Place, error) {
	placelistEntity, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return []model.Place{}, err
	}

	var foundPlaces []model.Place

	for _, placeEntity := range placelistEntity.Places {
		visited := false
		for _, visitor := range placeEntity.Visitors {
			if visitor.PublicID == userID {
				visited = true
				break
			}
		}

		place := model.Place{}
		err = copier.Copy(&placeEntity, &place)
		if err != nil {
			return []model.Place{}, err
		}
		place.Visited = visited

		foundPlaces = append(foundPlaces, place)
	}

	return foundPlaces, nil
}

func (s *placelistServiceImpl) Create(userID string, pc model.PlacelistCreate) (model.Placelist, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return model.Placelist{}, err
	}

	placelistEntity := entity.Placelist{
		ID:       utils.GenerateID(),
		PublicID: utils.GeneratePublicID(),
		Name:     pc.Name,
		AuthorID: userEntity.ID,
	}

	err = s.placelistRepository.Create(placelistEntity)
	if err != nil {
		return model.Placelist{}, err
	}

	placelist := model.Placelist{}
	err = copier.Copy(&placelistEntity, &placelist)
	if err != nil {
		return model.Placelist{}, err
	}

	return placelist, err
}

func (s *placelistServiceImpl) UpdateByID(placelistID string, userID string, pu model.PlacelistUpdate) (model.Placelist, error) {
	if pu.Status == enum.PlacelistCreated {
		return model.Placelist{}, errors.New("impossible to create placelist with update function")
	}

	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return model.Placelist{}, err
	}

	placelistEntity, err := s.placelistRepository.GetByPublicIDFull(placelistID)
	if err != nil {
		return model.Placelist{}, err
	}

	placelistEntity.Name = pu.Name

	if pu.Status == enum.PlacelistFollowed {
		placelistEntity.FollowedUsers = append(placelistEntity.FollowedUsers, userEntity)
	} else if pu.Status == enum.PlacelistNone {
		for i, follower := range placelistEntity.FollowedUsers {
			if follower.PublicID == userID {
				placelistEntity.FollowedUsers = append(placelistEntity.FollowedUsers[:i], placelistEntity.FollowedUsers[i+1:]...)
			}
		}
	}

	var placesEntities []entity.Place

	for _, placeID := range pu.PlacesIDs {
		placeEntity, err := s.placeRepository.GetByPublicID(placeID)
		if err == nil {
			placesEntities = append(placesEntities, placeEntity)
		}
	}

	placelistEntity.Places = placesEntities

	err = s.placelistRepository.Update(placelistEntity)
	if err != nil {
		return model.Placelist{}, err
	}

	placelist := model.Placelist{}
	err = copier.Copy(&placelistEntity, &placelist)
	if err != nil {
		return model.Placelist{}, err
	}

	return placelist, err
}
