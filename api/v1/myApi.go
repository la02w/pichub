package v1

import (
	"net/http"
	"pichub/model"
	"pichub/utils"
	"pichub/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	contentType := file.Header.Get("Content-Type")
	code, md5, url := 0, "", ""
	if contentType != "" && len(contentType) > 6 && contentType[:6] == "image/" {
		code, md5 = model.UploadImage(file)
		url = utils.GetHost(c.Request) + "/api/v1/i/" + md5
	} else {
		code = errmsg.ERROR_FILE_TYPE
		url = "null"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"url":  url,
	})
}
func GetImage(c *gin.Context) {
	md5 := c.Param("md5")
	data := model.GetImageData(md5)
	c.Data(http.StatusOK, "image/png", data.Data)
}
