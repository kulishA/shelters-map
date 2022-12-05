package repository

import (
	"github.com/jmoiron/sqlx"
	"sheltes-map/internal/core"
)

type IShelterRepository interface {
	GetNearestShelter(latitude, longitude float32) (*core.Shelter, error)
}

type IHeatingPointRepository interface {
	GetNearestHeatingPoint(latitude, longitude float32) (*core.HeatingPoint, error)
}

type Repository struct {
	Shelter      IShelterRepository
	HeatingPoint IHeatingPointRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Shelter:      NewShelterRepository(db),
		HeatingPoint: NewHeatingPointRepository(db),
	}
}
