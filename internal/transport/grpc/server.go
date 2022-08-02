package grpc

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"sheltes-map/internal/config"
	"sheltes-map/internal/service"
	"sheltes-map/internal/transport/grpc/handler"
	"sheltes-map/proto"
)

type IServer interface {
	ListenAndServe(conf *config.ServerConfig) error
	Stop()
}

type Deps struct {
	Logger  *logrus.Logger
	Handler *handler.Handler
}

type Server struct {
	grpc *grpc.Server
	*Deps
}

func NewGrpcServer(service *service.Service) *Server {
	deps := &Deps{
		Logger:  logrus.New(),
		Handler: handler.NewHandler(service),
	}
	server := &Server{
		grpc: grpc.NewServer(),
		Deps: deps,
	}

	return server
}

func (s *Server) ListenAndServe(conf *config.ServerConfig) error {
	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	lis, err := net.Listen(conf.Network, addr)
	if err != nil {
		return err
	}
	proto.RegisterShelterApiServer(s.grpc, s.Handler.ShelterHandler)

	s.Deps.Logger.Infof("Server listening on: %s", addr)

	if err := s.grpc.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.grpc.GracefulStop()
}
