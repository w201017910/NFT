package rounter

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"nft/database"
	"nft/util"
	"path"
	"strconv"
)

var Client *ethclient.Client

func ReadFile(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	var username string
	if e == nil {
		username = cookie.Value
	}
	fmt.Println("name", username)
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("file_err:", err)
		return
	}
	filepath := path.Join("./images", file.Filename)
	fmt.Println(filepath)
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	cid := util.UpLoad(filepath)
	if cid != "" {
		del := util.DelImage(filepath)
		if del == nil {
			userInfo := database.QueryUser(username)
			fmt.Println("Address", userInfo.Address)
			type_ := c.PostForm("type")
			name := c.PostForm("name")
			intro := c.PostForm("intro")
			cids := "https://ipfs.io/ipfs/" + cid
			tokenId := util.MintToken(Client, userInfo.Keystore[2:], common.HexToAddress(userInfo.Address), cids, type_, name)
			id, _ := strconv.Atoi(tokenId.String())
			database.CreateImg(id, userInfo.Address, userInfo.Address, cids, type_, intro)
			c.JSON(http.StatusOK, gin.H{"tokenId": id, "message": "ok", "url": "https://ipfs.io/ipfs/" + cid})
		}
		fmt.Println(del)
	}
}
