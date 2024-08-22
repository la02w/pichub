package router

import (
	v1 "pichub/api/v1"
	"pichub/middleware"
	"pichub/utils"

	"github.com/gin-contrib/multitemplate"

	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/index.html")
	p.AddFromFiles("front", "web/front/index.html")
	return p
}
func InitRouter() {
	r := gin.Default()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Cors())
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/static", "./web/front/")
	r.Static("/admin", "./web/admin/")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})
	router := r.Group("/api/v1")
	{
		router.POST("/upload", v1.UploadImage)
		router.GET("/i/:md5", v1.GetImage)
	}
	_ = r.Run(utils.SERVER_PORT)
}
