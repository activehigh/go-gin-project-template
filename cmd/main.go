package main

import (
	"github.com/activehigh/go-gin-project-template/pkg/v1/configs"
	v1 "github.com/activehigh/go-gin-project-template/pkg/v1/server"
)

func main() {
	config := configs.New()
	config.LoadFromEnv()
	server := v1.NewServer(config)
	server.Start()
}
