package rounter

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"nft/database"
	"nft/util"
	"path"
	"strconv"
)

func ReadFile(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	var username string
	if e == nil {
		username = cookie.Value
	}
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("file_err:", err)
		return
	}
	filepath := path.Join("./images", file.Filename)
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
			type_ := c.PostForm("type")
			name := c.PostForm("name")
			intro := c.PostForm("intro")
			cids := "http://175.178.215.53:8080/ipfs/" + cid
			tokenId := util.MintToken(userInfo.Keystore[2:], common.HexToAddress(userInfo.Address), cids, type_, name)
			id, _ := strconv.Atoi(tokenId.String())
			database.CreateImg(id, name, userInfo.Address, userInfo.Address, cids, type_, intro)
			//c.JSON(http.StatusOK, gin.H{"message": "ok", "url": "https://ipfs.io/ipfs/" + cid})
		}
		fmt.Println(del)
	}
}
