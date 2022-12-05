package core

import "sheltes-map/proto"

type HeatingPoint struct {
	City              string  `db:"city"`
	Address           string  `db:"address"`
	District          string  `db:"district"`
	Latitude          float32 `db:"latitude"`
	Longitude         float32 `db:"longitude"`
	CloserType        string  `db:"closer_type"`
	StartDate         string  `db:"start_date"`
	Capacity          int64   `db:"capacity"`
	MobilePhone       string  `db:"mobile_phone"`
	ResponsiblePerson string  `db:"responsible_person"`
	Phone             string  `db:"phone"`
	KitchenRoom       bool    `db:"kitchen_room"`
	MotherRoom        bool    `db:"mother_room"`
	DisabilitiesRoom  bool    `db:"disabilities_room"`
}

func (s *HeatingPoint) ToResponse() *proto.HeatingPointResponse {
	return &proto.HeatingPointResponse{
		City:              s.City,
		Address:           s.Address,
		District:          s.District,
		Latitude:          s.Latitude,
		Longitude:         s.Longitude,
		CloserType:        s.CloserType,
		StartDate:         s.StartDate,
		Capacity:          s.Capacity,
		MobilePhone:       s.MobilePhone,
		ResponsiblePerson: s.ResponsiblePerson,
		Phone:             s.Phone,
		KitchenRoom:       s.KitchenRoom,
		MotherRoom:        s.MotherRoom,
		DisabilitiesRoom:  s.DisabilitiesRoom,
	}
}
