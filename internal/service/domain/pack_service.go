package domain

import (
	"errors"

	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/enum/pack_status"
	"locpack-backend/pkg/types"
	"locpack-backend/pkg/utils/random"
)

type packServiceImpl struct {
	packRepository  storage.PackRepository
	placeRepository storage.PlaceRepository
	userRepository  storage.UserRepository
}

func NewPackService(
	packRepository storage.PackRepository,
	placeRepository storage.PlaceRepository,
	userRepository storage.UserRepository,
) service.PackService {
	return &packServiceImpl{packRepository, placeRepository, userRepository}
}

func (s *packServiceImpl) GetByID(packID string, userID string) (model.Pack, error) {
	packEntity, err := s.packRepository.GetByPublicIDFull(packID)
	if err != nil {
		return model.Pack{}, err
	}

	foundPack := model.Pack{
		ID:     packEntity.PublicID,
		Name:   packEntity.Name,
		Status: s.getPackStatus(packEntity, userID),
		Places: s.mapPlaceEntitiesToModels(packEntity.Places, userID),
		Author: model.User{
			ID:       packEntity.Author.PublicID,
			Username: packEntity.Author.Username,
		},
	}

	return foundPack, nil
}

func (s *packServiceImpl) GetByNameOrAuthor(query string, userID string) ([]model.Pack, error) {
	packsEntities, err := s.packRepository.GetByNameOrAuthorFull(query)
	if err != nil {
		return []model.Pack{}, err
	}

	var foundPacks []model.Pack
	for _, packEntity := range packsEntities {
		pack := model.Pack{
			ID:     packEntity.PublicID,
			Name:   packEntity.Name,
			Status: s.getPackStatus(packEntity, userID),
			Places: s.mapPlaceEntitiesToModels(packEntity.Places, userID),
			Author: model.User{
				ID:       packEntity.Author.PublicID,
				Username: packEntity.Author.Username,
			},
		}
		foundPacks = append(foundPacks, pack)
	}

	return foundPacks, nil
}

func (s *packServiceImpl) GetFollowedByUserID(userID string) ([]model.Pack, error) {
	userEntity, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []model.Pack{}, err
	}

	var foundPacks []model.Pack
	for _, packEntity := range userEntity.FollowedPacks {
		pack := model.Pack{
			ID:     packEntity.PublicID,
			Name:   packEntity.Name,
			Status: pack_status.Followed,
			Places: s.mapPlaceEntitiesToModels(packEntity.Places, userID),
			Author: model.User{
				ID:       packEntity.Author.PublicID,
				Username: packEntity.Author.Username,
			},
		}
		foundPacks = append(foundPacks, pack)
	}

	return foundPacks, nil
}

func (s *packServiceImpl) GetCreatedByUserID(userID string) ([]model.Pack, error) {
	userEntity, err := s.userRepository.GetByPublicIDFull(userID)
	if err != nil {
		return []model.Pack{}, err
	}

	var foundPacks []model.Pack
	for _, packEntity := range userEntity.CreatedPacks {
		pack := model.Pack{
			ID:     packEntity.PublicID,
			Name:   packEntity.Name,
			Status: pack_status.Created,
			Places: s.mapPlaceEntitiesToModels(packEntity.Places, userID),
			Author: model.User{
				ID:       packEntity.Author.PublicID,
				Username: packEntity.Author.Username,
			},
		}
		foundPacks = append(foundPacks, pack)
	}

	return foundPacks, nil
}

func (s *packServiceImpl) Create(userID string, pc model.PackCreate) (model.Pack, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return model.Pack{}, err
	}

	packEntity := entity.Pack{
		ID:       random.GenerateID(),
		PublicID: random.GeneratePublicID(),
		Name:     pc.Name,
		AuthorID: userEntity.ID,
	}

	err = s.packRepository.Create(packEntity)
	if err != nil {
		return model.Pack{}, err
	}

	pack := model.Pack{
		ID:     packEntity.PublicID,
		Name:   packEntity.Name,
		Status: pack_status.Created,
		Places: []model.Place{},
		Author: model.User{
			ID:       userEntity.PublicID,
			Username: userEntity.Username,
		},
	}

	return pack, err
}

func (s *packServiceImpl) UpdateByID(packID string, userID string, pu model.PackUpdate) (model.Pack, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return model.Pack{}, err
	}

	packEntity, err := s.packRepository.GetByPublicIDFull(packID)
	if err != nil {
		return model.Pack{}, err
	}

	status := s.getPackStatus(packEntity, userID)

	if status == pack_status.None {
		if pu.Status != pack_status.Followed {
			return model.Pack{}, errors.New("it is possible only to follow this pack")
		}
		packEntity.FollowedUsers = append(packEntity.FollowedUsers, userEntity)
		status = pack_status.Followed
	} else if status == pack_status.Created {
		packEntity.Name = pu.Name
		var placeEntities []entity.Place
		for _, placeID := range pu.PlacesIDs {
			placeEntity, err := s.placeRepository.GetByPublicID(placeID)
			if err == nil {
				placeEntities = append(placeEntities, placeEntity)
			}
		}
		packEntity.Places = placeEntities
	} else if status == pack_status.Followed {
		if pu.Status != pack_status.None {
			return model.Pack{}, errors.New("it is possible only to unfollow this pack")
		}
		for i, follower := range packEntity.FollowedUsers {
			if follower.PublicID == userID {
				packEntity.FollowedUsers = append(packEntity.FollowedUsers[:i], packEntity.FollowedUsers[i+1:]...)
			}
		}
		status = pack_status.None
	}

	err = s.packRepository.Update(packEntity)
	if err != nil {
		return model.Pack{}, err
	}

	pack := model.Pack{
		ID:     packEntity.PublicID,
		Name:   packEntity.Name,
		Status: status,
		Places: s.mapPlaceEntitiesToModels(packEntity.Places, userID),
		Author: model.User{
			ID:       packEntity.Author.PublicID,
			Username: packEntity.Author.Username,
		},
	}

	return pack, nil
}

func (s *packServiceImpl) mapPlaceEntitiesToModels(placeEntities []entity.Place, userID string) []model.Place {
	places := []model.Place{}
	for _, placeEntity := range placeEntities {
		visited := false
		for _, visitor := range placeEntity.Visitors {
			if visitor.PublicID == userID {
				visited = true
				break
			}
		}

		place := model.Place{
			ID:      placeEntity.PublicID,
			Name:    placeEntity.Name,
			Address: placeEntity.Address,
			Visited: visited,
		}

		places = append(places, place)
	}

	return places
}

func (s *packServiceImpl) getPackStatus(packEntity entity.Pack, userID string) types.PackStatus {
	if packEntity.Author.PublicID == userID {
		return pack_status.Created
	}

	followed := false
	for _, follower := range packEntity.FollowedUsers {
		if follower.PublicID == userID {
			followed = true
			break
		}
	}
	if followed {
		return pack_status.Followed
	}

	return pack_status.None
}
