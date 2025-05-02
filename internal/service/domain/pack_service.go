package domain

import (
	"errors"

	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/enum/pack_status"
	"locpack-backend/pkg/utils/random"

	"github.com/jinzhu/copier"
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

	followed := false
	for _, follower := range packEntity.FollowedUsers {
		if follower.PublicID == userID {
			followed = true
			break
		}
	}

	status := pack_status.None
	if packEntity.Author.PublicID == userID {
		status = pack_status.Created
	} else if followed {
		status = pack_status.Followed
	}

	foundPack := model.Pack{
		ID:             packEntity.PublicID,
		Name:           packEntity.Name,
		AuthorID:       packEntity.Author.PublicID,
		AuthorUsername: packEntity.Author.Username,
		Status:         status,
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
		followed := false
		for _, follower := range packEntity.FollowedUsers {
			if follower.PublicID == userID {
				followed = true
				break
			}
		}

		status := pack_status.None
		if packEntity.Author.PublicID == userID {
			status = pack_status.Created
		} else if followed {
			status = pack_status.Followed
		}

		pack := model.Pack{
			ID:             packEntity.PublicID,
			Name:           packEntity.Name,
			AuthorID:       packEntity.Author.PublicID,
			AuthorUsername: packEntity.Author.Username,
			Status:         status,
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
			ID:             packEntity.PublicID,
			Name:           packEntity.Name,
			AuthorID:       packEntity.Author.PublicID,
			AuthorUsername: packEntity.Author.Username,
			Status:         pack_status.Followed,
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

	for _, packEntity := range userEntity.FollowedPacks {
		pack := model.Pack{
			ID:             packEntity.PublicID,
			Name:           packEntity.Name,
			AuthorID:       packEntity.Author.PublicID,
			AuthorUsername: packEntity.Author.Username,
			Status:         pack_status.Created,
		}
		foundPacks = append(foundPacks, pack)
	}

	return foundPacks, nil
}

func (s *packServiceImpl) GetPlacesByID(packID string, userID string) ([]model.Place, error) {
	packEntity, err := s.packRepository.GetByPublicIDFull(packID)
	if err != nil {
		return []model.Place{}, err
	}

	var foundPlaces []model.Place

	for _, placeEntity := range packEntity.Places {
		visited := false
		for _, visitor := range placeEntity.Visitors {
			if visitor.PublicID == userID {
				visited = true
				break
			}
		}

		place := model.Place{}
		err = copier.Copy(&place, &placeEntity)
		if err != nil {
			return []model.Place{}, err
		}
		place.Visited = visited

		foundPlaces = append(foundPlaces, place)
	}

	return foundPlaces, nil
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

	pack := model.Pack{}
	err = copier.Copy(&pack, &packEntity)
	if err != nil {
		return model.Pack{}, err
	}

	return pack, err
}

func (s *packServiceImpl) UpdateByID(packID string, userID string, pu model.PackUpdate) (model.Pack, error) {
	if pu.Status == pack_status.Created {
		return model.Pack{}, errors.New("impossible to create pack with update function")
	}

	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return model.Pack{}, err
	}

	packEntity, err := s.packRepository.GetByPublicIDFull(packID)
	if err != nil {
		return model.Pack{}, err
	}

	if packEntity.Author.ID != userEntity.ID {
		return model.Pack{}, errors.New("user is not author")
	}

	packEntity.Name = pu.Name

	if pu.Status == pack_status.Followed {
		packEntity.FollowedUsers = append(packEntity.FollowedUsers, userEntity)
	} else if pu.Status == pack_status.None {
		for i, follower := range packEntity.FollowedUsers {
			if follower.PublicID == userID {
				packEntity.FollowedUsers = append(packEntity.FollowedUsers[:i], packEntity.FollowedUsers[i+1:]...)
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

	packEntity.Places = placesEntities

	err = s.packRepository.Update(packEntity)
	if err != nil {
		return model.Pack{}, err
	}

	pack := model.Pack{}
	err = copier.Copy(&pack, &packEntity)
	if err != nil {
		return model.Pack{}, err
	}

	return pack, err
}
