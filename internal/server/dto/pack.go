package dto

import "locpack-backend/pkg/types"

type Pack struct {
	ID     string           `json:"id"`
	Name   string           `json:"name"`
	Status types.PackStatus `json:"status"`
	Author User             `json:"author"`
	Places []Place          `json:"places"`
}

type PackCreate struct {
	Name string `json:"name"`
}

type PackUpdate struct {
	Name      string           `json:"name"`
	PlacesIDs []string         `json:"places_ids"`
	Status    types.PackStatus `json:"status"`
}
