package fakes

import (
	"errors"
	"placelists/internal/storage/entities"
)

type UserRepositoryFakeImpl struct {
	Users []entities.User
}

func NewUserRepositoryFake() *UserRepositoryFakeImpl {
	return &UserRepositoryFakeImpl{}
}

func (r *UserRepositoryFakeImpl) GetByPublicID(id string) (*entities.User, error) {
	for _, user := range r.Users {
		if user.PublicID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepositoryFakeImpl) GetByPublicIDFull(id string) (*entities.User, error) {
	user, err := r.GetByPublicID(id)
	return user, err
}

func (r *UserRepositoryFakeImpl) Create(u *entities.User) error {
	for _, user := range r.Users {
		if user.ID == u.ID {
			return errors.New("user already exists")
		}
	}
	r.Users = append(r.Users, *u)
	return nil
}

func (r *UserRepositoryFakeImpl) Update(u *entities.User) error {
	for i, user := range r.Users {
		if user.ID == u.ID {
			r.Users[i] = *u
			return nil
		}
	}
	return errors.New("user not found")
}
