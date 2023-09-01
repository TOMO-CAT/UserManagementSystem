package common

const (
	ErrorCodeOk = 0

	// Register: 10001 ~ 19999
	ErrorCodeRegisterInvalidUserName = 10001
	ErrorCodeRegisterInvalidPassword = 10002
	ErrorCodeRegisterInternalError   = 10003
)

const (
	ErrorMsgOk = "ok"

	// Register interface
	ErrorMsgRegisterInvalidUserName = "invalid user name"
	ErrorMsgRegisterInvalidPassword = "invalid password"

	// Login interface
)
