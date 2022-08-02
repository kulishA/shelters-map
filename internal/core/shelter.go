package core

import "sheltes-map/proto"

type Shelter struct {
	City          string  `db:"city"`
	Address       string  `db:"address"`
	AddressNumber string  `db:"address_number"`
	Latitude      float32 `db:"latitude"`
	Longitude     float32 `db:"longitude"`
	CloserType    string  `db:"closer_type"`
	ShelterType   string  `db:"shelter_type"`
	BuildingType  string  `db:"building_type"`
	Owner         string  `db:"owner"`
	Phone         string  `db:"phone"`
	Ramp          string  `db:"ramp"`
}

func (s *Shelter) ToResponse() *proto.ShelterResponse {
	return &proto.ShelterResponse{
		City:          s.City,
		Address:       s.Address,
		AddressNumber: s.AddressNumber,
		Latitude:      s.Latitude,
		Longitude:     s.Longitude,
		CloserType:    s.CloserType,
		ShelterType:   s.ShelterType,
		BuildingType:  s.BuildingType,
		Owner:         s.Owner,
		Phone:         s.Phone,
		Ramp:          s.Ramp,
	}
}
