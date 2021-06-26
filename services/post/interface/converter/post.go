package converter

import (
	"post/domain"
	"post/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertPosts(posts []*domain.Post) []*pb.Post {
	var pbPosts []*pb.Post
	for _, post := range posts {
		pbPost := ConvertPost(post)
		pbPosts = append(pbPosts, pbPost)
	}
	return pbPosts
}

func ConvertPost(post *domain.Post) *pb.Post {
	pbPost := &pb.Post{
		Id:        post.Id,
		Text:      post.Text,
		UserId:    post.UserId,
		CreatedAt: timestamppb.New(post.CreatedAt),
		UpdatedAt: timestamppb.New(post.UpdatedAt),
	}
	return pbPost
}
