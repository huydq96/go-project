package main

import (
	"log"

	"go-project/config"
	"go-project/internal/server"
)

func main() {
	cfgPath, err := config.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.InitConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)
	log.Fatal(s.Start())
}
