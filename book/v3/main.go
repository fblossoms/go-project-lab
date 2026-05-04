package main

import (
	"fmt"
	"go18/book/v3/config"
	"go18/book/v3/exception"
	"go18/book/v3/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.toml"
	}
	config.LoadConfigFromYaml(path)

	server := gin.New()
	server.Use(gin.Logger(), exception.Recovery())

	handlers.Book.Registry(server)

	ac := config.C().Application

	err := server.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
