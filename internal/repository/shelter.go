package repository

type ShelterRepository struct {
	db *Database
}

func NewShelterRepository(db *Database) *ShelterRepository {
	return &ShelterRepository{db: db}
}

func (r *ShelterRepository) Save() (bool, error) {
	return false, nil
}
