package grpcserver

import (
	"context"
	"regexp"

	proto "github.com/TOMO-CAT/UserManagementSystem/proto/service"
)

func (s *server) Register(ctx context.Context, req *proto.RegisterRequest) (res *proto.RegisterResponse, err error) {
	return
}

func checkUserName(userName string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", userName)
	return ok
}

func checkPasswork(password string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password)
	return ok
}
