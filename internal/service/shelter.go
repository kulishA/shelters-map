package service

import (
	"sheltes-map/internal/core"
	"sheltes-map/internal/repository"
)

type ShelterService struct {
	repo repository.IShelterRepository
}

func NewShelterService(repo repository.IShelterRepository) *ShelterService {
	return &ShelterService{repo: repo}
}

func (s *ShelterService) GetNearestShelter(latitude, longitude float32) (*core.Shelter, error) {
	return s.repo.GetNearestShelter(latitude, longitude)
}
