package service

import (
	"fmt"
	"search-engine/cache"
	"search-engine/dao"
	"search-engine/engine"
	"search-engine/model"
)

// DocService 文档管理服务
// 负责保持 MySQL 和内存索引的一致性
type DocService struct {
	dao *dao.DocDAO
	idx *engine.InvertedIndex
}

func NewDocService(dao *dao.DocDAO, idx *engine.InvertedIndex) *DocService {
	return &DocService{dao: dao, idx: idx}
}

// CreateDoc 创建文档：写 MySQL → 加入内存索引 → 清空相关缓存
func (s *DocService) CreateDoc(req *model.CreateDocRequest) (*model.Document, error) {
	doc := &model.Document{
		Title:   req.Title,
		Content: req.Content,
		URL:     req.URL,
	}
	if err := s.dao.Create(doc); err != nil {
		return nil, fmt.Errorf("写入数据库失败: %w", err)
	}

	// 同步到内存索引
	s.idx.AddDocument(&engine.DocMeta{
		ID:      doc.ID,
		Title:   doc.Title,
		Content: doc.Content,
		URL:     doc.URL,
	})

	// 文档变更后清空搜索缓存，避免返回旧结果
	cache.Del("search:*")

	return doc, nil
}

// DeleteDoc 删除文档：从 MySQL 删除 → 从内存索引删除
func (s *DocService) DeleteDoc(id uint64) error {
	if _, err := s.dao.GetByID(id); err != nil {
		return fmt.Errorf("文档不存在: %w", err)
	}
	if err := s.dao.Delete(id); err != nil {
		return fmt.Errorf("删除失败: %w", err)
	}
	s.idx.RemoveDocument(id)
	cache.Del("search:*")
	return nil
}

// GetAllDocs 查询所有文档列表
func (s *DocService) GetAllDocs() ([]*model.Document, error) {
	return s.dao.GetAll()
}

// RebuildIndex 重新从 MySQL 加载所有文档并重建索引
func (s *DocService) RebuildIndex() (int, error) {
	docs, err := s.dao.GetAll()
	if err != nil {
		return 0, err
	}
	// 重建：先用新索引替换（简单方案：逐个 Add，AddDocument 内部会先删旧记录）
	for _, doc := range docs {
		s.idx.AddDocument(&engine.DocMeta{
			ID:      doc.ID,
			Title:   doc.Title,
			Content: doc.Content,
			URL:     doc.URL,
		})
	}
	cache.Del("search:*")
	return len(docs), nil
}
