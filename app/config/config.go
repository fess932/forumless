package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Forums []Forum `toml:"forum"`
}

type Forum struct {
	Name string
	Host string
	Port string
}

func New() Config {
	var conf Config

	if _, err := toml.DecodeFile("app/config/config.toml", &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
