package consts

import "google.golang.org/grpc/status"

const (
	AuthTypeEmail    = "email"
	AuthTypePassword = "password"
)

var (
	ErrInvalidType   = status.Error(10001, "invalid signIn type")
	ErrEmailFailed   = status.Error(10002, "Send Email Fail")
	ErrCodeNotExist  = status.Error(10003, "Code Not Exist")
	ErrCodeWrong     = status.Error(10004, "Code Wrong")
	ErrPasswordWrong = status.Error(10005, "Password Wrong")
	ErrEmailExist    = status.Error(10006, "Email already exist")
	ErrEmailNotExist = status.Error(10006, "Email not exist")
)
