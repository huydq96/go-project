package main

import (
	"log"

	"go-project/config"
	"go-project/internal/server"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)
	log.Fatal(s.Start())
}
