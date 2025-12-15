package router

import (
	"{{.ProjectName}}/internal/handler"
	"{{.ProjectName}}/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API 路由
	api := r.Group("/api")
	{
		// Item 路由
		itemHandler := handler.NewItemHandler(db)
		items := api.Group("/items")
		{
			items.GET("", itemHandler.GetItems)
			items.GET("/:id", itemHandler.GetItem)
			items.POST("", itemHandler.CreateItem)
			items.PUT("/:id", itemHandler.UpdateItem)
			items.DELETE("/:id", itemHandler.DeleteItem)
		}
	}

	return r
}

