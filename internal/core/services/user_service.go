package services

import "placelists/internal/app/api/dtos"

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (s *UserService) GetUserByUsername(username string) (*dtos.User, error) {
	user := &dtos.User{Username: username}
	return user, nil
}

func (s *UserService) UpdateUserByUsername(username string, uu dtos.UserUpdate) (*dtos.User, error) {
	return nil, nil
}
