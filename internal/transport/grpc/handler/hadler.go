package handler

import (
	"sheltes-map/internal/service"
	"sheltes-map/proto"
)

type Handler struct {
	ShelterHandler proto.ShelterApiServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{ShelterHandler: NewShelterHandler(service)}
}
