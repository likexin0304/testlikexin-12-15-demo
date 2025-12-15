package main

import (
	"log"
	"os"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/internal/model"
	"{{.ProjectName}}/internal/router"
	"{{.ProjectName}}/pkg/logger"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 初始化日志
	logger.Init()

	// 加载配置
	cfg := config.Load()

	// 连接数据库
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移数据库
	if err := db.AutoMigrate(&model.Item{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	logger.Info("Database connected and migrated successfully")

	// 设置路由
	r := router.SetupRouter(db)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "{{.BackendPort}}"
	}

	logger.Info("Starting server on port " + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

