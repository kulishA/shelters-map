package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sheltes-map/internal/core"
)

type HeatingPointRepository struct {
	*sqlx.DB
}

func NewHeatingPointRepository(DB *sqlx.DB) *HeatingPointRepository {
	return &HeatingPointRepository{DB: DB}
}

func (r *HeatingPointRepository) GetNearestHeatingPoint(latitude, longitude float32) (*core.HeatingPoint, error) {
	var heatingPoint core.HeatingPoint

	query := fmt.Sprintf("SELECT city, address, district, latitude, longitude, closer_type, "+
		"start_date, capacity, mobile_phone, phone, responsible_person, kitchen_room, mother_room, disabilities_room "+
		"FROM %s ORDER BY location <-> point '(%f, %f)' LIMIT 1", shelterTable, latitude, longitude)

	row := r.QueryRow(query)

	if err := row.Scan(&heatingPoint.City, &heatingPoint.Address, &heatingPoint.District, &heatingPoint.Latitude, &heatingPoint.Longitude,
		&heatingPoint.CloserType, &heatingPoint.StartDate, &heatingPoint.Capacity, &heatingPoint.MobilePhone, &heatingPoint.Phone,
		&heatingPoint.ResponsiblePerson, heatingPoint.KitchenRoom, heatingPoint.MotherRoom, heatingPoint.DisabilitiesRoom); err != nil {
		return nil, err
	}

	return &heatingPoint, nil
}
