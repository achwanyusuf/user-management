package main
import (
	"context"
	"github.com/achwanyusuf/user-management/proto"
)

func (s *server) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	token := readToken(rep, request.Token)
	if token == false{
		return &proto.DeleteResponse{
			Message: "Unauthorized Credentials",
		}, nil
	}
	
	deleted := deleteUser(rep, request.UserId)
	if deleted == false{
		return &proto.DeleteResponse{
			Message: "Can't Delete User",
		}, nil
	}
	return &proto.DeleteResponse{
		Message: "",
	}, nil
}