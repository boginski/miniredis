package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/api/set", handler.SetKeyValue)
	router.GET("/api/get/key", handler.GetKey)
	router.GET("/api/get/pattern", handler.GetPatternKey)
	router.DELETE("/api/delete", handler.DeleteKey)

	router.Run(":8000")
}
