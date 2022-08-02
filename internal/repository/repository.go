package repository

import (
	"sheltes-map/internal/core"
	"sheltes-map/pkg/database"
)

type IShelterRepository interface {
	GetNearestShelter(latitude, longitude float32) (*core.Shelter, error)
}

type Repository struct {
	Shelter IShelterRepository
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{Shelter: NewShelterRepository(db)}
}
