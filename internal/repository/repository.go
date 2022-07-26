package repository

type IShelterRepository interface {
	Save() (bool, error)
}

type Repository struct {
	Shelter IShelterRepository
}

func NewRepository(db *Database) *Repository {
	return &Repository{Shelter: NewShelterRepository(db)}
}
