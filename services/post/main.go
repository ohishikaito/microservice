package main

import (
	"log"
	"net"
	"os"
	"post/infrastructure"
	"post/interface/controller"
	"post/interface/repository"
	"post/pb"
	"post/usecase"
)

func main() {
	db := infrastructure.NewGormConnect()

	postRepository := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postController := controller.NewPostController(postUsecase)

	grpcServer := infrastructure.NewGrpcServer()
	pb.RegisterPostServiceServer(grpcServer, postController)

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
