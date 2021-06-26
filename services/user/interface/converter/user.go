package converter

import (
	"fmt"
	"user/domain"
	"user/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUsers(users []*domain.User) []*pb.User {
	var pbUsers []*pb.User
	for _, user := range users {
		pbUser := ConvertUser(user)
		pbUsers = append(pbUsers, pbUser)
	}
	return pbUsers
}

func ConvertUser(user *domain.User) *pb.User {
	fmt.Println(user.Gender)
	pbUser := &pb.User{
		Id:              user.Id,
		LastName:        user.LastName,
		FirstName:       user.FirstName,
		Email:           user.Email,
		TelephoneNumber: user.TelephoneNumber,
		Gender:          user.Gender,
		// FIXME: user.Genderが0だったら、pbでjson:omitemptyしてるせいでjsonにGenderが含まれない
		// Gender: 0,
		// clientがRailsだとRFC3339に変えるけど、clientがGoだとepoch timeで返してる
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	return pbUser
}
