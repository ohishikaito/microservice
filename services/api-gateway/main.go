package main

import (
	"api_gateway/infrastructure"
	"api_gateway/pb"
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	router := gin.Default()

	grpcClientConn := infrastructure.NewGrpcClientConn(os.Getenv("USER_SERVER_NAME"))
	defer grpcClientConn.Close()
	userClient := pb.NewUserServiceClient(grpcClientConn)

	grpcClientConn = infrastructure.NewGrpcClientConn(os.Getenv("USER_SERVER_NAME"))
	defer grpcClientConn.Close()
	postClient := pb.NewPostServiceClient(grpcClientConn)

	client := NewClient(
		userClient,
		postClient,
	)

	router.GET("/users", client.GetUsers)
	router.GET("/posts", client.GetPosts)

	router.Run(":" + os.Getenv("PORT"))
}

type client struct {
	user pb.UserServiceClient
	post pb.PostServiceClient
}

func NewClient(
	userClient pb.UserServiceClient,
	postClient pb.PostServiceClient,
) *client {
	return &client{
		user: userClient,
		post: postClient,
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

func (c *client) GetPosts(ctx *gin.Context) {
	posts, err := c.post.GetPosts(context.Background(), &emptypb.Empty{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}
