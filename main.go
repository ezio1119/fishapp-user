package main

import (
	"net"
	"time"

	"github.com/ezio1119/fishapp-user/conf"
	"github.com/ezio1119/fishapp-user/infrastructure"
	"github.com/ezio1119/fishapp-user/infrastructure/middleware"
	"github.com/ezio1119/fishapp-user/interfaces/controllers"
	"github.com/ezio1119/fishapp-user/interfaces/repository"
	"github.com/ezio1119/fishapp-user/usecase/interactor"
)

func main() {
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

	userController := controllers.NewUserController(
		interactor.NewUserInteractor(
			repository.NewUserRepository(dbConn),
			repository.NewBlackListRepository(redisC),
			time.Duration(conf.C.Sv.Timeout)*time.Second,
		))

	server := infrastructure.NewGrpcServer(middleware.InitMiddleware(), userController)

	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		panic(err)
	}
	if err := server.Serve(list); err != nil {
		panic(err)
	}

}
