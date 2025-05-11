package model

import "locpack-backend/pkg/types"

type Pack struct {
	ID     string
	Name   string
	Author User
	Status types.PackStatus
	Places []Place
}

type PackCreate struct {
	Name string
}

type PackUpdate struct {
	Name      string
	Status    types.PackStatus
	PlacesIDs []string
}
