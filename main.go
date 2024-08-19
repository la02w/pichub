package main

import (
	"net/http"
	"path/filepath"
	"pichub/middleware"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/i", "./upload")
	router.Static("/static", "./static")
	router.Use(middleware.Cors())
	router.GET("/", func(c *gin.Context) {
		// 将index.html作为响应发送
		c.File("./static/index.html")
	})
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		fullURL := getHost(c.Request)
		// 指定基础的文件夹路径
		basePath := "./upload/"
		file, _ := c.FormFile("file")
		// 2024/08/16/1723794210.png
		filePath := createFilePath(file.Filename)
		// ./upload/2024/08/16/1723794210.png
		uploadPath := basePath + filePath
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, uploadPath)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			// https://xxx.xx/i/2024/08/16/1723794210.png
			"imageUrl": fullURL + "/i/" + filePath,
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
