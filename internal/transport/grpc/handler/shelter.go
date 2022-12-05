package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sheltes-map/internal/service"
	"sheltes-map/proto"
)

type ShelterHandler struct {
	services *service.Service
	proto.LocationApiServer
}

func NewShelterHandler(services *service.Service) *ShelterHandler {
	return &ShelterHandler{services: services}
}

func (h *ShelterHandler) GetFirstNearbyShelter(ctx context.Context, in *proto.LocationRequest) (*proto.ShelterResponse, error) {
	shelter, err := h.services.Shelter.GetNearestShelter(in.GetLatitude(), in.GetLongitude())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("error: %s", err.Error()))
	}

	return shelter.ToResponse(), nil
}
