package rounter

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
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

	name := ctx.PostForm("userName")
	password := ctx.PostForm("passwd")
	email := ctx.PostForm("mail")
	// 获取文件(注意这个地方的file要和html模板中的name一致)
	file, err := ctx.FormFile("picFile")
	privateKey := ctx.PostForm("privateKey")
	auth := ctx.PostForm("auth")
	keyStoreFile, keyFileErr := ctx.FormFile("keyStoreFile")
	newName := "static/picture/img-15.png"
	if err != nil {
		fmt.Println("获取数据失败")
	} else {
		fmt.Println("接收的数据", file.Filename)
		//获取文件的后缀名
		extstring := path.Ext(file.Filename)
		//根据当前时间鹾生成一个新的文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		//新的文件名
		fileName := fileNameStr + extstring
		//保存上传文件
		newName = "./frontend/static/updata/" + fileName
		ctx.SaveUploadedFile(file, newName)
	}
	if privateKey != "" {
		keyFile, mnemonic, err := util.ImportAccountByPrivateKey(privateKey, password)
		if err != "" {
			ctx.JSON(200, gin.H{
				"err": err,
			})
		}
		privateKey_, _ := crypto.HexToECDSA(privateKey)
		publicKey := crypto.PubkeyToAddress(privateKey_.PublicKey)
		database.InsertUser(name, password, publicKey.String(), email, newName, keyFile, mnemonic)
	}
	if keyFileErr != nil {
		fmt.Println("获取数据失败")
	} else {
		keyStorePath := path.Join("./keyStoreFile", keyStoreFile.Filename)
		fmt.Println("keyStorePath", keyStorePath)
		ctx.SaveUploadedFile(keyStoreFile, keyStorePath)
		publicKey, mnemonic, importKeyErr := util.ImportAccountByKetstoreFile(keyStorePath, auth)
		isPublicKeyExist := database.QueryUserByAddress(publicKey).Address
		if importKeyErr != nil {
			ctx.JSON(200, gin.H{
				"err": "导入失败！",
			})
		} else if isPublicKeyExist != "" {
			ctx.JSON(200, gin.H{
				"err": "该账号已存在！",
			})
		} else {
			database.InsertUser(name, password, publicKey, email, newName, "./"+keyStorePath, mnemonic)
		}
	}
	publicKey, keystorePath, mnemonic, err := util.NewAccounts(password)
	if err != nil {
		ctx.JSON(200, err)
	}
	database.InsertUser(name, password, publicKey, email, newName, keystorePath, mnemonic)
	//path, address, mneonic := util.NewAccount("./keystore", password)

	//database.InsertUser(name, password, address, email, newName, path, mneonic)
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
