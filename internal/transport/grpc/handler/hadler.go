package handler

import (
	"sheltes-map/internal/service"
	"sheltes-map/proto"
)

type Handler struct {
	LocationHandler proto.LocationApiServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{LocationHandler: NewLocationHandler(service)}
}
