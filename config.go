package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Basedir  string `yaml:"basedir"`
}

type Config struct {
	Appconf AppConfig `yaml:"app"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("config file not found, using default config")
			return &Config{
				Appconf: AppConfig{
					Password: "",
					User:     "",
					Basedir:  "",
				},
			}, nil
		}
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
