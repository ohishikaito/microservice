package controller

import (
	"context"
	"user/domain"
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
	return &pb.GetUsersResponse{
		Users: converter.ConvertUsers(users),
	}, nil
}

func (c *userController) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := c.userUsecase.GetUserById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		User: converter.ConvertUser(user),
	}, nil
}

func (c *userController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	reqUser := &domain.User{
		LastName:  req.LastName,
		FirstName: req.FirstName,
	}
	user, err := c.userUsecase.CreateUser(reqUser)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: converter.ConvertUser(user),
	}, nil
}

func (c *userController) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	reqUser := &domain.User{
		Id:        req.Id,
		LastName:  req.LastName,
		FirstName: req.FirstName,
	}
	user, err := c.userUsecase.UpdateUser(reqUser)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		User: converter.ConvertUser(user),
	}, nil
}

func (c *userController) DestroyUser(ctx context.Context, req *pb.DestroyUserRequest) (*emptypb.Empty, error) {
	if _, err := c.userUsecase.GetUserById(req.Id); err != nil {
		return nil, err
	}
	if err := c.userUsecase.DeleteUserById(req.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
