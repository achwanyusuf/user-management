package main
import (
	"context"
	"github.com/achwanyusuf/user-management/proto"
	"github.com/achwanyusuf/user-management/utils/cryptography"
	"time"
)

func (s *server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	userData := userData(rep, request.Email)
	if userData == nil {
		return &proto.LoginResponse{
			Message: "Can't find any data",
		}, nil
	}

	// Validating Password
	isValid := cryptography.ValidatingHash(userData.Password,[]byte(request.Password))
	if isValid == false{
		return &proto.LoginResponse{
			Message: "Password is incorrect",
		}, nil
	}
	token := cryptography.GenerateToken(request.Email)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userData.Token = token
	userData.LatestLogin = currentTime
	//updating data
	updateUser(rep, userData)
	return &proto.LoginResponse{
		UserId: userData.UserId,
		Email: userData.Email,
		Address: userData.Address,
		Token: token,
		LatestLogin: currentTime,
		Message: "",
	}, nil
}
