package infrastructure

import (
	"context"
	"fmt"

	"github.com/ezio1119/fishapp-user/infrastructure/middleware"
	"github.com/ezio1119/fishapp-user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(middLe *middleware.Middleware, uc pb.UserServiceServer) *grpc.Server {

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middLe.UnaryLogingInterceptor(),
			middLe.UnaryValidationInterceptor(),
			middLe.UnaryAuthInterceptor(),
			middLe.UnaryRecoveryInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			middLe.StreamLogingInterceptor(),
			middLe.StreamRecoveryInterceptor(),
		),
	)

	pb.RegisterUserServiceServer(server, uc)
	reflection.Register(server)
	return server
}

func ErrorUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("インター")
		return handler(ctx, req)
	})
}
