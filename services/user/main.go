package main

import (
	"log"
	"net"
	"os"
	"user/infrastructure"
	"user/interface/controller"
	"user/interface/repository"
	"user/pb"
	"user/usecase"
)

func main() {
	db := infrastructure.NewGormConnect()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	grpcServer := infrastructure.NewGrpcServer()
	pb.RegisterUserServiceServer(grpcServer, userController)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVER_PORT")) // [::]:50051
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	log.Print("grpcServer serve")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("serve err", err)
	}
}
