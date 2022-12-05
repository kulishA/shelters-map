package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sheltes-map/internal/service"
	"sheltes-map/proto"
)

type HeatingPointHandler struct {
	services *service.Service
	proto.LocationApiServer
}

func NewHeatingPointHandler(services *service.Service) *HeatingPointHandler {
	return &HeatingPointHandler{services: services}
}

func (h *HeatingPointHandler) GetFirstNearbyHeatingPoint(ctx context.Context, in *proto.LocationRequest) (*proto.HeatingPointResponse, error) {
	heatingPoint, err := h.services.HeatingPoint.GerNearestHeatingPoint(in.GetLatitude(), in.GetLongitude())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("error: %s", err.Error()))
	}

	return heatingPoint.ToResponse(), nil
}
