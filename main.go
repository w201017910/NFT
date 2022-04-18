package main

import (
	"github.com/gin-gonic/gin"
	"nft/database"
	"nft/rounter"
)

func main() {
	database.Init()
	Start("localhost:8080")
}
func Start(addr string) (err error) {
	r := gin.Default()
	r.Static("/frontend", "./frontend")
	r.LoadHTMLGlob("./template/*")
	r.GET("/index-2.html", rounter.Index2)
	r.GET("/404.html", rounter.NotFind)
	r.GET("/collection.html", rounter.Collection)
	r.GET("/drag.html", rounter.Drag)
	r.GET("/homepage.html", rounter.Homepage)
	r.GET("/item-details.html", rounter.ItemDetails)
	r.GET("/login.html", rounter.LoginPage)
	r.GET("/ranking.html", rounter.Ranking)
	r.GET("/register.html", rounter.RegisterPage)
	r.GET("/swap.html", rounter.Swap)
	r.GET("/wallet.html", rounter.Wallet)
	r.GET("/", rounter.Index1)
	r.POST("/register", rounter.Register)
	r.POST("/login", rounter.Login)
	r.POST("/logout", rounter.LoginOut)
	err = r.Run(addr)
	return err
}
