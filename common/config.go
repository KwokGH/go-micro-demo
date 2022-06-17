package common

import (
	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	source := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)

	cfg, err := config.NewConfig()
	if err != nil {
		return cfg, err
	}

	return cfg, cfg.Load(source)
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func GetMysqlConfig(config config.Config, path ...string) (MysqlConfig, error) {
	cfg := MysqlConfig{}

	config.Get(path...).Scan(&cfg)

	return cfg, nil
}
