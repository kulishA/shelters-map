package service

import (
	"sheltes-map/internal/core"
	"sheltes-map/internal/repository"
)

type IShelterService interface {
	GetNearestShelter(latitude, longitude float32) (*core.Shelter, error)
}

type IHeatingPointService interface {
	GerNearestHeatingPoint(latitude, longitude float32) (*core.HeatingPoint, error)
}

type Service struct {
	Shelter      IShelterService
	HeatingPoint IHeatingPointService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Shelter:      NewShelterService(repo.Shelter),
		HeatingPoint: NewHeatingPointService(repo.HeatingPoint),
	}
}
