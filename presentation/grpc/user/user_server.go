package demo_microservice

import (
	"context"
	"github.com/S-L-T/demo-microservice/domain/entity"
	"github.com/S-L-T/demo-microservice/domain/use_case"
	"github.com/S-L-T/demo-microservice/helper"
	"strconv"
	"time"
)

type UserServer struct {
	userUseCase use_case.User
	UnimplementedUserServiceServer
}

func NewGRPCUserServer(u use_case.User) UserServer {
	return UserServer{userUseCase: u}
}

func (u UserServer) AddUser(ctx context.Context, request *AddUserRequest) (*AddUserResponse, error) {
	user := entity.User{
		FirstName: request.User.FirstName,
		LastName:  request.User.LastName,
		Nickname:  request.User.Nickname,
		Password:  request.User.Password,
		Email:     request.User.Email,
		Country:   request.User.Country,
	}

	id, err := u.userUseCase.AddUser(user)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return &AddUserResponse{}, err
	}
	return &AddUserResponse{
		Id: id,
	}, nil
}

func (u UserServer) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	user := entity.User{
		ID:        request.Id,
		FirstName: request.User.FirstName,
		LastName:  request.User.LastName,
		Nickname:  request.User.Nickname,
		Password:  request.User.Password,
		Email:     request.User.Email,
		Country:   request.User.Country,
	}

	err := u.userUseCase.UpdateUser(user)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return &UpdateUserResponse{}, err
	}

	return &UpdateUserResponse{}, nil
}

func (u UserServer) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error) {
	err := u.userUseCase.DeleteUser(request.Id)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return &DeleteUserResponse{}, err
	}

	return &DeleteUserResponse{}, nil
}

func (u UserServer) GetPaginatedUsers(ctx context.Context, request *GetPaginatedUsersRequest) (*GetPaginatedUsersResponse, error) {
	filter := entity.Filter{
		FirstName: request.Filter.FirstName,
		LastName:  request.Filter.LastName,
		Nickname:  request.Filter.Nickname,
		Email:     request.Filter.Email,
		Country:   request.Filter.Country,
	}

	pageNum, err := strconv.ParseUint(request.PageNum, 10, 64)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return &GetPaginatedUsersResponse{}, err
	}

	pageSize, err := strconv.ParseUint(request.PageSize, 10, 64)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return &GetPaginatedUsersResponse{}, err
	}

	returnedUsers, err := u.userUseCase.GetPaginatedUsers(filter, pageNum, pageSize)

	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return &GetPaginatedUsersResponse{}, err
	}

	var users []*User
	for _, u := range returnedUsers {
		users = append(users, &User{
			Id:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.Nickname,
			Password:  u.Password,
			Email:     u.Email,
			Country:   u.Country,
			CreatedAt: u.CreatedAt.Format(time.RFC3339),
			UpdatedAt: u.CreatedAt.Format(time.RFC3339),
		})
	}

	return &GetPaginatedUsersResponse{
		Users: users,
	}, nil
}

func (u UserServer) mustEmbedUnimplementedUserServiceServer() {
	panic("implement me")
}
