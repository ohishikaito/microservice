package client

import (
	"context"
	"net/http"

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
	// posts, err := c.postClient.GetPosts(context.Background(), &emptypb.Empty{})
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"errorMessage": string(err.Error()),
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, posts)
}
