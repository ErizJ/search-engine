package router

import (
	"net/http"
	"search-engine/handler"
	"search-engine/middleware"

	"github.com/gin-gonic/gin"
)

// Setup 注册所有路由
func Setup(docH *handler.DocHandler, searchH *handler.SearchHandler) *gin.Engine {
	r := gin.Default()

	// 跨域中间件
	r.Use(middleware.CORS())

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"pong": true})
	})

	v1 := r.Group("/api/v1")
	{
		// 搜索
		v1.GET("/search", searchH.Search)

		// 文档管理
		docs := v1.Group("/docs")
		{
			docs.GET("", docH.ListDocs)
			docs.POST("", docH.CreateDoc)
			docs.DELETE("/:id", docH.DeleteDoc)
		}

		// 索引管理
		v1.POST("/index/rebuild", docH.RebuildIndex)
	}

	return r
}
