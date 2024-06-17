package config

import (
	"os"
	"reflect"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int `yaml:"port"`
}

func NewConfig(path string) Config {
	c := loadFile[Config](path)
	return c
}

func loadFile[T any](path string) T {
	var t T
	err := load(path, &t)
	if err != nil {
		panic(err)
	}
	return t
}

func load(configFile string, obj interface{}) error {
	content, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	if reflect.TypeOf(obj).Kind() == reflect.Pointer {
		typ := reflect.ValueOf(obj).Elem().Type()
		if typ.Kind() == reflect.Struct {
			if err := defaults.Set(obj); err != nil {
				return err
			}
		}
	}

	return yaml.Unmarshal(content, obj)
}
