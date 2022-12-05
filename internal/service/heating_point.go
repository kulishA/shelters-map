package service

import (
	"sheltes-map/internal/core"
	"sheltes-map/internal/repository"
)

type HeatingPointService struct {
	repo repository.IHeatingPointRepository
}

func NewHeatingPointService(repo repository.IHeatingPointRepository) *HeatingPointService {
	return &HeatingPointService{repo: repo}
}

func (s *HeatingPointService) GerNearestHeatingPoint(latitude, longitude float32) (*core.HeatingPoint, error) {
	return s.repo.GetNearestHeatingPoint(latitude, longitude)
}
