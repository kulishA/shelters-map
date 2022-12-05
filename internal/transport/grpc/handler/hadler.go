package handler

import (
	"sheltes-map/internal/service"
	"sheltes-map/proto"
)

type Handler struct {
	ShelterHandler proto.LocationApiServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{ShelterHandler: NewShelterHandler(service)}
}
