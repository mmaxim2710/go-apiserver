package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		BindAddr string `yaml:"bind_addr"`
	} `yaml:"server"`

	SessionKey string `yaml:"session_key"`
	LogLevel   string `yaml:"log_level"`

	DB struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		User   string `yaml:"user"`
		DBName string `yaml:"db_name"`
	} `yaml:"db"`
}

func NewConfig(cfgPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
