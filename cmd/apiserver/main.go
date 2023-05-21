package main

import (
	"flag"
	"github.com/mmaxim2710/firstWebApp/internal/app/apiserver"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"github.com/mmaxim2710/firstWebApp/internal/app/logger"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func main() {
	flag.Parse()

	newConfig, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	if err := logger.ConfigureLogger(newConfig.LogLevel); err != nil {
		panic(err)
	}

	if err := apiserver.Start(newConfig); err != nil {
		panic(err)
	}
}
