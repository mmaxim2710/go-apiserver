package main

import (
	"flag"
	"github.com/mmaxim2710/firstWebApp/internal/app/apiserver"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"log"
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
		log.Fatal("err in main(): ", err)
	}

	s := apiserver.New(newConfig)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
