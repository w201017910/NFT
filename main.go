package main

import "github.com/gin-gonic/gin"

func main() {
	Start("localhost:8080")
}
func Start(addr string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()

	// 静态文件服务

	r.Static("/web", "./frontend")

	// 启动web服务
	err = r.Run(addr)
	return err
}
