package config

import (
	"os"

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
	c := new(Config)
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "etc/config.yaml"
	}
	err := conf.Load(path, c)
	if err != nil {
		return nil, err
	}
	err = c.SetUp()
	if err != nil {
		return nil, err
	}
	return c, nil
}
