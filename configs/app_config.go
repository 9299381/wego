package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type AppConfig struct {
	Config
	Env  string `json:"env"`
	Name string `json:"app_name"`
}

func (it *AppConfig) Load() contracts.Iconfig {

	config := &AppConfig{
		Env:  wego.Env("APP_Env","prod"),
		Name: wego.Env("APP_NAME","app"),
	}
	return config
}

func (it *AppConfig)Get(key string)  string {
	return it.GetKey(it,key)
}