package main

import (
	"fmt"
	"os"

	"github.com/net-byte/opensocks/config"
	"github.com/net-byte/opensocks/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	key := os.Getenv("KEY")
	if key == "" {
		key = "6w9z$C&F)J@NcRfUjXn2r4u7x!A%D*G-"
	}
	config := config.Config{}
	config.ServerMode = true
	config.ServerAddr = fmt.Sprintf(":%v", port)
	config.Key = key
	server.Start(config)
}
