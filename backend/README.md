# {{.ProjectName}}

{{.ProjectDescription}}

## 技术栈

- **Go** 1.23
- **Gin** - Web 框架
- **GORM** - ORM
- **PostgreSQL** - 数据库

## 快速开始

### 1. 安装依赖

```bash
go mod download
```

### 2. 配置环境变量

复制 `.env.example` 到 `.env` 并修改配置：

```bash
cp .env.example .env
```

### 3. 运行项目

```bash
go run cmd/api/main.go
```

服务器将在 `http://localhost:{{.BackendPort}}` 启动

## API 端点

### 健康检查
```
GET /health
```

### Items API
```
GET    /api/items      # 获取所有物品
GET    /api/items/:id  # 获取单个物品
POST   /api/items      # 创建物品
PUT    /api/items/:id  # 更新物品
DELETE /api/items/:id  # 删除物品
```

### 示例请求

创建物品：
```bash
curl -X POST http://localhost:{{.BackendPort}}/api/items \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sample Item",
    "description": "This is a sample item",
    "price": 29.99
  }'
```

获取所有物品：
```bash
curl http://localhost:{{.BackendPort}}/api/items
```

## Docker 部署

### 构建镜像

```bash
docker build -t {{.ProjectSlug}} .
```

### 运行容器

```bash
docker run -p {{.BackendPort}}:{{.BackendPort}} \
  -e DATABASE_URL="your-database-url" \
  {{.ProjectSlug}}
```

## 项目结构

```
.
├── cmd/
│   └── api/
│       └── main.go          # 应用入口
├── config/
│   └── config.go            # 配置管理
├── internal/
│   ├── handler/             # HTTP 处理器
│   ├── middleware/          # 中间件
│   ├── model/               # 数据模型
│   └── router/              # 路由配置
├── pkg/
│   ├── logger/              # 日志工具
│   └── response/            # 响应工具
├── .env.example             # 环境变量示例
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## 开发

### 添加新的 API 端点

1. 在 `internal/model` 中定义数据模型
2. 在 `internal/handler` 中创建处理器
3. 在 `internal/router` 中注册路由

### 数据库迁移

GORM 会自动迁移数据库schema。在 `main.go` 中添加你的模型：

```go
db.AutoMigrate(&model.YourModel{})
```

## License

MIT

