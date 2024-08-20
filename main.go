package main

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"pichub/middleware"
	"pichub/utils"
	"strconv"
	"time"

	"context"
	"net/url"

	"github.com/gin-gonic/gin"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	router := gin.Default()
	router.Static("/i", "./upload")
	router.Static("/static", "./web")
	router.Use(middleware.Cors())
	router.GET("/", func(c *gin.Context) {
		// 将index.html作为响应发送
		c.File("./web/index.html")
	})
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		code := 415
		msg := "抱歉,您上传的数据,格式不支持！"
		contentType := file.Header.Get("Content-Type")
		allowedMimeTypes := []string{"image/jpeg", "image/png", "image/gif", "image/bmp"}
		for _, mime := range allowedMimeTypes {
			if contentType == mime {
				switch {
				case utils.PICK_SERVICE == "local":
					msg = localUpload(utils.LOCAL_BASE_FOLDER, file, c)
				case utils.PICK_SERVICE == "tencent":
					msg = cosUpload(file.Filename, file)
				}
				code = 200
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
	})
	router.Run(utils.SERVER_PORT)
}

func getHost(request *http.Request) string {
	scheme := "http"
	if request.TLS != nil {
		scheme = "https"
	}
	host := request.Host
	return scheme + "://" + host
}

func createFilePath(filename string) string {
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	timestamp := now.Unix()
	datePath := year + "/" + month + "/" + day + "/"
	return datePath + strconv.Itoa(int(timestamp)) + filepath.Ext(filename)
}

func cosUpload(filename string, file *multipart.FileHeader) string {
	u, _ := url.Parse(utils.TENCENT_COS_URL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.TENCENT_COS_SECRETID,
			SecretKey: utils.TENCENT_COS_SECRETKEY,
		},
	})
	name := "pichub/" + createFilePath(filename)
	src, _ := file.Open()
	defer src.Close()
	_, err := c.Object.Put(context.Background(), name, src, nil)
	if err != nil {
		panic(err)
	}
	return utils.TENCENT_COS_URL + "/" + name

}

func localUpload(basePath string, file *multipart.FileHeader, c *gin.Context) string { // 指定basePath基础的文件夹路径
	fullURL := getHost(c.Request)
	filePath := createFilePath(file.Filename) // 2024/08/16/1723794210.png
	uploadPath := basePath + filePath         // ./upload/2024/08/16/1723794210.png
	c.SaveUploadedFile(file, uploadPath)      // 上传文件至指定的完整文件路径
	return fullURL + "/i/" + filePath
}
