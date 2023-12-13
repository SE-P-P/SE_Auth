package config

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type MySql struct {
	URL string
}

type Email struct {
	Sender   string
	Password string
	SMTPAddr string
	SMTPPort int
}

type Config struct {
	service.ServiceConf
	MySql MySql
	Redis *redis.RedisConf
	Email Email
}

func NewConfig() (*Config, error) {
	sc := []constant.ServerConfig{{
		IpAddr: "124.223.119.145",
		Port:   8848,
	}}

	cc := constant.ClientConfig{
		NamespaceId:         "",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		LogLevel:            "error",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		return nil, err
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "config.yaml",
		Group:  "DEFAULT_GROUP",
	})
	config, err := SetConfig(content)
	if err != nil {
		return nil, err
	}
	err = config.SetUp()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func SetConfig(content string) (config *Config, err error) {
	c := new(Config)
	err = conf.LoadFromYamlBytes([]byte(content), &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
