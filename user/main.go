package main

import (
	"app/pb"
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	userController := NewUserController()
	grpcServer := grpc.NewServer()
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

type userController struct {
	//
}

func NewUserController() *userController {
	return &userController{}
}

func (c *userController) GetUsers(ctx context.Context, req *emptypb.Empty) (*pb.GetUsersResponse, error) {
	user := &pb.User{
		Id:        1,
		LastName:  "大石",
		FirstName: "海渡",
	}
	var users []*pb.User
	users = append(users, user)
	pbUsers := &pb.GetUsersResponse{
		Users: users,
	}
	return pbUsers, nil
}
