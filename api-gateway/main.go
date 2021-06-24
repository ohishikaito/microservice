package main

import (
	"app/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	fmt.Println("hello world")
	request()
}

func request() {
	conn := newGrpcClientConn()
	client := pb.NewUserServiceClient(conn)
	result, err := client.GetUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println("--------- response data ---------")
	fmt.Println(result)
	fmt.Println("---------------------------------")
}

func newGrpcClientConn() *grpc.ClientConn {
	// targetUrl := os.Getenv("USER_SERVER_NAME") + ":" + os.Getenv("GRPC_SERVER_PORT")
	targetUrl := "user_app_1:50051"
	conn, err := grpc.Dial(targetUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}
