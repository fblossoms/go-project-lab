package main

import (
	"fmt"
	"go18/book/v3/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	handlers.Book.Registry(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
