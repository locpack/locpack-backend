package domain

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
)

func TestUserService_GetByID(t *testing.T) {
	t.Parallel()

	userID := "user-123"
	userUUID := uuid.New()

	tests := []struct {
		name        string
		userID      string
		setupMocks  func(*storage.MockUserRepository)
		expected    model.User
		expectError bool
	}{
		{
			name:   "success",
			userID: userID,
			setupMocks: func(userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)
			},
			expected: model.User{
				ID:       userID,
				Username: "testuser",
			},
			expectError: false,
		},
		{
			name:   "user not found",
			userID: "not-exist",
			setupMocks: func(userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID("not-exist").Return(entity.User{}, errors.New("user not found"))
			},
			expected:    model.User{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := storage.NewMockUserRepository(t)
			tt.setupMocks(userRepo)

			svc := NewUserService(userRepo)

			result, err := svc.GetByID(tt.userID)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.ID, result.ID)
				assert.Equal(t, tt.expected.Username, result.Username)
			}
		})
	}
}

func TestUserService_UpdateByID(t *testing.T) {
	t.Parallel()

	userID := "user-123"
	userUUID := uuid.New()

	tests := []struct {
		name        string
		userID      string
		input       model.UserUpdate
		setupMocks  func(*storage.MockUserRepository)
		expected    model.User
		expectError bool
	}{
		{
			name:   "success",
			userID: userID,
			input: model.UserUpdate{
				Username: "updateduser",
			},
			setupMocks: func(userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "olduser",
				}, nil)

				userRepo.EXPECT().Update(mock.MatchedBy(func(u entity.User) bool {
					return u.Username == "updateduser" && u.PublicID == "updateduser"
				})).Return(nil)
			},
			expected: model.User{
				ID:       "updateduser",
				Username: "updateduser",
			},
			expectError: false,
		},
		{
			name:   "user not found",
			userID: "not-exist",
			input: model.UserUpdate{
				Username: "irrelevant",
			},
			setupMocks: func(userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID("not-exist").Return(entity.User{}, errors.New("not found"))
			},
			expected:    model.User{},
			expectError: true,
		},
		{
			name:   "update fails",
			userID: userID,
			input: model.UserUpdate{
				Username: "baduser",
			},
			setupMocks: func(userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "olduser",
				}, nil)

				userRepo.EXPECT().Update(mock.Anything).Return(errors.New("update failed"))
			},
			expected:    model.User{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := storage.NewMockUserRepository(t)
			tt.setupMocks(userRepo)

			svc := NewUserService(userRepo)

			result, err := svc.UpdateByID(tt.userID, tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.ID, result.ID)
				assert.Equal(t, tt.expected.Username, result.Username)
			}
		})
	}
}
