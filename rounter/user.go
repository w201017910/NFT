package rounter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nft/database"
	"nft/util"
	"path"
	"strconv"
	"time"
)

func Register(ctx *gin.Context) {
	//获取普通文本

	name := ctx.PostForm("1")
	password := ctx.PostForm("2")
	email := ctx.PostForm("3")
	// 获取文件(注意这个地方的file要和html模板中的name一致)
	file, err := ctx.FormFile("0")
	newName := "static/picture/img-15.png"
	if err != nil {
		fmt.Println("获取数据失败")
	} else {
		fmt.Println("接收的数据", file.Filename)
		//获取文件名称
		fmt.Println(file.Filename)
		//文件大小
		fmt.Println(file.Size)
		//获取文件的后缀名
		extstring := path.Ext(file.Filename)
		fmt.Println(extstring)
		//根据当前时间鹾生成一个新的文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		//新的文件名
		fileName := fileNameStr + extstring
		//保存上传文件
		newName = "./frontend/static/updata/" + fileName
		ctx.SaveUploadedFile(file, newName)
	}
	path, address, mneonic := util.NewAccount("./keystore", password)

	database.InsertUser(name, password, address, email, newName, path, mneonic)
	ctx.SetCookie("name", name, 1000, "/", "localhost", false, true)
	ctx.Redirect(http.StatusMovedPermanently, "/")

}
func Login(c *gin.Context) {
	name := c.PostForm("1")
	password := c.PostForm("2")
	person := database.SignIn(name, password)
	if person != nil {
		c.SetCookie("name", name, 1000, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"judge": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"judge": false,
		})
	}

}
func LoginOut(c *gin.Context) {

	c.SetCookie("name", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}
