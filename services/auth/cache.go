package auth

import (
	"github.com/go-redis/redis"

	"github.com/caarlos0/env"

	"fmt"
)

type redisConf struct {
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	//Host     string `env:"REDIS_HOST" envDefault:"redis"`
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD"`
}

var redisClient *redis.Client

func init() {
	var conf redisConf
	env.Parse(&conf)
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password,
	})

	_, err := redisClient.Ping().Result()
	if nil != err {
		panic(fmt.Errorf("Failed to ping redis server: %v", err))
	}
}

const (
	keyPrefix = "wizardcloud.cn.weixin.mp."
)