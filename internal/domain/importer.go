package domain

type Shelter struct {
	City          string `csv:"city"`
	District      string `csv:"district"`
	Address       string `csv:"address"`
	AddressNumber string `csv:"address_number"`
	Latitude      string `csv:"latitude"`
	Longitude     string `csv:"longitude"`
	CloserType    string `csv:"closure_type"`
	ShelterType   string `csv:"shelter_type"`
	BuildingType  string `csv:"building_type"`
	Owner         string `csv:"owner"`
	Phone         string `csv:"phone"`
	Ramp          string `csv:"ramp"`
}
