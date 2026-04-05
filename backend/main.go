package main

import (
	"log"
	"search-engine/cache"
	"search-engine/config"
	"search-engine/dao"
	"search-engine/db"
	"search-engine/engine"
	"search-engine/handler"
	"search-engine/model"
	"search-engine/router"
	"search-engine/service"
)

func main() {
	// ① 加载配置
	cfg := config.Load()

	// ② 初始化 MySQL
	db.InitMySQL(cfg.MySQL.DSN)
	sqlDB := db.GetDB()

	// 自动建表（若表已存在则跳过）
	if err := sqlDB.AutoMigrate(&model.Document{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// ③ 初始化 Redis（失败时降级，不影响搜索）
	cache.InitRedis(cfg.Redis.Addr, cfg.Redis.Password)

	// ④ 构建内存倒排索引
	idx := engine.NewInvertedIndex()
	docDAO := dao.NewDocDAO(sqlDB)

	docs, err := docDAO.GetAll()
	if err != nil {
		log.Printf("⚠️  加载文档失败: %v", err)
	} else {
		for _, doc := range docs {
			idx.AddDocument(&engine.DocMeta{
				ID:      doc.ID,
				Title:   doc.Title,
				Content: doc.Content,
				URL:     doc.URL,
			})
		}
		log.Printf("✅ 索引构建完成，共加载 %d 篇文档", len(docs))
	}

	// ⑤ 初始化 Service 和 Handler
	docSvc := service.NewDocService(docDAO, idx)
	searchSvc := service.NewSearchService(idx)

	docH := handler.NewDocHandler(docSvc)
	searchH := handler.NewSearchHandler(searchSvc)

	// ⑥ 注册路由并启动服务
	r := router.Setup(docH, searchH)

	addr := ":" + cfg.Server.Port
	log.Printf("🚀 搜索引擎后端启动于 http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
