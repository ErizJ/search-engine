package service

import (
	"encoding/json"
	"fmt"
	"search-engine/cache"
	"search-engine/engine"
	"search-engine/model"
	"time"
)

const (
	cachePrefix = "search:"
	cacheTTL    = 5 * time.Minute
	snippetLen  = 180 // 摘要最大字符数
)

// SearchService 搜索核心服务
type SearchService struct {
	idx *engine.InvertedIndex
}

func NewSearchService(idx *engine.InvertedIndex) *SearchService {
	return &SearchService{idx: idx}
}

// Search 执行搜索，支持 Redis 缓存和分页
//
// 完整流程：
//  1. 查 Redis 缓存 → 命中直接返回
//  2. 对 query 分词 → 查倒排索引 → TF-IDF 排序
//  3. 分页截取结果 → 生成摘要高亮
//  4. 写入 Redis 缓存
func (s *SearchService) Search(query string, page, size int) (*model.SearchResponse, error) {
	start := time.Now()

	// 参数校正
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 50 {
		size = 10
	}

	// ① 尝试读取 Redis 缓存
	cacheKey := fmt.Sprintf("%s%s:p%d:s%d", cachePrefix, query, page, size)
	if cached, err := cache.Get(cacheKey); err == nil {
		var resp model.SearchResponse
		if json.Unmarshal([]byte(cached), &resp) == nil {
			resp.Cached = true
			return &resp, nil
		}
	}

	// ② 查询倒排索引（返回按 TF-IDF 分数降序排列的全部结果）
	scoredDocs := s.idx.Search(query)
	total := len(scoredDocs)

	// ③ 分页
	offset := (page - 1) * size
	pages := (total + size - 1) / size

	var pageDocs []engine.ScoredDoc
	if offset < total {
		end := offset + size
		if end > total {
			end = total
		}
		pageDocs = scoredDocs[offset:end]
	}

	// ④ 构建结果（生成摘要 + 高亮）
	results := make([]model.SearchResult, 0, len(pageDocs))
	for _, sd := range pageDocs {
		doc := s.idx.GetDoc(sd.DocID)
		if doc == nil {
			continue
		}
		results = append(results, model.SearchResult{
			DocID:   doc.ID,
			Title:   engine.HighlightTitle(doc.Title, query),
			URL:     doc.URL,
			Snippet: engine.GenerateSnippet(doc.Content, query, snippetLen),
			Score:   sd.Score,
		})
	}

	resp := &model.SearchResponse{
		Query:   query,
		Total:   total,
		Page:    page,
		Size:    size,
		Pages:   pages,
		Results: results,
		Cost:    formatDuration(time.Since(start)),
		Cached:  false,
	}

	// ⑤ 写入 Redis 缓存（序列化失败不影响返回）
	if data, err := json.Marshal(resp); err == nil {
		_ = cache.Set(cacheKey, string(data), cacheTTL)
	}

	return resp, nil
}

// formatDuration 格式化耗时为可读字符串，如 "1.23ms"
func formatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return fmt.Sprintf("%.2fμs", float64(d.Microseconds()))
	}
	return fmt.Sprintf("%.2fms", float64(d.Microseconds())/1000)
}

