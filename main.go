package main

import (
	"context"
	"net"
	"time"

	"github.com/ezio1119/fishapp-auth/conf"
	"github.com/ezio1119/fishapp-auth/infrastructure"
	"github.com/ezio1119/fishapp-auth/infrastructure/middleware"
	"github.com/ezio1119/fishapp-auth/interfaces/controllers"
	"github.com/ezio1119/fishapp-auth/interfaces/repository"
	"github.com/ezio1119/fishapp-auth/pb"
	"github.com/ezio1119/fishapp-auth/usecase/interactor"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	dbConn, err := infrastructure.NewGormConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	redisC, err := infrastructure.NewRedisClient()
	if err != nil {
		panic(err)
	}
	defer redisC.Close()

	grpcConn, err := grpc.DialContext(ctx, conf.C.API.ProfileAPI, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()
	profileC := pb.NewProfileServiceClient(grpcConn)

	authController := controllers.NewAuthController(
		interactor.NewAuthInteractor(
			repository.NewUserRepository(dbConn),
			repository.NewProfileRepository(profileC),
			repository.NewBlackListRepository(redisC),
			time.Duration(conf.C.Sv.Timeout)*time.Second,
		))

	server := infrastructure.NewGrpcServer(middleware.InitMiddleware(), authController)

	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		panic(err)
	}
	err = server.Serve(list)
	if err != nil {
		panic(err)
	}

}
