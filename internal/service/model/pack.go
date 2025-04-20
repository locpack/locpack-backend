package model

import "locpack-backend/pkg/types"

type Pack struct {
	ID             string `copier:"PublicID"`
	Name           string
	AuthorID       string
	AuthorUsername string
	Status         types.PackStatus
}

type PackCreate struct {
	Name string
}

type PackUpdate struct {
	Name      string
	Status    types.PackStatus
	PlacesIDs []string
}
