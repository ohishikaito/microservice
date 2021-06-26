package converter

import (
	"user/domain"
	"user/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUsers(users []*domain.User) (*pb.GetUsersResponse, error) {
	var pbUsers []*pb.User
	for _, user := range users {
		pbUser := ConvertUser(user)
		pbUsers = append(pbUsers, pbUser)
	}
	return &pb.GetUsersResponse{Users: pbUsers}, nil
}

func ConvertUser(user *domain.User) *pb.User {
	pbUser := &pb.User{
		Id:              user.Id,
		LastName:        user.LastName,
		FirstName:       user.FirstName,
		Email:           user.Email,
		TelephoneNumber: user.TelephoneNumber,
		Gender:          uint64(user.Gender),
		// clientがRailsだとRFC3339に変えるけど、clientがGoだとepoch timeで返してる
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	return pbUser
}
