package domain

import (
	"testing"

	"github.com/gin-gonic/gin"
	"locpack-backend/internal/storage"
)

func setupServiceTest(t *testing.T) (*packServiceImpl, *placeServiceImpl, *userServiceImpl, *storage.MockPackRepository, *storage.MockPlaceRepository, *storage.MockUserRepository) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	packRepo := new(storage.MockPackRepository)
	placeRepo := new(storage.MockPlaceRepository)
	userRepo := new(storage.MockUserRepository)
	packSvc := NewPackService(packRepo, placeRepo, userRepo).(*packServiceImpl)
	placeSvc := NewPlaceService(placeRepo, userRepo).(*placeServiceImpl)
	userSvc := NewUserService(userRepo).(*userServiceImpl)

	return packSvc, placeSvc, userSvc, packRepo, placeRepo, userRepo
}
