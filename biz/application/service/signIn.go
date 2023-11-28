package service

import (
	"context"

	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"SEproject/biz/application/dto/signIn"
	"SEproject/biz/infrastructure/config"
	"SEproject/biz/infrastructure/mapper"
)

type ISignInService interface {
	ResetPw(ctx context.Context, req *signIn.ResetPwReq) (*signIn.ResetPwResp, error)
	SignIn(ctx context.Context, req *signIn.SignInReq) (*signIn.SignInResp, error)
	Register(ctx context.Context, req *signIn.RegisterReq) (*signIn.RegisterResp, error)
}

type SignInService struct {
	Config     *config.Config
	UserMapper mapper.SignInModel
	Redis      *redis.Redis
}

var SignInServiceSet = wire.NewSet(
	wire.Struct(new(SignInService), "*"),
	wire.Bind(new(ISignInService), new(*SignInService)),
)

func (s *SignInService) ResetPw(ctx context.Context, req *signIn.ResetPwReq) (*signIn.ResetPwResp, error) {

	//TODO implement me
	panic("implement me")
}

func (s *SignInService) SignIn(ctx context.Context, req *signIn.SignInReq) (*signIn.SignInResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SignInService) Register(ctx context.Context, req *signIn.RegisterReq) (*signIn.RegisterResp, error) {
	//TODO implement me
	panic("implement me")
}
