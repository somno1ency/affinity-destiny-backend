package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Datasource map[string]string
}
