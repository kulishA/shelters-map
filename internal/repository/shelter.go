package repository

import (
	"fmt"
	"sheltes-map/internal/domain"
)

type ShelterRepository struct {
	db *Database
}

func NewShelterRepository(db *Database) *ShelterRepository {
	return &ShelterRepository{db: db}
}

func (r *ShelterRepository) Save(shelter *domain.Shelter) (bool, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (city, address, address_number, latitude, longitude, closer_type, shelter_type, building_type, owner, phone, ramp) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", shelterTable)

	row := r.db.DB.QueryRow(query, shelter.City, shelter.Address, shelter.AddressNumber,
		shelter.Latitude, shelter.Longitude, shelter.CloserType, shelter.ShelterType,
		shelter.BuildingType, shelter.Owner, shelter.Phone, shelter.Ramp)

	if err := row.Scan(&id); err != nil {
		return false, err
	}

	return true, nil
}
