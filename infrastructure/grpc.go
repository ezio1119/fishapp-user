package infrastructure

import (
	"github.com/ezio1119/fishapp-auth/infrastructure/middleware"
	"github.com/ezio1119/fishapp-auth/pb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(middLe *middleware.Middleware, uc pb.AuthServiceServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middLe.LoggerInterceptor(),
			middLe.ValidatorInterceptor(),
			middLe.AuthInterceptor(),
			// middLe.RecoveryInterceptor(),
		)),
	)
	pb.RegisterAuthServiceServer(server, uc)
	reflection.Register(server)
	return server
}
