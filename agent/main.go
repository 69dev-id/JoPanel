package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"jopanel/agent/auth"
	"jopanel/agent/executor"
)

func main() {
	// Security: Agent should listen on localhost only or a unix socket
	// For dev, localhost:8081
	r := gin.Default()

	// Auth Middleware
	r.Use(auth.RequireSecret())

	r.POST("/execute", executor.HandleCommand)

	port := os.Getenv("AGENT_PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("JoPanel Agent listening on 127.0.0.1:%s", port)
	r.Run("127.0.0.1:" + port)
}
