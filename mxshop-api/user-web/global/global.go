package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop-api/user-web/config"
	__proto "mxshop-api/user-web/proto"
)

var (
	SystemConfig  *config.SystemConfig = &config.SystemConfig{}
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	NacosConfig   *config.NacosConfig  = &config.NacosConfig{}
	Trans         ut.Translator
	UserSrvClient __proto.UserClient
)

const (
	EnvMXSHOP = "MXSHOP_DEBUG"
)
