package config

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	IP   string `toml:"server_ip"`
	Port string `toml:"server_port"`
}

var Get config

func ReadConfig() {
	if _, err := toml.DecodeFile("config/config.toml", &Get); err != nil {
		panic(err)
	}
}
