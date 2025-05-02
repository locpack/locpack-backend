package domain

import (
	"errors"

	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/utils/random"

	"github.com/jinzhu/copier"
)

type placeServiceImpl struct {
	placeRepository storage.PlaceRepository
	userRepository  storage.UserRepository
}

func NewPlaceService(
	placeRepository storage.PlaceRepository,
	userRepository storage.UserRepository,
) service.PlaceService {
	return &placeServiceImpl{placeRepository, userRepository}
}

func (s *placeServiceImpl) GetByID(placeID string, userID string) (model.Place, error) {
	placeEntity, err := s.placeRepository.GetByPublicIDFull(placeID)
	if err != nil {
		return model.Place{}, err
	}

	visited := false
	for _, visitor := range placeEntity.Visitors {
		if visitor.PublicID == userID {
			visited = true
			break
		}
	}

	var place model.Place

	err = copier.Copy(&place, &placeEntity)
	if err != nil {
		return model.Place{}, err
	}
	place.Visited = visited

	return place, nil
}

func (s *placeServiceImpl) GetByNameOrAddress(query string, userID string) ([]model.Place, error) {
	placesEntities, err := s.placeRepository.GetByNameOrAddressFull(query)
	if err != nil {
		return []model.Place{}, err
	}

	var foundPlaces []model.Place

	for _, placeEntity := range placesEntities {
		visited := false
		for _, visitor := range placeEntity.Visitors {
			if visitor.PublicID == userID {
				visited = true
				break
			}
		}

		var place model.Place
		err = copier.Copy(&place, &placeEntity)
		if err != nil {
			return []model.Place{}, err
		}
		place.Visited = visited

		foundPlaces = append(foundPlaces, place)
	}

	return foundPlaces, nil
}

func (s *placeServiceImpl) Create(userID string, pc model.PlaceCreate) (model.Place, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return model.Place{}, err
	}

	var visitors []entity.User
	if pc.Visited {
		visitors = append(visitors, userEntity)
	}

	placeEntity := entity.Place{
		ID:       random.GenerateID(),
		PublicID: random.GeneratePublicID(),
		Name:     pc.Name,
		Address:  pc.Address,
		AuthorID: userEntity.ID,
		Visitors: visitors,
	}

	err = s.placeRepository.Create(placeEntity)
	if err != nil {
		return model.Place{}, err
	}

	place := model.Place{}
	err = copier.Copy(&place, &placeEntity)
	if err != nil {
		return model.Place{}, err
	}

	place.Visited = pc.Visited

	return place, err
}

func (s *placeServiceImpl) UpdateByID(placeID string, userID string, pu model.PlaceUpdate) (model.Place, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return model.Place{}, err
	}

	placeEntity, err := s.placeRepository.GetByPublicIDFull(placeID)
	if err != nil {
		return model.Place{}, err
	}

	if placeEntity.Author.ID != userEntity.ID {
		return model.Place{}, errors.New("user is not author")
	}

	placeEntity.Name = pu.Name
	placeEntity.Address = pu.Address

	if pu.Visited {
		placeEntity.Visitors = append(placeEntity.Visitors, userEntity)
	} else {
		for i, visitor := range placeEntity.Visitors {
			if visitor.PublicID == userID {
				placeEntity.Visitors = append(placeEntity.Visitors[:i], placeEntity.Visitors[i+1:]...)
			}
		}
	}

	err = s.placeRepository.Update(placeEntity)
	if err != nil {
		return model.Place{}, err
	}

	place := model.Place{}
	err = copier.Copy(&place, &placeEntity)
	if err != nil {
		return model.Place{}, err
	}

	return place, err
}
