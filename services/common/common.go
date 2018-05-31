package common

import (
	"fmt"
	"github.com/caarlos0/env"
)

type componentAppConf struct {
	AppId     string `env:"COMP_APP_ID,required"`
	AppSecret string `env:"COMP_APP_SECRET,required"`
}

var compAppConf componentAppConf

func init() {
	if err := env.Parse(&compAppConf); nil != err {
		panic(fmt.Errorf("Parse component app config error: %v", err))
	}
}

func GetCompAppId() string {
	return compAppConf.AppId
}

func GetCompAppSecret() string {
	return compAppConf.AppSecret
}
