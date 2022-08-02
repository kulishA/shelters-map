package repository

import (
	"fmt"
	"sheltes-map/internal/core"
	"sheltes-map/pkg/database"
)

type ShelterRepository struct {
	db *database.Database
}

func NewShelterRepository(db *database.Database) *ShelterRepository {
	return &ShelterRepository{db: db}
}

func (r *ShelterRepository) GetNearestShelter(latitude, longitude float32) (*core.Shelter, error) {
	var shelter core.Shelter
	query := fmt.Sprintf("SELECT city, address, address_number, latitude, longitude, closer_type, shelter_type, building_type, owner, phone, ramp "+
		"FROM %s ORDER BY location <-> point '(%f, %f)' LIMIT 1", shelterTable, latitude, longitude)

	row := r.db.DB.QueryRow(query)

	if err := row.Scan(&shelter.City, &shelter.Address, &shelter.AddressNumber, &shelter.Latitude, &shelter.Longitude,
		&shelter.CloserType, &shelter.ShelterType, &shelter.BuildingType, &shelter.Owner, &shelter.Phone, &shelter.Ramp); err != nil {
		return nil, err
	}

	return &shelter, nil
}
