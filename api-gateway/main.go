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
	userController := NewUserController(userClient)
	router.GET("/users", userController.GetUsers)
	router.Run(":" + os.Getenv("PORT"))
}

type userController struct {
	userClient pb.UserServiceClient
}

func NewUserController(userClient pb.UserServiceClient) *userController {
	return &userController{
		userClient: userClient,
	}
}

func (c *userController) GetUsers(ctx *gin.Context) {
	users, err := c.userClient.GetUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
