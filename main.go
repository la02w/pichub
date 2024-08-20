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
		url := ""
		switch {
		case utils.PICK_SERVICE == "local":
			url = localUpload(utils.LOCAL_BASE_FOLDER, file, c)
		case utils.PICK_SERVICE == "tencent":
			url = cosUpload(file.Filename, file)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   200,
			"imageUrl": url,
		})
	})
	router.Run(":2356")
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
