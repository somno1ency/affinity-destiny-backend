package config

import (
	zeroRds "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource map[string]string
	RedisConf  zeroRds.RedisConf
}
