package main

import (
	"context"
	"github.com/achwanyusuf/user-management/proto"
	"github.com/achwanyusuf/user-management/model"
)

func (s *server) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	token := readToken(rep, request.Token)
	if token == false{
		return &proto.UpdateResponse{
			Message: "Unauthorized Credentials",
		}, nil
	}
	data := model.UserData{UserId: request.UserId, Email: request.Email, Address: request.Address}
	updated := updateUser2(rep, &data)
	if updated == false{
		return &proto.UpdateResponse{
			Message: "Error Updating",
		}, nil
	}
	return &proto.UpdateResponse{
		UserId: request.UserId,
		Email: request.Email,
		Address: request.Address,
	}, nil
}