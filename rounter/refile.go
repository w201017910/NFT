package rounter

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
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
				isCidExist := database.QueryImgByCid("http://175.178.215.53:8080/ipfs/" + cid)
				fmt.Println("isCidExist != nil: ", isCidExist != nil)
				if isCidExist != nil {
					c.JSON(http.StatusOK, gin.H{
						"exist": "该作品在链上已存在！",
					})
				} else {
					userInfo := database.QueryUser(username)
					type_ := c.PostForm("type")
					name := c.PostForm("name")
					intro := c.PostForm("intro")
					cids := "http://175.178.215.53:8080/ipfs/" + cid
					opts, _ := util.BindOptsByKeystore(database.QueryUser(cookie.Value).Keystore, database.QueryUser(cookie.Value).Password)
					tokenId, mintErr := util.MintToken(opts, common.HexToAddress(userInfo.Address), cids, type_, name)
					if mintErr != nil {
						c.JSON(200, gin.H{
							"err": "铸币失败！",
						})
					} else {
						id, _ := strconv.Atoi(tokenId.String())
						database.CreateImg(id, name, userInfo.Address, userInfo.Address, cids, type_, intro)
						c.JSON(http.StatusOK, gin.H{"message": "ok"})
					}
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "false",
				})
			}
		}
	} else {
		c.JSON(200, gin.H{
			"err": "未登录！",
		})
	}
}
