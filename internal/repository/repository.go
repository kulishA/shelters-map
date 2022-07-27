package repository

import "sheltes-map/internal/domain"

type IShelterRepository interface {
	Save(shelter *domain.Shelter) (bool, error)
}

type Repository struct {
	Shelter IShelterRepository
}

func NewRepository(db *Database) *Repository {
	return &Repository{Shelter: NewShelterRepository(db)}
}
