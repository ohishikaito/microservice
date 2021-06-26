package client

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *client) GetUsers(ctx *gin.Context) {
	users, err := c.userClient.GetUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *client) GetUser(ctx *gin.Context) {
	// users, err := c.userClient.GetUsers(context.Background(), &emptypb.Empty{})
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"errorMessage": string(err.Error()),
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, users)
}
