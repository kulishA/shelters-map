package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sheltes-map/internal/service"
	"sheltes-map/proto"
)

type LocationHandler struct {
	services *service.Service
	proto.LocationApiServer
}

func NewLocationHandler(services *service.Service) *LocationHandler {
	return &LocationHandler{services: services}
}

func (h *LocationHandler) GetFirstNearbyShelter(ctx context.Context, in *proto.LocationRequest) (*proto.ShelterResponse, error) {
	shelter, err := h.services.Shelter.GetNearestShelter(in.GetLatitude(), in.GetLongitude())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("error: %s", err.Error()))
	}

	return shelter.ToResponse(), nil
}

func (h *LocationHandler) GetFirstNearbyHeatingPoint(ctx context.Context, in *proto.LocationRequest) (*proto.HeatingPointResponse, error) {
	heatingPoint, err := h.services.HeatingPoint.GerNearestHeatingPoint(in.GetLatitude(), in.GetLongitude())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("error: %s", err.Error()))
	}

	return heatingPoint.ToResponse(), nil
}
