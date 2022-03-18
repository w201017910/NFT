package main

import (
	"github.com/gin-gonic/gin"
	"ipfs-app/routers"
)

func main() {
	router := gin.Default()
	router.POST("/upLoad", routers.ReadFile)
	router.Run("127.0.0.1:12345")
}
