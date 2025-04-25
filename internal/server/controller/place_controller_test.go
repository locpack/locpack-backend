package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"locpack-backend/internal/server/dto"
	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPlaceController_GetPlacesByQuery(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		query          string
		mockSetup      func(*service.MockPlaceService)
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
			mockSetup: func(mock *service.MockPlaceService) {
				mock.On("GetByNameOrAddress", "test", "user1").Return(nil, errors.New("service error"))
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
			mockSetup: func(mock *service.MockPlaceService) {
				packs := []model.Place{
					{
						ID:   "place1",
						Name: "Test Place",
					},
				}

				mock.On("GetByNameOrAddress", "test", "user1").Return(packs, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: []dto.Place{
					{
						ID:   "place1",
						Name: "Test Place",
					},
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockService service.MockPlaceService
			if tt.mockSetup != nil {
				tt.mockSetup(&mockService)
			}

			ctx, recorder := setupTestContext(t, "GET", "/api/v1/places?query="+tt.query, nil)
			ctx.Set("userID", tt.userID)

			controller := NewPlaceController(&mockService)
			controller.GetPlacesByQuery(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualPacks []dto.Place
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

func TestPlaceController_PostPlace(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name             string
		userID           string
		requestBody      any
		mockSetup        func(s *service.MockPlaceService)
		expectedBody     dto.ResponseWrapper
		expectedCode     int
		overrideBindJSON bool
	}

	validInput := dto.PlaceCreate{
		Name: "Test Place",
	}

	testCases := []testCase{
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
			mockSetup: func(s *service.MockPlaceService) {
				s.On("Create", "123", mock.Anything).Return(model.Place{}, errors.New("service error"))
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
			mockSetup: func(s *service.MockPlaceService) {
				s.On("Create", "123", mock.Anything).Return(model.Place{Name: "Test Place"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.Place{
					Name: "Test Place",
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.MockPlaceService)
			controller := NewPlaceController(mockService)

			var requestBody io.Reader
			if tt.requestBody != nil {
				bodyBytes, err := json.Marshal(tt.requestBody)
				assert.NoError(t, err)
				requestBody = bytes.NewBuffer(bodyBytes)
			}

			ctx, recorder := setupTestContext(t, "POST", "/api/v1/places", requestBody)

			if tt.userID != "" {
				ctx.Set("userID", tt.userID)
			}

			if tt.overrideBindJSON {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte("{invalid-json")))
			}

			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			controller.PostPlace(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualPack dto.Place
				err = json.Unmarshal(dataBytes, &actualPack)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Data.(dto.Place).Name, actualPack.Name)
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

func TestPlaceController_GetPlaceByID(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name         string
		userID       string
		placeID      string
		mockSetup    func(s *service.MockPlaceService)
		expectedBody dto.ResponseWrapper
		expectedCode int
	}

	testCases := []testCase{
		{
			name:         "missing userID",
			userID:       "",
			placeID:      "456",
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:         "missing placeID",
			userID:       "123",
			placeID:      "",
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:    "service error",
			userID:  "123",
			placeID: "456",
			mockSetup: func(s *service.MockPlaceService) {
				s.On("GetByID", "456", "123").Return(model.Place{}, errors.New("service error"))
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:    "success",
			userID:  "123",
			placeID: "456",
			mockSetup: func(s *service.MockPlaceService) {
				s.On("GetByID", "456", "123").Return(model.Place{Name: "My Place"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.Place{Name: "My Place"},
				Meta: dto.Meta{Success: true},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.MockPlaceService)
			controller := NewPlaceController(mockService)

			ctx, recorder := setupTestContext(t, "GET", "/api/v1/places/"+tt.placeID, nil)

			if tt.userID != "" {
				ctx.Set("userID", tt.userID)
			}
			if tt.placeID != "" {
				ctx.Params = gin.Params{gin.Param{Key: "id", Value: tt.placeID}}
			}
			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			controller.GetPlaceByID(ctx)

			var body dto.ResponseWrapper
			err := json.NewDecoder(recorder.Body).Decode(&body)
			assert.NoError(t, err)

			if tt.expectedBody.Data != nil {
				expected := tt.expectedBody.Data.(dto.Place)
				actualBytes, _ := json.Marshal(body.Data)
				var actual dto.Place
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

func TestPlaceController_PutPlaceByID(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name             string
		userID           string
		placeID          string
		requestBody      any
		mockSetup        func(s *service.MockPlaceService)
		expectedCode     int
		expectedBody     dto.ResponseWrapper
		overrideBindJSON bool
	}

	testCases := []testCase{
		{
			name:         "missing userID",
			userID:       "",
			placeID:      "123",
			requestBody:  dto.PlaceUpdate{Name: "Updated"},
			expectedCode: http.StatusBadRequest,
			expectedBody: dto.ResponseWrapper{
				Data:   nil,
				Meta:   dto.Meta{Success: false},
				Errors: []dto.Error{{Message: "Some error", Code: "000"}},
			},
		},
		{
			name:         "missing placeID",
			userID:       "456",
			placeID:      "",
			requestBody:  dto.PlaceUpdate{Name: "Updated"},
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
			placeID:          "456",
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
			placeID:     "123",
			requestBody: dto.PackUpdate{Name: "Updated"},
			mockSetup: func(s *service.MockPlaceService) {
				s.On("UpdateByID", "123", "456", mock.Anything).Return(model.Place{}, errors.New("update failed"))
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
			placeID:     "123",
			requestBody: dto.PackUpdate{Name: "Updated"},
			mockSetup: func(s *service.MockPlaceService) {
				s.On("UpdateByID", "123", "456", mock.Anything).Return(model.Place{Name: "Updated"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.Place{Name: "Updated"},
				Meta: dto.Meta{Success: true},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.MockPlaceService)
			controller := NewPlaceController(mockService)

			var requestBody io.Reader
			if tt.requestBody != nil {
				bodyBytes, err := json.Marshal(tt.requestBody)
				assert.NoError(t, err)
				requestBody = bytes.NewBuffer(bodyBytes)
			}

			ctx, recorder := setupTestContext(t, "PUT", "/api/v1/places/"+tt.placeID, requestBody)

			if tt.overrideBindJSON {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte("{invalid-json")))
			}

			if tt.userID != "" {
				ctx.Set("userID", tt.userID)
			}
			if tt.placeID != "" {
				ctx.Params = gin.Params{gin.Param{Key: "id", Value: tt.placeID}}
			}
			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			controller.PutPlaceByID(ctx)

			var body dto.ResponseWrapper
			err := json.NewDecoder(recorder.Body).Decode(&body)
			assert.NoError(t, err)

			if tt.expectedBody.Data != nil {
				expected := tt.expectedBody.Data.(dto.Place)
				dataBytes, _ := json.Marshal(body.Data)
				var actual dto.Place
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
