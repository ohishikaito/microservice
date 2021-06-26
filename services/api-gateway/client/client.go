package client

import (
	"api_gateway/pb"
)

type client struct {
	userClient pb.UserServiceClient
	postClient pb.PostServiceClient
}

func NewClient(
	UserServiceClient pb.UserServiceClient,
	PostServiceClient pb.PostServiceClient,
) *client {
	return &client{
		userClient: UserServiceClient,
		postClient: PostServiceClient,
	}
}
