package main
import (
	"context"
	"github.com/achwanyusuf/user-management/proto"
	uuid "github.com/google/uuid"
	"strings"
	"github.com/achwanyusuf/user-management/model"
	"github.com/achwanyusuf/user-management/utils/cryptography"
)

func (s *server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	token := readToken(rep, request.Token)
	if token == false{
		return &proto.CreateResponse{
			Message: "Unauthorized Credentials",
		}, nil
	}
	userId:=uuid.New().String()
	shortUserId:=strings.Replace(userId, "-", "", -1)
	password := cryptography.HashAndSalt([]byte(request.Password))
	data := model.UserData{UserId: shortUserId, Email: request.Email, Password: password, Address: request.Address}
	created := createUser(rep, &data)
	if created == false{
		return &proto.CreateResponse{
			Message: "User already exist",
		}, nil
	}
	return &proto.CreateResponse{
		UserId: shortUserId,
		Email: request.Email,
		Address: request.Address,
	}, nil
}