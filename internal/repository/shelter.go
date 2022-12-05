package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sheltes-map/internal/core"
)

type ShelterRepository struct {
	*sqlx.DB
}

func NewShelterRepository(DB *sqlx.DB) *ShelterRepository {
	return &ShelterRepository{DB: DB}
}

func (r *ShelterRepository) GetNearestShelter(latitude, longitude float32) (*core.Shelter, error) {
	var shelter core.Shelter
	query := fmt.Sprintf("SELECT city, address, address_number, latitude, longitude, closer_type, shelter_type, building_type, owner, phone, ramp "+
		"FROM %s ORDER BY location <-> point '(%f, %f)' LIMIT 1", shelterTable, latitude, longitude)

	row := r.QueryRow(query)

	if err := row.Scan(&shelter.City, &shelter.Address, &shelter.AddressNumber, &shelter.Latitude, &shelter.Longitude,
		&shelter.CloserType, &shelter.ShelterType, &shelter.BuildingType, &shelter.Owner, &shelter.Phone, &shelter.Ramp); err != nil {
		return nil, err
	}

	return &shelter, nil
}
