package main

import (
	"app/pb"
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	request()
}

func newGrpcClientConn() *grpc.ClientConn {
	targetUrl := os.Getenv("USER_SERVER_NAME") + ":" + os.Getenv("GRPC_SERVER_PORT")
	conn, err := grpc.Dial(targetUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

func request() {
	router := gin.Default()
	conn := newGrpcClientConn()
	userClient := pb.NewUserServiceClient(conn)
	client := NewClient(userClient)
	router.GET("/users", client.GetUsers)
	router.Run(":" + os.Getenv("PORT"))
}

type client struct {
	user pb.UserServiceClient
}

func NewClient(userClient pb.UserServiceClient) *client {
	return &client{
		user: userClient,
	}
}

func (c *client) GetUsers(ctx *gin.Context) {
	users, err := c.user.GetUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
