package domain

import (
	"errors"
	"testing"

	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPlaceService_GetByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		placeID       string
		userID        string
		setupMocks    func(*storage.MockPlaceRepository)
		expectedPlace model.Place
		expectError   bool
	}{
		{
			name:    "success - user is a visitor",
			placeID: "place123",
			userID:  "user123",
			setupMocks: func(repo *storage.MockPlaceRepository) {
				repo.EXPECT().GetByPublicIDFull("place123").Return(entity.Place{
					PublicID: "place123",
					Name:     "Test Place",
					Address:  "Test Address",
					Visitors: []entity.User{
						{PublicID: "user123"},
					},
				}, nil)
			},
			expectedPlace: model.Place{
				ID:      "place123",
				Name:    "Test Place",
				Address: "Test Address",
				Visited: true,
			},
			expectError: false,
		},
		{
			name:    "success - user is not a visitor",
			placeID: "place123",
			userID:  "user456",
			setupMocks: func(repo *storage.MockPlaceRepository) {
				repo.EXPECT().GetByPublicIDFull("place123").Return(entity.Place{
					PublicID: "place123",
					Name:     "Test Place",
					Address:  "Test Address",
					Visitors: []entity.User{
						{PublicID: "other-user"},
					},
				}, nil)
			},
			expectedPlace: model.Place{
				ID:      "place123",
				Name:    "Test Place",
				Address: "Test Address",
				Visited: false,
			},
			expectError: false,
		},
		{
			name:    "repository error",
			placeID: "place123",
			userID:  "user123",
			setupMocks: func(repo *storage.MockPlaceRepository) {
				repo.EXPECT().GetByPublicIDFull("place123").Return(entity.Place{}, errors.New("database error"))
			},
			expectedPlace: model.Place{},
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := storage.NewMockPlaceRepository(t)
			tt.setupMocks(mockRepo)

			service := &placeServiceImpl{
				placeRepository: mockRepo,
			}

			result, err := service.GetByID(tt.placeID, tt.userID)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedPlace, result)
		})
	}
}

func TestPlaceService_GetByNameOrAddress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		query       string
		userID      string
		setupMocks  func(*storage.MockPlaceRepository)
		expected    []model.Place
		expectError bool
	}{
		{
			name:   "success - visited match",
			query:  "park",
			userID: "user123",
			setupMocks: func(repo *storage.MockPlaceRepository) {
				repo.EXPECT().GetByNameOrAddressFull("park").Return([]entity.Place{
					{
						PublicID: "place123",
						Name:     "Central Park",
						Address:  "Park Avenue",
						Visitors: []entity.User{
							{PublicID: "user123"},
						},
					},
				}, nil)
			},
			expected: []model.Place{
				{
					ID:      "place123",
					Name:    "Central Park",
					Address: "Park Avenue",
					Visited: true,
				},
			},
			expectError: false,
		},
		{
			name:   "success - multiple places",
			query:  "cafe",
			userID: "user123",
			setupMocks: func(repo *storage.MockPlaceRepository) {
				repo.EXPECT().GetByNameOrAddressFull("cafe").Return([]entity.Place{
					{
						PublicID: "place123",
						Name:     "Cafe One",
						Address:  "First Street",
						Visitors: []entity.User{
							{PublicID: "user123"},
						},
					},
					{
						PublicID: "place456",
						Name:     "Cafe Two",
						Address:  "Second Street",
						Visitors: []entity.User{
							{PublicID: "other-user"},
						},
					},
				}, nil)
			},
			expected: []model.Place{
				{
					ID:      "place123",
					Name:    "Cafe One",
					Address: "First Street",
					Visited: true,
				},
				{
					ID:      "place456",
					Name:    "Cafe Two",
					Address: "Second Street",
					Visited: false,
				},
			},
			expectError: false,
		},
		{
			name:   "empty result",
			query:  "nonexistent",
			userID: "user123",
			setupMocks: func(repo *storage.MockPlaceRepository) {
				repo.EXPECT().GetByNameOrAddressFull("nonexistent").Return([]entity.Place{}, nil)
			},
			expected:    []model.Place(nil),
			expectError: false,
		},
		{
			name:   "repository error",
			query:  "error",
			userID: "user123",
			setupMocks: func(repo *storage.MockPlaceRepository) {
				repo.EXPECT().GetByNameOrAddressFull("error").Return(nil, errors.New("database error"))
			},
			expected:    []model.Place{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := storage.NewMockPlaceRepository(t)
			tt.setupMocks(mockRepo)

			service := &placeServiceImpl{
				placeRepository: mockRepo,
			}

			result, err := service.GetByNameOrAddress(tt.query, tt.userID)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPlaceService_Create(t *testing.T) {
	t.Parallel()

	userID := "user123"
	userUUID := uuid.New()

	tests := []struct {
		name        string
		userID      string
		input       model.PlaceCreate
		setupMocks  func(*storage.MockPlaceRepository, *storage.MockUserRepository)
		expected    model.Place
		expectError bool
	}{
		{
			name:   "success - visited true",
			userID: userID,
			input: model.PlaceCreate{
				Name:    "New Place",
				Address: "New Address",
				Visited: true,
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)

				placeRepo.EXPECT().Create(mock.AnythingOfType("entity.Place")).Run(func(p entity.Place) {
					assert.Equal(t, "New Place", p.Name)
					assert.Equal(t, "New Address", p.Address)
					assert.Equal(t, userUUID, p.AuthorID)
					assert.Len(t, p.Visitors, 1)
					assert.Equal(t, userID, p.Visitors[0].PublicID)
				}).Return(nil)
			},
			expected: model.Place{
				Name:    "New Place",
				Address: "New Address",
				Visited: true,
			},
			expectError: false,
		},
		{
			name:   "success - visited false",
			userID: userID,
			input: model.PlaceCreate{
				Name:    "Another Place",
				Address: "Another Address",
				Visited: false,
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)

				placeRepo.EXPECT().Create(mock.AnythingOfType("entity.Place")).Run(func(p entity.Place) {
					assert.Equal(t, "Another Place", p.Name)
					assert.Equal(t, "Another Address", p.Address)
					assert.Equal(t, userUUID, p.AuthorID)
					assert.Len(t, p.Visitors, 0)
				}).Return(nil)
			},
			expected: model.Place{
				Name:    "Another Place",
				Address: "Another Address",
				Visited: false,
			},
			expectError: false,
		},
		{
			name:   "user not found",
			userID: "nonexistent",
			input: model.PlaceCreate{
				Name:    "New Place",
				Address: "New Address",
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID("nonexistent").Return(entity.User{}, errors.New("user not found"))
			},
			expected:    model.Place{},
			expectError: true,
		},
		{
			name:   "repository error",
			userID: userID,
			input: model.PlaceCreate{
				Name:    "Error Place",
				Address: "Error Address",
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)

				placeRepo.EXPECT().Create(mock.AnythingOfType("entity.Place")).Return(errors.New("database error"))
			},
			expected:    model.Place{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeRepo := storage.NewMockPlaceRepository(t)
			userRepo := storage.NewMockUserRepository(t)
			tt.setupMocks(placeRepo, userRepo)

			service := &placeServiceImpl{
				placeRepository: placeRepo,
				userRepository:  userRepo,
			}

			result, err := service.Create(tt.userID, tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Name, result.Name)
				assert.Equal(t, tt.expected.Address, result.Address)
				assert.Equal(t, tt.expected.Visited, result.Visited)
			}
		})
	}
}

func TestPlaceService_UpdateByID(t *testing.T) {
	t.Parallel()

	userID := "user123"
	placeID := "place123"
	userUUID := uuid.New()

	tests := []struct {
		name        string
		placeID     string
		userID      string
		input       model.PlaceUpdate
		setupMocks  func(*storage.MockPlaceRepository, *storage.MockUserRepository)
		expected    model.Place
		expectError bool
	}{
		{
			name:    "success - add visit",
			placeID: placeID,
			userID:  userID,
			input: model.PlaceUpdate{
				Name:    "Updated Place",
				Address: "Updated Address",
				Visited: true,
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)

				placeRepo.EXPECT().GetByPublicIDFull(placeID).Return(entity.Place{
					PublicID: placeID,
					Name:     "Original Place",
					Address:  "Original Address",
					Visitors: []entity.User{},
				}, nil)

				placeRepo.EXPECT().Update(mock.AnythingOfType("entity.Place")).Run(func(p entity.Place) {
					assert.Equal(t, "Updated Place", p.Name)
					assert.Equal(t, "Updated Address", p.Address)
					assert.Len(t, p.Visitors, 1)
					assert.Equal(t, userID, p.Visitors[0].PublicID)
				}).Return(nil)
			},
			expected: model.Place{
				ID:      placeID,
				Name:    "Updated Place",
				Address: "Updated Address",
				Visited: true,
			},
			expectError: false,
		},
		{
			name:    "success - remove visit",
			placeID: placeID,
			userID:  userID,
			input: model.PlaceUpdate{
				Name:    "Updated Place",
				Address: "Updated Address",
				Visited: false,
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)

				placeRepo.EXPECT().GetByPublicIDFull(placeID).Return(entity.Place{
					PublicID: placeID,
					Name:     "Original Place",
					Address:  "Original Address",
					Visitors: []entity.User{
						{PublicID: userID},
					},
				}, nil)

				placeRepo.EXPECT().Update(mock.AnythingOfType("entity.Place")).Run(func(p entity.Place) {
					assert.Equal(t, "Updated Place", p.Name)
					assert.Equal(t, "Updated Address", p.Address)
					assert.Len(t, p.Visitors, 0)
				}).Return(nil)
			},
			expected: model.Place{
				ID:      placeID,
				Name:    "Updated Place",
				Address: "Updated Address",
				Visited: false,
			},
			expectError: false,
		},
		{
			name:    "user not found",
			placeID: placeID,
			userID:  "nonexistent",
			input: model.PlaceUpdate{
				Name:    "Updated Place",
				Address: "Updated Address",
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID("nonexistent").Return(entity.User{}, errors.New("user not found"))
			},
			expected:    model.Place{},
			expectError: true,
		},
		{
			name:    "place not found",
			placeID: "nonexistent",
			userID:  userID,
			input: model.PlaceUpdate{
				Name:    "Updated Place",
				Address: "Updated Address",
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)

				placeRepo.EXPECT().GetByPublicIDFull("nonexistent").Return(entity.Place{}, errors.New("place not found"))
			},
			expected:    model.Place{},
			expectError: true,
		},
		{
			name:    "update error",
			placeID: placeID,
			userID:  userID,
			input: model.PlaceUpdate{
				Name:    "Error Place",
				Address: "Error Address",
				Visited: true,
			},
			setupMocks: func(placeRepo *storage.MockPlaceRepository, userRepo *storage.MockUserRepository) {
				userRepo.EXPECT().GetByPublicID(userID).Return(entity.User{
					ID:       userUUID,
					PublicID: userID,
					Username: "testuser",
				}, nil)

				placeRepo.EXPECT().GetByPublicIDFull(placeID).Return(entity.Place{
					PublicID: placeID,
					Name:     "Original Place",
					Address:  "Original Address",
				}, nil)

				placeRepo.EXPECT().Update(mock.AnythingOfType("entity.Place")).Return(errors.New("database error"))
			},
			expected:    model.Place{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeRepo := storage.NewMockPlaceRepository(t)
			userRepo := storage.NewMockUserRepository(t)
			tt.setupMocks(placeRepo, userRepo)

			service := &placeServiceImpl{
				placeRepository: placeRepo,
				userRepository:  userRepo,
			}

			result, err := service.UpdateByID(tt.placeID, tt.userID, tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.ID, result.ID)
				assert.Equal(t, tt.expected.Name, result.Name)
				assert.Equal(t, tt.expected.Address, result.Address)
			}
		})
	}
}
