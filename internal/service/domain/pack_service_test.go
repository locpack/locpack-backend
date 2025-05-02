package domain

import (
	"errors"
	"testing"

	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/enum/pack_status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPackService_GetByID(t *testing.T) {
	t.Parallel()

	packSvc, _, _, packRepo, _, _ := setupServiceTest(t)

	userID := "user1"
	packID := "pack1"
	packEntity := entity.Pack{
		PublicID: packID,
		Name:     "Test Pack",
		Author:   entity.User{PublicID: userID, Username: "author"},
		FollowedUsers: []entity.User{
			{PublicID: userID},
		},
	}

	tests := []struct {
		name     string
		setup    func()
		expected model.Pack
		wantErr  bool
	}{
		{
			name: "success - created",
			setup: func() {
				packRepo.On("GetByPublicIDFull", packID).Return(packEntity, nil).Once()
			},
			expected: model.Pack{
				ID:             packID,
				Name:           "Test Pack",
				AuthorID:       userID,
				AuthorUsername: "author",
				Status:         pack_status.Created,
			},
		},
		{
			name: "error fetching pack",
			setup: func() {
				packRepo.On("GetByPublicIDFull", packID).Return(entity.Pack{}, errors.New("not found")).Once()
			},
			wantErr: true,
		},
		{
			name: "user neither author nor follower",
			setup: func() {
				packRepo.On("GetByPublicIDFull", packID).Return(entity.Pack{
					PublicID: packID,
					Name:     "Pack",
					Author:   entity.User{PublicID: "other"},
					FollowedUsers: []entity.User{
						{PublicID: "someone-else"},
					},
				}, nil).Once()
			},
			expected: model.Pack{
				ID:             packID,
				Name:           "Pack",
				AuthorID:       "other",
				AuthorUsername: "",
				Status:         pack_status.None,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := packSvc.GetByID(packID, userID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, got)
			}

			packRepo.AssertExpectations(t)
		})
	}
}

func TestPackService_GetByNameOrAuthor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		userID     string
		query      string
		mockReturn []entity.Pack
		mockError  error
		expected   []model.Pack
		expectErr  bool
	}{
		{
			name:   "Created and Followed packs",
			userID: "user1",
			query:  "Pack",
			mockReturn: []entity.Pack{
				{
					PublicID:      "p1",
					Name:          "My Pack",
					Author:        entity.User{PublicID: "user1", Username: "u1"},
					FollowedUsers: []entity.User{{PublicID: "user2"}},
				},
				{
					PublicID:      "p2",
					Name:          "Another Pack",
					Author:        entity.User{PublicID: "user2", Username: "u2"},
					FollowedUsers: []entity.User{{PublicID: "user1"}},
				},
			},
			expected: []model.Pack{
				{Name: "My Pack", Status: pack_status.Created},
				{Name: "Another Pack", Status: pack_status.Followed},
			},
		},
		{
			name:      "Repository error",
			userID:    "user1",
			query:     "Pack",
			mockError: errors.New("db error"),
			expectErr: true,
		},
		{
			name:   "Pack with status None",
			userID: "user3",
			query:  "Pack",
			mockReturn: []entity.Pack{
				{
					PublicID: "p3",
					Name:     "Unrelated Pack",
					Author:   entity.User{PublicID: "someoneelse"},
				},
			},
			expected: []model.Pack{
				{Name: "Unrelated Pack", Status: pack_status.None},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packSvc, _, _, packRepo, _, _ := setupServiceTest(t)
			packRepo.On("GetByNameOrAuthorFull", tt.query).Return(tt.mockReturn, tt.mockError).Once()

			res, err := packSvc.GetByNameOrAuthor(tt.query, tt.userID)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, res, len(tt.expected))
			for i, p := range res {
				assert.Equal(t, tt.expected[i].Status, p.Status)
				assert.Equal(t, tt.expected[i].Name, p.Name)
			}
		})
	}
}

func TestPackService_GetFollowedByUserID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		userID    string
		mockUser  entity.User
		mockError error
		expected  int
		expectErr bool
	}{
		{
			name:   "Single followed pack",
			userID: "user1",
			mockUser: entity.User{FollowedPacks: []entity.Pack{
				{
					PublicID: "p1",
					Name:     "Followed Pack",
					Author: entity.User{
						PublicID: "u2",
						Username: "author"},
				}},
			},
			expected: 1,
		},
		{
			name:      "User not found",
			userID:    "missing",
			mockError: errors.New("not found"),
			expectErr: true,
		},
		{
			name:     "No followed packs",
			userID:   "user1",
			mockUser: entity.User{FollowedPacks: []entity.Pack{}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packSvc, _, _, _, _, userRepo := setupServiceTest(t)
			userRepo.On("GetByPublicIDFull", tt.userID).Return(tt.mockUser, tt.mockError).Once()

			res, err := packSvc.GetFollowedByUserID(tt.userID)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, res, tt.expected)
			for _, p := range res {
				assert.Equal(t, pack_status.Followed, p.Status)
			}
		})
	}
}

func TestPackService_GetCreatedByUserID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		userID    string
		mockUser  entity.User
		mockError error
		expected  int
		expectErr bool
	}{
		{
			name:   "Single created pack",
			userID: "user1",
			mockUser: entity.User{FollowedPacks: []entity.Pack{
				{
					PublicID: "p1",
					Name:     "Created Pack",
					Author: entity.User{
						PublicID: "user1",
						Username: "author",
					}},
			}},
			expected: 1,
		},
		{
			name:      "User not found",
			userID:    "missing",
			mockError: errors.New("not found"),
			expectErr: true,
		},
		{
			name:     "No created packs",
			userID:   "user1",
			mockUser: entity.User{FollowedPacks: []entity.Pack{}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packSvc, _, _, _, _, userRepo := setupServiceTest(t)
			userRepo.On("GetByPublicIDFull", tt.userID).Return(tt.mockUser, tt.mockError).Once()

			res, err := packSvc.GetCreatedByUserID(tt.userID)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, res, tt.expected)
			for _, p := range res {
				assert.Equal(t, pack_status.Created, p.Status)
			}
		})
	}
}

func TestPackService_Create(t *testing.T) {
	tests := []struct {
		name       string
		setupMocks func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository)
		wantErr    bool
	}{
		{
			name: "success",
			setupMocks: func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository) {
				userRepo.On("GetByPublicID", "user1").Return(entity.User{PublicID: "user1"}, nil)
				packRepo.On("Register", mock.AnythingOfType("entity.Pack")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "user repo error",
			setupMocks: func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository) {
				userRepo.On("GetByPublicID", "user1").Return(entity.User{}, errors.New("user not found"))
			},
			wantErr: true,
		},
		{
			name: "pack repo create error",
			setupMocks: func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository) {
				userRepo.On("GetByPublicID", "user1").Return(entity.User{PublicID: "user1"}, nil)
				packRepo.On("Register", mock.AnythingOfType("entity.Pack")).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packSvc, _, _, packRepo, _, userRepo := setupServiceTest(t)
			tt.setupMocks(userRepo, packRepo)

			pack, err := packSvc.Create("user1", model.PackCreate{Name: "My Pack"})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "My Pack", pack.Name)
			}
		})
	}
}

func TestPackService_UpdateByID(t *testing.T) {
	tests := []struct {
		name       string
		update     model.PackUpdate
		setupMocks func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository, placeRepo *storage.MockPlaceRepository)
		wantErr    bool
	}{
		{
			name: "success followed",
			update: model.PackUpdate{
				Name:      "New Name",
				Status:    pack_status.Followed,
				PlacesIDs: []string{"place1"},
			},
			setupMocks: func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository, placeRepo *storage.MockPlaceRepository) {
				userRepo.On("GetByPublicID", "user1").Return(entity.User{PublicID: "user1"}, nil)
				packRepo.On("GetByPublicIDFull", "pack1").Return(entity.Pack{
					PublicID: "pack1",
					Name:     "Old Name",
					Author:   entity.User{PublicID: "user2"},
				}, nil)
				placeRepo.On("GetByPublicID", "place1").Return(entity.Place{PublicID: "place1"}, nil)
				packRepo.On("Update", mock.AnythingOfType("entity.Pack")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "error on created status",
			update: model.PackUpdate{
				Name:   "New Name",
				Status: pack_status.Created,
			},
			setupMocks: func(_ *storage.MockUserRepository, _ *storage.MockPackRepository, _ *storage.MockPlaceRepository) {},
			wantErr:    true,
		},
		{
			name: "user repo error",
			update: model.PackUpdate{
				Name:   "New Name",
				Status: pack_status.Followed,
			},
			setupMocks: func(userRepo *storage.MockUserRepository, _ *storage.MockPackRepository, _ *storage.MockPlaceRepository) {
				userRepo.On("GetByPublicID", "user1").Return(entity.User{}, errors.New("user not found"))
			},
			wantErr: true,
		},
		{
			name: "pack repo error",
			update: model.PackUpdate{
				Name:   "New Name",
				Status: pack_status.Followed,
			},
			setupMocks: func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository, placeRepo *storage.MockPlaceRepository) {
				userRepo.On("GetByPublicID", "user1").Return(entity.User{PublicID: "user1"}, nil)
				packRepo.On("GetByPublicIDFull", "pack1").Return(entity.Pack{}, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name: "update repo error",
			update: model.PackUpdate{
				Name:      "New Name",
				Status:    pack_status.Followed,
				PlacesIDs: []string{"place1"},
			},
			setupMocks: func(userRepo *storage.MockUserRepository, packRepo *storage.MockPackRepository, placeRepo *storage.MockPlaceRepository) {
				userRepo.On("GetByPublicID", "user1").Return(entity.User{PublicID: "user1"}, nil)
				packRepo.On("GetByPublicIDFull", "pack1").Return(entity.Pack{
					PublicID: "pack1",
					Name:     "Old Name",
					Author:   entity.User{PublicID: "user2"},
				}, nil)
				placeRepo.On("GetByPublicID", "place1").Return(entity.Place{PublicID: "place1"}, nil)
				packRepo.On("Update", mock.AnythingOfType("entity.Pack")).Return(errors.New("update error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packSvc, _, _, packRepo, placeRepo, userRepo := setupServiceTest(t)
			tt.setupMocks(userRepo, packRepo, placeRepo)

			updated, err := packSvc.UpdateByID("pack1", "user1", tt.update)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "New Name", updated.Name)
			}
		})
	}
}

func TestPackService_GetPlacesByID(t *testing.T) {
	tests := []struct {
		name           string
		packID         string
		userID         string
		mockSetup      func(packRepo *storage.MockPackRepository)
		expectedPlaces []model.Place
		expectErr      bool
	}{
		{
			name:   "success - user visited one place",
			packID: "pack1",
			userID: "user1",
			mockSetup: func(packRepo *storage.MockPackRepository) {
				packRepo.On("GetByPublicIDFull", "pack1").Return(entity.Pack{
					PublicID: "pack1",
					Places: []entity.Place{
						{
							PublicID: "place1",
							Name:     "Visited Place",
							Visitors: []entity.User{
								{PublicID: "user1"},
							},
						},
					},
				}, nil)
			},
			expectedPlaces: []model.Place{
				{ID: "place1", Name: "Visited Place", Visited: true},
			},
			expectErr: false,
		},
		{
			name:   "success - user did not visit",
			packID: "pack1",
			userID: "user1",
			mockSetup: func(packRepo *storage.MockPackRepository) {
				packRepo.On("GetByPublicIDFull", "pack1").Return(entity.Pack{
					PublicID: "pack1",
					Places: []entity.Place{
						{
							PublicID: "place2",
							Name:     "Not Visited Place",
							Visitors: []entity.User{
								{PublicID: "user2"},
							},
						},
					},
				}, nil)
			},
			expectedPlaces: []model.Place{
				{ID: "place2", Name: "Not Visited Place", Visited: false},
			},
			expectErr: false,
		},
		{
			name:   "error - pack not found",
			packID: "pack404",
			userID: "user1",
			mockSetup: func(packRepo *storage.MockPackRepository) {
				packRepo.On("GetByPublicIDFull", "pack404").Return(entity.Pack{}, errors.New("not found"))
			},
			expectedPlaces: nil,
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packSvc, _, _, packRepo, _, _ := setupServiceTest(t)
			if tt.mockSetup != nil {
				tt.mockSetup(packRepo)
			}

			got, err := packSvc.GetPlacesByID(tt.packID, tt.userID)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedPlaces), len(got))
				for i := range tt.expectedPlaces {
					assert.Equal(t, tt.expectedPlaces[i].Visited, got[i].Visited)
					assert.Equal(t, tt.expectedPlaces[i].ID, got[i].ID)
					assert.Equal(t, tt.expectedPlaces[i].Name, got[i].Name)
				}
			}

			packRepo.AssertExpectations(t)
		})
	}
}

func TestPackService_UpdateByID_ErrorOnCreateStatus(t *testing.T) {
	packSvc, _, _, _, _, _ := setupServiceTest(t)

	_, err := packSvc.UpdateByID("pack1", "user1", model.PackUpdate{
		Status: pack_status.Created,
	})

	assert.Error(t, err)
}
