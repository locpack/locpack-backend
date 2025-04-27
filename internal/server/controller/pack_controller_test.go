package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"locpack-backend/internal/server/dto"
	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
)

func TestPackController_GetPacksByQuery(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		query          string
		mockSetup      func(*service.MockPackService)
		expectedStatus int
		expectedBody   dto.ResponseWrapper
	}{
		{
			name:           "missing user id in context",
			expectedStatus: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:           "missing query in query params",
			userID:         "user1",
			expectedStatus: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "service returns error",
			userID: "user1",
			query:  "test",
			mockSetup: func(mock *service.MockPackService) {
				mock.On("GetByNameOrAuthor", "test", "user1").Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "success",
			userID: "user1",
			query:  "test",
			mockSetup: func(mock *service.MockPackService) {
				packs := []model.Pack{
					{
						ID:   "pack1",
						Name: "Test Pack",
					},
				}

				mock.On("GetByNameOrAuthor", "test", "user1").Return(packs, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: []dto.Pack{
					{
						ID:   "pack1",
						Name: "Test Pack",
					},
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockService service.MockPackService
			if tt.mockSetup != nil {
				tt.mockSetup(&mockService)
			}

			ctx, recorder := setupControllerTest(t, "GET", "/api/v1/packs?query="+tt.query, nil)
			ctx.Set("userID", tt.userID)

			controller := NewPackController(&mockService)

			controller.GetPacksByQuery(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualPacks []dto.Pack
				err = json.Unmarshal(dataBytes, &actualPacks)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Data, actualPacks)
			} else {
				assert.Nil(t, body.Data)
			}

			assert.Equal(t, tt.expectedBody.Meta, body.Meta)
			assert.Equal(t, tt.expectedBody.Errors, body.Errors)
			assert.Equal(t, tt.expectedStatus, recorder.Code)

			mockService.AssertExpectations(t)
		})
	}
}

func TestPackController_PostPack(t *testing.T) {
	t.Parallel()

	validInput := dto.PackCreate{
		Name: "Test Pack",
	}

	testCases := []struct {
		name             string
		userID           string
		requestBody      any
		mockSetup        func(s *service.MockPackService)
		expectedBody     dto.ResponseWrapper
		expectedCode     int
		overrideBindJSON bool
	}{
		{
			name:         "missing userID",
			userID:       "",
			requestBody:  validInput,
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:             "invalid JSON",
			userID:           "123",
			overrideBindJSON: true,
			expectedCode:     http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:        "service failure",
			userID:      "123",
			requestBody: validInput,
			mockSetup: func(s *service.MockPackService) {
				s.On("Create", "123", mock.Anything).Return(model.Pack{}, errors.New("service error"))
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:        "success",
			userID:      "123",
			requestBody: validInput,
			mockSetup: func(s *service.MockPackService) {
				s.On("Create", "123", mock.Anything).Return(model.Pack{Name: "Test Pack"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.Pack{
					Name: "Test Pack",
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.MockPackService)
			controller := NewPackController(mockService)

			var requestBody io.Reader
			if tt.requestBody != nil {
				bodyBytes, err := json.Marshal(tt.requestBody)
				assert.NoError(t, err)
				requestBody = bytes.NewBuffer(bodyBytes)
			}

			ctx, recorder := setupControllerTest(t, "POST", "/api/v1/packs", requestBody)

			if tt.userID != "" {
				ctx.Set("userID", tt.userID)
			}

			if tt.overrideBindJSON {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte("{invalid-json")))
			}

			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			controller.PostPack(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualPack dto.Pack
				err = json.Unmarshal(dataBytes, &actualPack)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Data.(dto.Pack).Name, actualPack.Name)
			} else {
				assert.Nil(t, body.Data)
			}

			assert.Equal(t, tt.expectedBody.Meta, body.Meta)
			assert.Equal(t, tt.expectedBody.Errors, body.Errors)
			assert.Equal(t, tt.expectedCode, recorder.Code)

			mockService.AssertExpectations(t)
		})
	}
}

func TestPackController_GetPacksFollowed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         string
		mockSetup      func(*service.MockPackService)
		expectedStatus int
		expectedBody   dto.ResponseWrapper
	}{
		{
			name:           "missing user id in context",
			expectedStatus: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "service returns error",
			userID: "user1",
			mockSetup: func(mock *service.MockPackService) {
				mock.On("GetFollowedByUserID", "user1").Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "success",
			userID: "user1",
			mockSetup: func(mock *service.MockPackService) {
				packs := []model.Pack{
					{
						ID:   "pack1",
						Name: "Test Pack",
					},
				}

				mock.On("GetFollowedByUserID", "user1").Return(packs, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: []dto.Pack{
					{
						ID:   "pack1",
						Name: "Test Pack",
					},
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockService service.MockPackService
			if tt.mockSetup != nil {
				tt.mockSetup(&mockService)
			}

			ctx, recorder := setupControllerTest(t, "GET", "/api/v1/packs/followed", nil)
			ctx.Set("userID", tt.userID)

			controller := NewPackController(&mockService)

			controller.GetPacksFollowed(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualPacks []dto.Pack
				err = json.Unmarshal(dataBytes, &actualPacks)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Data, actualPacks)
			} else {
				assert.Nil(t, body.Data)
			}

			assert.Equal(t, tt.expectedBody.Meta, body.Meta)
			assert.Equal(t, tt.expectedBody.Errors, body.Errors)
			assert.Equal(t, tt.expectedStatus, recorder.Code)

			mockService.AssertExpectations(t)
		})
	}
}

func TestPackController_GetPacksCreated(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         string
		mockSetup      func(*service.MockPackService)
		expectedStatus int
		expectedBody   dto.ResponseWrapper
	}{
		{
			name:           "missing user id in context",
			expectedStatus: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "service returns error",
			userID: "user1",
			mockSetup: func(mock *service.MockPackService) {
				mock.On("GetCreatedByUserID", "user1").Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "success",
			userID: "user1",
			mockSetup: func(mock *service.MockPackService) {
				packs := []model.Pack{
					{
						ID:   "pack1",
						Name: "Test Pack",
					},
				}

				mock.On("GetCreatedByUserID", "user1").Return(packs, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: []dto.Pack{
					{
						ID:   "pack1",
						Name: "Test Pack",
					},
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockService service.MockPackService
			if tt.mockSetup != nil {
				tt.mockSetup(&mockService)
			}

			ctx, recorder := setupControllerTest(t, "GET", "/api/v1/packs/created", nil)
			ctx.Set("userID", tt.userID)

			controller := NewPackController(&mockService)

			controller.GetPacksCreated(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualPacks []dto.Pack
				err = json.Unmarshal(dataBytes, &actualPacks)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Data, actualPacks)
			} else {
				assert.Nil(t, body.Data)
			}

			assert.Equal(t, tt.expectedBody.Meta, body.Meta)
			assert.Equal(t, tt.expectedBody.Errors, body.Errors)
			assert.Equal(t, tt.expectedStatus, recorder.Code)

			mockService.AssertExpectations(t)
		})
	}
}

func TestPackController_GetPackByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		userID       string
		packID       string
		mockSetup    func(s *service.MockPackService)
		expectedBody dto.ResponseWrapper
		expectedCode int
	}{
		{
			name:         "missing userID",
			userID:       "",
			packID:       "456",
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:         "missing packID",
			userID:       "123",
			packID:       "",
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "service error",
			userID: "123",
			packID: "456",
			mockSetup: func(s *service.MockPackService) {
				s.On("GetByID", "456", "123").Return(model.Pack{}, errors.New("service error"))
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:   "success",
			userID: "123",
			packID: "456",
			mockSetup: func(s *service.MockPackService) {
				s.On("GetByID", "456", "123").Return(model.Pack{Name: "My Pack"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.Pack{Name: "My Pack"},
				Meta: dto.Meta{Success: true},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.MockPackService)
			controller := NewPackController(mockService)

			ctx, recorder := setupControllerTest(t, "GET", "/api/v1/packs/"+tt.packID, nil)

			if tt.userID != "" {
				ctx.Set("userID", tt.userID)
			}
			if tt.packID != "" {
				ctx.Params = gin.Params{gin.Param{Key: "id", Value: tt.packID}}
			}
			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			controller.GetPackByID(ctx)

			var body dto.ResponseWrapper
			err := json.NewDecoder(recorder.Body).Decode(&body)
			assert.NoError(t, err)

			if tt.expectedBody.Data != nil {
				expected := tt.expectedBody.Data.(dto.Pack)
				actualBytes, _ := json.Marshal(body.Data)
				var actual dto.Pack
				_ = json.Unmarshal(actualBytes, &actual)
				assert.Equal(t, expected.Name, actual.Name)
			} else {
				assert.Nil(t, body.Data)
			}

			assert.Equal(t, tt.expectedBody.Meta, body.Meta)
			assert.Equal(t, tt.expectedBody.Errors, body.Errors)
			assert.Equal(t, tt.expectedCode, recorder.Code)

			mockService.AssertExpectations(t)
		})
	}
}

func TestPackController_PutPackByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		userID           string
		packID           string
		requestBody      any
		mockSetup        func(s *service.MockPackService)
		expectedCode     int
		expectedBody     dto.ResponseWrapper
		overrideBindJSON bool
	}{
		{
			name:         "missing userID",
			userID:       "",
			packID:       "123",
			requestBody:  dto.PackUpdate{Name: "Updated"},
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:         "missing packID",
			userID:       "456",
			packID:       "",
			requestBody:  dto.PackUpdate{Name: "Updated"},
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:             "invalid JSON",
			userID:           "123",
			packID:           "456",
			overrideBindJSON: true,
			expectedCode:     http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:        "service error",
			userID:      "456",
			packID:      "123",
			requestBody: dto.PackUpdate{Name: "Updated"},
			mockSetup: func(s *service.MockPackService) {
				s.On("UpdateByID", "123", "456", mock.Anything).Return(model.Pack{}, errors.New("update failed"))
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:        "success",
			userID:      "456",
			packID:      "123",
			requestBody: dto.PackUpdate{Name: "Updated"},
			mockSetup: func(s *service.MockPackService) {
				s.On("UpdateByID", "123", "456", mock.Anything).Return(model.Pack{Name: "Updated"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.Pack{Name: "Updated"},
				Meta: dto.Meta{Success: true},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.MockPackService)
			controller := NewPackController(mockService)

			var requestBody io.Reader
			if tt.requestBody != nil {
				bodyBytes, err := json.Marshal(tt.requestBody)
				assert.NoError(t, err)
				requestBody = bytes.NewBuffer(bodyBytes)
			}

			ctx, recorder := setupControllerTest(t, "PUT", "/api/v1/packs/"+tt.packID, requestBody)

			if tt.overrideBindJSON {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte("{invalid-json")))
			}

			if tt.userID != "" {
				ctx.Set("userID", tt.userID)
			}
			if tt.packID != "" {
				ctx.Params = gin.Params{gin.Param{Key: "id", Value: tt.packID}}
			}
			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			controller.PutPackByID(ctx)

			var body dto.ResponseWrapper
			err := json.NewDecoder(recorder.Body).Decode(&body)
			assert.NoError(t, err)

			if tt.expectedBody.Data != nil {
				expected := tt.expectedBody.Data.(dto.Pack)
				dataBytes, _ := json.Marshal(body.Data)
				var actual dto.Pack
				_ = json.Unmarshal(dataBytes, &actual)
				assert.Equal(t, expected.Name, actual.Name)
			} else {
				assert.Nil(t, body.Data)
			}

			assert.Equal(t, tt.expectedBody.Meta, body.Meta)
			assert.Equal(t, tt.expectedBody.Errors, body.Errors)
			assert.Equal(t, tt.expectedCode, recorder.Code)

			mockService.AssertExpectations(t)
		})
	}
}
