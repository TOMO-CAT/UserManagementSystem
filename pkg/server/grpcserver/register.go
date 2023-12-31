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

	// 参数校验
	if !checkUserName(req.GetUserName()) {
		wrapResp(res, false, common.ErrorCodeRegisterInvalidUserName, common.ErrorMsgRegisterInvalidUserName)
		return
	}
	if !checkPassword(req.GetUserPassword()) {
		wrapResp(res, false, common.ErrorCodeRegisterInvalidPassword, common.ErrorMsgRegisterInvalidPassword)
		return
	}

	// ·

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
	res.IsSuccess = isSuccess
	res.ErrorCode = int32(errno)
	res.ErrorMsg = errMsg
}
