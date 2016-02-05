package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Settings *SiteSettings

type SiteSettings struct {
	Server   *ServerSettings
	Database *DatabaseSettings
	Redis    *RedisSettings
}

type ServerSettings struct {
	Address string `yaml:"address"`
}

type DatabaseSettings struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}

type RedisSettings struct {
	Address  string `yaml:"address"`
	Database int64  `yaml:"database"`
}

func InitSettings(file string) (err error) {
	Settings = &SiteSettings{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(data, Settings); err != nil {
		return
	}
	return
}
