package v1

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"pichub/model"
	"pichub/utils"
	"pichub/utils/errmsg"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
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
	if data.Size > 1048576 && data.Type == "image/jpeg" || data.Type == "image/png" {
		img, _, err := image.Decode(bytes.NewReader(data.Data))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		thumbnail := resize.Thumbnail(200, 200, img, resize.Lanczos3)
		var buf bytes.Buffer
		if data.Type == "image/jpeg" {
			err = jpeg.Encode(&buf, thumbnail, nil)
		}
		if data.Type == "image/png" {
			err = png.Encode(&buf, thumbnail)
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, data.Type, buf.Bytes())
	}
	c.Data(http.StatusOK, data.Type, data.Data)
}
