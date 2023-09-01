package grpcserver

import (
	"context"
	"regexp"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/server/common"
	proto "github.com/TOMO-CAT/UserManagementSystem/proto/service"
)

func (s *server) Register(ctx context.Context, req *proto.RegisterRequest) (res *proto.RegisterResponse, err error) {
	res = &proto.RegisterResponse{
		IsSuccess: true,
		ErrorCode: common.ErrorCodeOk,
		ErrorMsg:  common.ErrorMsgOk,
	}

	return
}

func checkUserName(userName string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", userName)
	return ok
}

func checkPassword(password string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password)
	return ok
}

func wrapResp(res *proto.RegisterResponse, isSuccess bool, errno int, errMsg string) {

}
