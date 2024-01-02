package provider

import (
	"github.com/google/wire"

	"SEproject/biz/infrastructure/redis"

	"SEproject/biz/application/service"
	"SEproject/biz/infrastructure/config"
	"SEproject/biz/infrastructure/mapper"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

// Provider 提供controller依赖的对象
type Provider struct {
	Config        *config.Config
	SignInService service.ISignInService
}

func Get() *Provider {
	return provider
}

var ApplicationSet = wire.NewSet(
	service.SignInServiceSet,
)
var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	mapper.NewSignInModel,
	mapper.NewSaveModel,
	redis.NewRedis,
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	InfrastructureSet,
)
