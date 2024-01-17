package global

import (
	"gorm.io/gorm"
	"mvshop_srvs/user_srv/config"
)

var (
	DB *gorm.DB
	// ServerConfig *config.ServerConfig
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)

func init() {

}
