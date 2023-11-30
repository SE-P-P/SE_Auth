package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gopkg.in/gomail.v2"

	"SEproject/biz/application/dto/signIn"
	"SEproject/biz/infrastructure/config"
	"SEproject/biz/infrastructure/consts"
	"SEproject/biz/infrastructure/mapper"
)

type ISignInService interface {
	ResetPw(ctx context.Context, req *signIn.ResetPwReq) (*signIn.ResetPwResp, error)
	SignIn(ctx context.Context, req *signIn.SignInReq) (*signIn.SignInResp, error)
	Register(ctx context.Context, req *signIn.RegisterReq) (*signIn.RegisterResp, error)
	SendCode(ctx context.Context, req *signIn.SendCodeReq) (*signIn.SendCodeResp, error)
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

type MailboxConf struct {
	Title     string
	Body      string
	Recipient string
	Sender    string
	SPassword string
	SMTPAddr  string
	SMTPPort  int
}

func (s *SignInService) SendCode(ctx context.Context, req *signIn.SendCodeReq) (*signIn.SendCodeResp, error) {
	var mailConf MailboxConf
	mailConf.Title = "您的验证码"
	mailConf.Recipient = req.GetEmail()
	mailConf.Sender = s.Config.Email.Sender
	mailConf.SPassword = s.Config.Email.Password
	mailConf.SMTPPort = s.Config.Email.SMTPPort
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	err := s.Redis.SetexCtx(ctx, req.Email, code, 300)
	if err != nil {
		return nil, err
	}

	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, code)

	m := gomail.NewMessage()

	m.SetHeader(`From`, mailConf.Sender, "SE实践项目组")
	m.SetHeader(`To`, mailConf.Recipient)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	t := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword)
	t.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = t.DialAndSend(m)
	if err != nil {
		return nil, err
	}
	return &signIn.SendCodeResp{}, nil
}

func (s *SignInService) ResetPw(ctx context.Context, req *signIn.ResetPwReq) (*signIn.ResetPwResp, error) {
	r, err := s.Redis.GetCtx(ctx, req.Email)
	if err != nil {
		return nil, err
	} else if r == "" {
		return nil, consts.ErrCodeNotExist
	} else {
		if r == req.GetCode() {
			user, err := s.UserMapper.FindByEmail(ctx, req.Email)
			if errors.Is(err, mapper.ErrNotFound) {
				return nil, consts.ErrEmailNotExist
			} else if err != nil {
				return nil, err
			} else {
				user.Password = req.Password
				err := s.UserMapper.Update(ctx, user)
				if err != nil {
					return nil, err
				}
				return &signIn.ResetPwResp{}, nil
			}
		} else {
			return nil, consts.ErrCodeWrong
		}
	}
}

func (s *SignInService) SignIn(ctx context.Context, req *signIn.SignInReq) (*signIn.SignInResp, error) {
	resp := &signIn.SignInResp{}
	var err error
	switch req.SignInType {
	case consts.AuthTypeEmail:
		resp.UserId, err = s.signInByEmail(ctx, req.Email, req.GetCode())
	case consts.AuthTypePassword:
		resp.UserId, err = s.signInByPassword(ctx, req)
	default:
		return nil, consts.ErrInvalidType
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SignInService) Register(ctx context.Context, req *signIn.RegisterReq) (*signIn.RegisterResp, error) {
	r, err := s.Redis.GetCtx(ctx, req.Email)
	if err != nil {
		return nil, err
	} else if r == "" {
		return nil, consts.ErrCodeNotExist
	} else {
		if r == req.GetCode() {
			_, err := s.UserMapper.FindByEmail(ctx, req.Email)
			if errors.Is(err, mapper.ErrNotFound) {
				result, err := s.UserMapper.Insert(ctx, &mapper.SignIn{
					Email:    req.Email,
					Password: req.Password,
				})
				if err != nil {
					return nil, err
				}
				userId, err := result.LastInsertId()
				if err != nil {
					return nil, err
				}
				return &signIn.RegisterResp{UserId: strconv.FormatInt(userId, 10)}, nil
			} else if err != nil {
				return nil, err
			} else {
				return nil, consts.ErrEmailExist
			}
		} else {
			return nil, consts.ErrCodeWrong
		}
	}
}

func (s *SignInService) signInByPassword(ctx context.Context, req *signIn.SignInReq) (string, error) {
	user, err := s.UserMapper.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	} else if user.Password == req.GetPassword() {
		return strconv.FormatInt(user.Id, 10), nil
	}
	return "", consts.ErrPasswordWrong
}

func (s *SignInService) signInByEmail(ctx context.Context, email string, code string) (string, error) {
	r, err := s.Redis.GetCtx(ctx, email)
	if err != nil {
		return "", err
	} else if r == "" {
		return "", consts.ErrCodeNotExist
	} else {
		if r == code {
			user, err := s.UserMapper.FindByEmail(ctx, email)
			if err != nil {
				return "", err
			}
			return strconv.FormatInt(user.Id, 10), nil
		}
		return "", consts.ErrCodeWrong
	}
}
