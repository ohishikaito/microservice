package client

import (
	"api_gateway/pb"
	"context"
	"net/http"
	"strconv"

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
	ctx.JSON(http.StatusOK, user)
}
