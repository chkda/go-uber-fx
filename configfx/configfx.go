package configfx

import (
	"io/ioutil"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

func ProvideConfig() *ApplicationConfig {
	conf := ApplicationConfig{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}

var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
