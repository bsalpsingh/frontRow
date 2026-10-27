package router

import (
	"github.com/bsalpsingh/frontRow/internal/router/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Create a Gin router with default middleware (logger and recovery)
	e := gin.Default()
	v1 := e.Group("/api/v1")
	handler.EventRouter((v1))

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	e.Run()
}
