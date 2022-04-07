package rounter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nft/database"
)

func Index2(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	if e == nil {
		isLogin = true
	}
	if isLogin {
		c.HTML(http.StatusOK, "index-2.html", gin.H{
			"isLogin":  isLogin,
			"username": cookie.Value,
			"img":      database.QueryUser(cookie.Value).Picture,
		})
	} else {
		c.HTML(http.StatusOK, "index-2.html", gin.H{
			"isLogin": isLogin,
		})
	}
}
func NotFind(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{})
}
func Collection(c *gin.Context) {
	c.HTML(http.StatusOK, "collection.html", gin.H{})
}
func Drag(c *gin.Context) {
	c.HTML(http.StatusOK, "drag.html", gin.H{})
}
func Homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", gin.H{})
}
func Index1(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
func ItemDetails(c *gin.Context) {
	c.HTML(http.StatusOK, "item-details.html", gin.H{})
}
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func Ranking(c *gin.Context) {
	c.HTML(http.StatusOK, "ranking.html", gin.H{})
}
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func Swap(c *gin.Context) {
	c.HTML(http.StatusOK, "swap.html", gin.H{})
}
func Wallet(c *gin.Context) {
	c.HTML(http.StatusOK, "wallet.html", gin.H{})
}
