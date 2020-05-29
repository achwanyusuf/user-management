package main

import (
	"context"
	"github.com/achwanyusuf/user-management/proto"
)

func (s *server) ReadOne(ctx context.Context, request *proto.ReadOneRequest) (*proto.ReadOneResponse, error) {
	token := readToken(rep, request.Token)
	if token == false{
		return &proto.ReadOneResponse{
			Message: "Unauthorized Credentials",
		}, nil
	}
	userData := userDataById(rep, request.UserId)
	if userData == nil {
		return &proto.ReadOneResponse{
			Message: "Can't find any data",
		}, nil
	}
	return &proto.ReadOneResponse{
		UserId: userData.UserId,
		Email: userData.Email,
		Address: userData.Address,
		LatestLogin: userData.LatestLogin,
	}, nil
}


func (s *server) ReadAll(ctx context.Context, request *proto.ReadAllRequest) (*proto.ReadAllResponse, error) {
	type Dictionary map[string]interface{}
	token := readToken(rep, request.Token)
	if token == false{
		return &proto.ReadAllResponse{
			Message: "Unauthorized Credentials",
		}, nil
	}
	UserDatas := readAll(rep)
	return UserDatas, nil
}