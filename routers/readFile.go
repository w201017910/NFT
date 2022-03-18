package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ipfs-app/util"
	"net/http"
	"path"
)

func ReadFile(c *gin.Context) {
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
			c.JSON(http.StatusOK, gin.H{"message": "ok", "url": "https://ipfs.io/ipfs/" + cid})
		}
		fmt.Println(del)
	}
}
