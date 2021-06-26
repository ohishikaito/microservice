package infrastructure

import (
	"os"

	"google.golang.org/grpc"
)

func NewGrpcClientConn(serverName string) *grpc.ClientConn {
	target := serverName + ":" + os.Getenv("GRPC_SERVER_PORT")
	grpcClientConn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return grpcClientConn
}
