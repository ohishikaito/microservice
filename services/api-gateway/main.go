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
	router := gin.Default()
	conn := newGrpcClientConn()

	userClient := pb.NewUserServiceClient(conn)
	postClient := pb.NewPostServiceClient(conn)
	client := NewClient(
		userClient,
		postClient,
	)

	router.GET("/users", client.GetUsers)
	router.GET("/posts", client.GetPosts)

	router.Run(":" + os.Getenv("PORT"))
}

func newGrpcClientConn() *grpc.ClientConn {
	targetUrl := os.Getenv("USER_SERVER_NAME") + ":" + os.Getenv("GRPC_SERVER_PORT")
	conn, err := grpc.Dial(targetUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
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
