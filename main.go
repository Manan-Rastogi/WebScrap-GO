package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/facts", Facts)

	server.Run(":8080")

}
