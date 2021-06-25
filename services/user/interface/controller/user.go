package controller

import (
	"context"
	"user/interface/converter"
	"user/pb"
	"user/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) *userController {
	return &userController{
		userUsecase: uu,
	}
}

func (c *userController) GetUsers(ctx context.Context, req *emptypb.Empty) (*pb.GetUsersResponse, error) {
	users, err := c.userUsecase.GetUsers()
	if err != nil {
		return nil, err
	}
	return converter.ConvertUsers(users)

}

// func (c *userController) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.User, error) {
// 	user, err := c.userUsecase.GetUserById(req.Id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return converter.ConvertUser(user)
// }

// func (c *userController) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.User, error) {
// 	reqUser := &domain.User{
// 		LastName:  req.LastName,
// 		FirstName: req.FirstName,
// 	}
// 	user, err := c.userUsecase.CreateUser(reqUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return converter.ConvertUser(user)
// }

// func (c *userController) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.User, error) {
// 	reqUser := &domain.User{
// 		Id:        req.Id,
// 		LastName:  req.LastName,
// 		FirstName: req.FirstName,
// 	}
// 	user, err := c.userUsecase.UpdateUser(reqUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return converter.ConvertUser(user)
// }

// func (c *userController) DestroyUser(ctx context.Context, req *pb.DestroyUserReq) (*emptypb.Empty, error) {
// 	if _, err := c.userUsecase.GetUserById(req.Id); err != nil {
// 		return nil, err
// 	}
// 	if err := c.userUsecase.DeleteUserById(req.Id); err != nil {
// 		return nil, err
// 	}
// 	return &emptypb.Empty{}, nil
// }
