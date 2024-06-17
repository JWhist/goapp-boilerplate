package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type config struct {
	String    string `yaml:"string"`
	Numbers   []int  `yaml:"numbers"`
	Default   string `yaml:"default" default:"default"`
	NoDefault string `yaml:"no_default"`
}

func TestMustLoadFile(t *testing.T) {
	c := loadFile[config]("test.yaml")
	assert.Equal(t, "value", c.String)
	assert.Equal(t, []int{1, 2, 3}, c.Numbers)
	assert.Equal(t, "default", c.Default)
	assert.Equal(t, "", c.NoDefault)
}
