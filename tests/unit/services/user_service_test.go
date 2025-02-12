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

func TestUserServiceGetByPublicID(t *testing.T) {
	userRepo := &fakes.UserRepositoryFakeImpl{
		Users: []entities.User{
			{
				PublicID: "user-1",
				Username: "user-1",
			},
			{
				PublicID: "user-2",
				Username: "user-2",
			},
		},
	}

	service := domain.NewUserService(userRepo)

	tests := []struct {
		name              string
		publicID          string
		expectedResult    *models.User
		expectedErr       error
		expectedCondition []entities.User
	}{
		{
			name:     "User found",
			publicID: "user-1",
			expectedResult: &models.User{
				ID:       "user-1",
				Username: "user-1",
			},
			expectedErr:       nil,
			expectedCondition: userRepo.Users,
		},
		{
			name:              "User not found",
			publicID:          "user-3",
			expectedResult:    nil,
			expectedErr:       errors.New("user not found"),
			expectedCondition: userRepo.Users,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.GetByPublicID(tt.publicID)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResult, user)
			assert.Equal(t, tt.expectedCondition, userRepo.Users)
		})
	}
}

func TestUserServiceUpdateByPublicID(t *testing.T) {
	userRepo := &fakes.UserRepositoryFakeImpl{
		Users: []entities.User{
			{
				PublicID: "user-1",
				Username: "user-1",
			},
			{
				PublicID: "user-2",
				Username: "user-2",
			},
		},
	}

	service := domain.NewUserService(userRepo)

	tests := []struct {
		name              string
		publicID          string
		uu                *models.UserUpdate
		expectedErr       error
		expectedCondition []entities.User
	}{
		{
			name:     "Update user successfully",
			publicID: "user-1",
			uu: &models.UserUpdate{
				Username: "user-3",
			},
			expectedErr: nil,
			expectedCondition: []entities.User{
				{
					PublicID: "user-3",
					Username: "user-3",
				},
				{
					PublicID: "user-2",
					Username: "user-2",
				},
			},
		},
		{
			name:     "User not found",
			publicID: "user-0",
			uu: &models.UserUpdate{
				Username: "newuser",
			},
			expectedErr:       errors.New("user not found"),
			expectedCondition: userRepo.Users,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.UpdateByPublicID(tt.publicID, tt.uu)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedCondition, userRepo.Users)
		})
	}
}
