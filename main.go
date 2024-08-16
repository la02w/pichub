package main

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/i", "./upload")
	router.Static("/static", "./static")
	// router.Use(middleware.Cors())
	router.GET("/", func(c *gin.Context) {
		// 将index.html作为响应发送
		c.File("./static/index.html")
	})
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
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
			// https://pic.la02.cc/i/2024/08/16/1723794210.png
			"imageUrl": filepath.Join("https://pic.la02.cc/i/", filePath),
		})
	})
	router.Run(":2356")
}
func createFilePath(filename string) string {
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	timestamp := now.Unix()
	datePath := filepath.Join(year, month, day) + "/"
	return datePath + strconv.Itoa(int(timestamp)) + filepath.Ext(filename)
}