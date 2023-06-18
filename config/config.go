package config

import "github.com/JWhist/jwconfig"

type Config struct {
	Port int `yaml:"port"`
}

func NewConfig(path string) Config {
	c := jwconfig.LoadFile[Config](path)
	return c
}
