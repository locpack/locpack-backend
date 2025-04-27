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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserController_GetUserMy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         string
		mockSetup      func(*service.MockUserService)
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
			mockSetup: func(mock *service.MockUserService) {
				mock.On("GetByID", "user1").Return(model.User{}, errors.New("service error"))
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
			mockSetup: func(mock *service.MockUserService) {
				user := model.User{
					ID: "pack1",
				}

				mock.On("GetByID", "user1").Return(user, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.User{
					ID: "pack1",
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockService service.MockUserService
			if tt.mockSetup != nil {
				tt.mockSetup(&mockService)
			}

			ctx, recorder := setupControllerTest(t, http.MethodGet, "/api/v1/users/my", nil)
			ctx.Set("userID", tt.userID)

			controller := NewUserController(&mockService)

			controller.GetUserMy(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualUser dto.User
				err = json.Unmarshal(dataBytes, &actualUser)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Data, actualUser)
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

func TestUserController_GetUserByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         string
		mockSetup      func(*service.MockUserService)
		expectedStatus int
		expectedBody   dto.ResponseWrapper
	}{
		{
			name:   "service returns error",
			userID: "user1",
			mockSetup: func(m *service.MockUserService) {
				m.On("GetByID", mock.Anything).Return(model.User{}, errors.New("service error"))
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
			mockSetup: func(m *service.MockUserService) {
				user := model.User{
					ID: "pack1",
				}

				m.On("GetByID", mock.Anything).Return(user, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.User{
					ID: "pack1",
				},
				Meta:   dto.Meta{Success: true},
				Errors: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockService service.MockUserService
			if tt.mockSetup != nil {
				tt.mockSetup(&mockService)
			}

			ctx, recorder := setupControllerTest(t, http.MethodGet, "/api/v1/users/"+tt.userID, nil)
			ctx.Set("userID", tt.userID)

			controller := NewUserController(&mockService)

			controller.GetUserByID(ctx)

			var body dto.ResponseWrapper
			assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&body))

			if tt.expectedBody.Data != nil {
				dataBytes, err := json.Marshal(body.Data)
				assert.NoError(t, err)

				var actualUser dto.User
				err = json.Unmarshal(dataBytes, &actualUser)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Data, actualUser)
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

func TestUserController_PutUserByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		userID           string
		requestBody      any
		mockSetup        func(s *service.MockUserService)
		expectedCode     int
		expectedBody     dto.ResponseWrapper
		overrideBindJSON bool
	}{
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
			name:        "service error",
			userID:      "456",
			requestBody: dto.UserUpdate{Username: "Updated"},
			mockSetup: func(s *service.MockUserService) {
				s.On("UpdateByID", mock.Anything, mock.Anything).Return(model.User{}, errors.New("update failed"))
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
			requestBody: dto.UserUpdate{Username: "Updated"},
			mockSetup: func(s *service.MockUserService) {
				s.On("UpdateByID", mock.Anything, mock.Anything).Return(model.User{Username: "Updated"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: dto.ResponseWrapper{
				Data: dto.User{Username: "Updated"},
				Meta: dto.Meta{Success: true},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service.MockUserService)
			controller := NewUserController(mockService)

			var requestBody io.Reader
			if tt.requestBody != nil {
				bodyBytes, err := json.Marshal(tt.requestBody)
				assert.NoError(t, err)
				requestBody = bytes.NewBuffer(bodyBytes)
			}

			ctx, recorder := setupControllerTest(t, http.MethodPut, "/api/v1/users/"+tt.userID, requestBody)

			if tt.overrideBindJSON {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte("{invalid-json")))
			}

			if tt.userID != "" {
				ctx.Set("userID", tt.userID)
			}
			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			controller.PutUserByID(ctx)

			var body dto.ResponseWrapper
			err := json.NewDecoder(recorder.Body).Decode(&body)
			assert.NoError(t, err)

			if tt.expectedBody.Data != nil {
				expected := tt.expectedBody.Data.(dto.User)
				dataBytes, _ := json.Marshal(body.Data)
				var actual dto.User
				assert.NoError(t, json.Unmarshal(dataBytes, &actual))
				assert.Equal(t, expected, actual)
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
