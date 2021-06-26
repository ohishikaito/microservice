package main

import (
	"api_gateway/client"
	"api_gateway/infrastructure"
	"api_gateway/pb"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.NewLogger()

	grpcClientConn := infrastructure.NewGrpcClientConn(os.Getenv("USER_SERVER_NAME"))
	defer grpcClientConn.Close()
	userClient := pb.NewUserServiceClient(grpcClientConn)

	grpcClientConn = infrastructure.NewGrpcClientConn(os.Getenv("USER_SERVER_NAME"))
	defer grpcClientConn.Close()
	postClient := pb.NewPostServiceClient(grpcClientConn)

	gprcClient := client.NewClient(
		userClient,
		postClient,
	)

	router := gin.Default()
	router.GET("/users", gprcClient.GetUsers)
	router.GET("/users/:id", gprcClient.GetUser)
	router.GET("/users/:id/posts", gprcClient.GetUserPosts)
	router.GET("/posts", gprcClient.GetPosts)

	router.Run(":" + os.Getenv("PORT"))
}
