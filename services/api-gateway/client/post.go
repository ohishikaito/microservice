package client

import (
	"api_gateway/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *client) GetPosts(ctx *gin.Context) {
	posts, err := c.postClient.GetPosts(context.Background(), &emptypb.Empty{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (c *client) GetUserPosts(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	req := &pb.GetUserRequest{
		Id: id,
	}
	user, err := c.userClient.GetUser(context.Background(), req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}

	req2 := &pb.GetPostsByUserRequest{
		UserId: int64(user.User.Id),
	}
	posts, err := c.postClient.GetPostsByUser(context.Background(), req2)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (c *client) GetPost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	req := &pb.GetPostReq{
		Id: id,
	}
	post, err := c.postClient.GetPost(context.Background(), req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, post)
}
