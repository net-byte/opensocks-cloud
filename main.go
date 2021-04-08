package main

import (
	"os"

	"github.com/net-byte/opensocks/config"
	"github.com/net-byte/opensocks/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	username := os.Getenv("_USERNAME")
	if username == "" {
		username = "admin"
	}
	password := os.Getenv("_PASSWORD")
	if password == "" {
		password = "pass@123456"
	}
	config := config.Config{}
	config.ServerMode = true
	config.ServerAddr = ":" + port
	config.Username = username
	config.Password = password
	config.Init()
	server.Start(config)
}
