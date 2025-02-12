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
				Username: "testuser",
			},
		},
	}

	service := domain.NewUserService(userRepo)

	tests := []struct {
		name        string
		publicID    string
		expected    *models.User
		expectedErr error
	}{
		{
			name:     "User found",
			publicID: "user-1",
			expected: &models.User{
				ID:       "user-1",
				Username: "testuser",
			},
			expectedErr: nil,
		},
		{
			name:        "User not found",
			publicID:    "user-2",
			expected:    nil,
			expectedErr: errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.GetByPublicID(tt.publicID)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expected, user)
		})
	}
}

func TestUserServiceUpdateByPublicID(t *testing.T) {
	userRepo := &fakes.UserRepositoryFakeImpl{
		Users: []entities.User{
			{
				PublicID: "user-1",
				Username: "olduser",
			},
		},
	}

	service := domain.NewUserService(userRepo)

	tests := []struct {
		name        string
		publicID    string
		uu          *models.UserUpdate
		expectedErr error
	}{
		{
			name:     "Update user successfully",
			publicID: "user-1",
			uu: &models.UserUpdate{
				Username: "newuser",
			},
			expectedErr: nil,
		},
		{
			name:     "User not found",
			publicID: "user-2",
			uu: &models.UserUpdate{
				Username: "newuser",
			},
			expectedErr: errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.UpdateByPublicID(tt.publicID, tt.uu)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}