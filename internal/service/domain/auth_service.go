package domain

import (
	"strings"

	"github.com/jinzhu/copier"
	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/adapter"
)

type authServiceImpl struct {
	auth       adapter.Auth
	repository storage.UserRepository
}

func NewAuthService(auth adapter.Auth, repository storage.UserRepository) service.AuthService {
	return &authServiceImpl{auth, repository}
}

func (s *authServiceImpl) Register(register model.Register) (model.AccessToken, error) {
	userID, err := s.auth.Register(register.Username, register.Email, register.Password)
	if err != nil {
		return model.AccessToken{}, err
	}

	userEntity := entity.User{
		ID:       userID,
		PublicID: strings.ToLower(register.Username),
		Username: strings.ToLower(register.Username),
	}

	err = s.repository.Create(userEntity)
	if err != nil {
		return model.AccessToken{}, err
	}

	token, err := s.auth.Login(register.Username, register.Password)
	if err != nil {
		return model.AccessToken{}, err
	}

	accessToken := model.AccessToken{}
	err = copier.Copy(&accessToken, &token)
	if err != nil {
		return model.AccessToken{}, err
	}

	return accessToken, err
}

func (s *authServiceImpl) Login(login model.Login) (model.AccessToken, error) {
	token, err := s.auth.Login(login.Username, login.Password)
	if err != nil {
		return model.AccessToken{}, err
	}

	accessToken := model.AccessToken{}
	err = copier.Copy(&accessToken, &token)
	if err != nil {
		return model.AccessToken{}, err
	}

	return accessToken, err
}

func (s *authServiceImpl) Refresh(refresh model.Refresh) (model.AccessToken, error) {
	token, err := s.auth.Refresh(refresh.Value)
	if err != nil {
		return model.AccessToken{}, err
	}

	accessToken := model.AccessToken{}
	err = copier.Copy(&accessToken, &token)
	if err != nil {
		return model.AccessToken{}, err
	}

	return accessToken, err
}
