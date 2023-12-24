package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pintoter/mts-test/order-service/internal/service"
	desc "github.com/pintoter/mts-test/order-service/pkg/api/order-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

type Config interface {
	GetGrpcAddr() string
	GetMaxConnIdle() time.Duration
	GetMaxConnAge() time.Duration
	GetTime() time.Duration
	GetTimeout() time.Duration
}

type grpcServer struct {
	service *service.Service
	desc.UnimplementedOrderServiceServer
}

func New(service *service.Service) *grpcServer {
	return &grpcServer{
		service: service,
	}
}

func (s *grpcServer) Run(cfg Config) error {
	lis, err := net.Listen("tcp", cfg.GetGrpcAddr())
	if err != nil {
		return err
	}
	defer lis.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle: cfg.GetMaxConnIdle(),
				MaxConnectionAge:  cfg.GetMaxConnAge(),
				Time:              cfg.GetTime(),
				Timeout:           cfg.GetTimeout(),
			},
		),
	)

	desc.RegisterOrderServiceServer(grpcServer, s)

	go func() {
		log.Println("gRPC server is listening on:", cfg.GetGrpcAddr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("failed running gRPC server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Printf("got signal %v, starting shut down", v)
		grpcServer.GracefulStop()
	}

	log.Println("gRPC server correctly shutdown")

	return nil
}

func (s *grpcServer) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	log.Printf("got new request from UserId: %d for buy ItemId: %d\n", req.GetUserId(), req.GetItemId())

	err := s.service.CreateOrder(ctx, req.GetUserId(), req.GetItemId())
	if err != nil {
		return &desc.CreateOrderResponse{}, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CreateOrderResponse{}, nil
}
