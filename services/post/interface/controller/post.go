package controller

import (
	"context"
	"post/interface/converter"
	"post/pb"
	"post/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type postController struct {
	postUsecase usecase.PostUsecase
}

func NewPostController(uu usecase.PostUsecase) *postController {
	return &postController{
		postUsecase: uu,
	}
}

func (c *postController) GetPosts(ctx context.Context, req *emptypb.Empty) (*pb.GetPostsResponse, error) {
	posts, err := c.postUsecase.GetPosts()
	if err != nil {
		return nil, err
	}
	return &pb.GetPostsResponse{
		Posts: converter.ConvertPosts(posts),
	}, nil
}

// func (c *postController) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
// 	post, err := c.postUsecase.GetPostById(req.Id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.GetPostResponse{
// 		Post: converter.ConvertPost(post),
// 	}, nil
// }
