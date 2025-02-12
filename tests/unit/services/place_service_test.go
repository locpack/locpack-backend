package services_test

import (
	"errors"
	"placelists/internal/service/domain"
	"placelists/internal/service/models"
	"placelists/internal/storage/entities"
	"placelists/tests/unit/fakes"
	"testing"

	"github.com/go-playground/assert"
)

func TestPlaceServiceGetByPublicID(t *testing.T) {
	placeRepo := &fakes.PlaceRepositoryFakeImpl{
		Places: []entities.Place{
			{
				PublicID: "place-1",
				Name:     "Test Place",
				Address:  "123 Test St",
				Visitors: []entities.User{{PublicID: "user-1"}},
			},
		},
	}

	userRepo := &fakes.UserRepositoryFakeImpl{
		Users: []entities.User{{PublicID: "user-1"}},
	}

	service := domain.NewPlaceService(placeRepo, userRepo)

	tests := []struct {
		name        string
		placeID     string
		userID      string
		expectedErr error
		expected    *models.Place
	}{
		{
			name:        "Place found and user visited",
			placeID:     "place-1",
			userID:      "user-1",
			expectedErr: nil,
			expected: &models.Place{
				ID:      "place-1",
				Name:    "Test Place",
				Address: "123 Test St",
				Visited: true,
			},
		},
		{
			name:        "Place not found",
			placeID:     "place-2",
			userID:      "user-1",
			expectedErr: errors.New("place not found"),
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			place, err := service.GetByPublicID(tt.placeID, tt.userID)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expected, place)
		})
	}
}

func TestPlaceServiceGetByNameOrAddress(t *testing.T) {
	placeRepo := &fakes.PlaceRepositoryFakeImpl{
		Places: []entities.Place{
			{
				PublicID: "place-1",
				Name:     "Test Place",
				Address:  "123 Test St",
				Visitors: []entities.User{{PublicID: "user-1"}},
			},
		},
	}

	userRepo := &fakes.UserRepositoryFakeImpl{
		Users: []entities.User{{PublicID: "user-1"}},
	}

	service := domain.NewPlaceService(placeRepo, userRepo)

	tests := []struct {
		name        string
		query       string
		userID      string
		expectedErr error
		expected    *[]models.Place
	}{
		{
			name:        "Place found and user visited",
			query:       "Test",
			userID:      "user-1",
			expectedErr: nil,
			expected: &[]models.Place{
				{
					ID:      "place-1",
					Name:    "Test Place",
					Address: "123 Test St",
					Visited: true,
				},
			},
		},
		{
			name:        "No places found",
			query:       "Unknown",
			userID:      "user-1",
			expectedErr: errors.New("no places found"),
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			places, err := service.GetByNameOrAddress(tt.query, tt.userID)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expected, places)
		})
	}
}

func TestPlaceServiceCreate(t *testing.T) {
	placeRepo := &fakes.PlaceRepositoryFakeImpl{}

	userRepo := &fakes.UserRepositoryFakeImpl{
		Users: []entities.User{{PublicID: "user-1"}},
	}

	service := domain.NewPlaceService(placeRepo, userRepo)

	tests := []struct {
		name        string
		userID      string
		pc          *models.PlaceCreate
		expectedErr error
	}{
		{
			name:   "Create place successfully",
			userID: "user-1",
			pc: &models.PlaceCreate{
				Name:    "New Place",
				Address: "456 New St",
				Visited: true,
			},
			expectedErr: nil,
		},
		{
			name:   "User not found",
			userID: "user-2",
			pc: &models.PlaceCreate{
				Name:    "New Place",
				Address: "456 New St",
				Visited: true,
			},
			expectedErr: errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.Create(tt.userID, tt.pc)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestPlaceServiceUpdateByPublicID(t *testing.T) {
	placeRepo := &fakes.PlaceRepositoryFakeImpl{
		Places: []entities.Place{
			{
				PublicID: "place-1",
				Name:     "Old Place",
				Address:  "123 Old St",
				Visitors: []entities.User{{PublicID: "user-1"}},
			},
		},
	}

	userRepo := &fakes.UserRepositoryFakeImpl{
		Users: []entities.User{{PublicID: "user-1"}},
	}

	service := domain.NewPlaceService(placeRepo, userRepo)

	tests := []struct {
		name        string
		placeID     string
		userID      string
		pu          *models.PlaceUpdate
		expectedErr error
	}{
		{
			name:    "Update place successfully",
			placeID: "place-1",
			userID:  "user-1",
			pu: &models.PlaceUpdate{
				Name:    "Updated Place",
				Address: "456 Updated St",
				Visited: true,
			},
			expectedErr: nil,
		},
		{
			name:    "Place not found",
			placeID: "place-2",
			userID:  "user-1",
			pu: &models.PlaceUpdate{
				Name:    "Updated Place",
				Address: "456 Updated St",
				Visited: true,
			},
			expectedErr: errors.New("place not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.UpdateByPublicID(tt.placeID, tt.userID, tt.pu)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
