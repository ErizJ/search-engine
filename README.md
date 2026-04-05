# SearchGo 搜索引擎

基于 Go + Vue3 的轻量级搜索引擎，实现了倒排索引 + TF-IDF 排序。

## 快速启动

### 1. 启动 MySQL + Redis
```bash
docker-compose up -d mysql redis
```

### 2. 启动后端
```bash
cd backend
go mod tidy
go run main.go
```

### 3. 启动前端
```bash
cd frontend
npm install
npm run dev
```

浏览器打开 http://localhost:5173

---

## 环境变量（后端）

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `SERVER_PORT` | `8080` | 服务端口 |
| `MYSQL_DSN` | `root:password@tcp(127.0.0.1:3306)/searchengine?...` | MySQL 连接串 |
| `REDIS_ADDR` | `127.0.0.1:6379` | Redis 地址 |
| `REDIS_PASSWORD` | `` | Redis 密码 |

---

## API 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET  | `/api/v1/search?q=xxx&page=1&size=10` | 搜索 |
| GET  | `/api/v1/docs` | 文档列表 |
| POST | `/api/v1/docs` | 新增文档 |
| DELETE | `/api/v1/docs/:id` | 删除文档 |
| POST | `/api/v1/index/rebuild` | 重建索引 |

### 搜索示例
```bash
curl "http://localhost:8080/api/v1/search?q=golang&page=1&size=5"
```

### 新增文档示例
```bash
curl -X POST http://localhost:8080/api/v1/docs \
  -H "Content-Type: application/json" \
  -d '{"title":"我的文章","content":"这是文章内容","url":"https://example.com"}'
```
