package domain

import (
	"placelists/internal/service"
	"placelists/internal/service/models"
	"placelists/internal/storage"
	"placelists/internal/storage/entities"
	"placelists/pkg/rdg"

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

func (s *placeServiceImpl) GetByID(placeID string, userID string) (models.Place, error) {
	placeEntity, err := s.placeRepository.GetByPublicIDFull(placeID)
	if err != nil {
		return models.Place{}, err
	}

	visited := false
	for _, visitor := range placeEntity.Visitors {
		if visitor.PublicID == userID {
			visited = true
			break
		}
	}

	place := models.Place{}
	err = copier.Copy(&placeEntity, &place)
	if err != nil {
		return models.Place{}, err
	}
	place.Visited = visited

	return place, nil
}

func (s *placeServiceImpl) GetByNameOrAddress(query string, userID string) ([]models.Place, error) {
	placesEntities, err := s.placeRepository.GetByNameOrAddressFull(query)
	if err != nil {
		return []models.Place{}, err
	}

	foundPlaces := []models.Place{}

	for _, placeEntity := range placesEntities {
		visited := false
		for _, visitor := range placeEntity.Visitors {
			if visitor.PublicID == userID {
				visited = true
				break
			}
		}

		place := models.Place{}
		err = copier.Copy(&placeEntity, &place)
		if err != nil {
			return []models.Place{}, err
		}
		place.Visited = visited

		foundPlaces = append(foundPlaces, place)
	}

	return foundPlaces, nil
}

func (s *placeServiceImpl) Create(userID string, pc models.PlaceCreate) (models.Place, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return models.Place{}, err
	}

	visitors := []entities.User{}
	if pc.Visited {
		visitors = append(visitors, userEntity)
	}

	placeEntity := entities.Place{
		ID:       rdg.GenerateID(),
		PublicID: rdg.GeneratePublicID(),
		Name:     pc.Name,
		Address:  pc.Address,
		AuthorID: userEntity.ID,
		Visitors: visitors,
	}

	err = s.placeRepository.Create(placeEntity)
	if err != nil {
		return models.Place{}, err
	}

	place := models.Place{}
	err = copier.Copy(&placeEntity, &place)
	if err != nil {
		return models.Place{}, err
	}

	return place, err
}

func (s *placeServiceImpl) UpdateByID(placeID string, userID string, pu models.PlaceUpdate) (models.Place, error) {
	userEntity, err := s.userRepository.GetByPublicID(userID)
	if err != nil {
		return models.Place{}, err
	}

	placeEntity, err := s.placeRepository.GetByPublicIDFull(placeID)
	if err != nil {
		return models.Place{}, err
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
		return models.Place{}, err
	}

	place := models.Place{}
	err = copier.Copy(&placeEntity, &place)
	if err != nil {
		return models.Place{}, err
	}

	return place, err
}
