package model

// R 统一 JSON 响应封装
type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(data interface{}) R {
	return R{Code: 0, Message: "success", Data: data}
}

func Fail(msg string) R {
	return R{Code: 1, Message: msg}
}

// SearchResult 单条搜索结果
type SearchResult struct {
	DocID   uint64  `json:"doc_id"`
	Title   string  `json:"title"`   // 可能含 <em> 高亮标签
	URL     string  `json:"url"`
	Snippet string  `json:"snippet"` // 带 <em> 高亮的摘要片段
	Score   float64 `json:"score"`
}

// SearchResponse 搜索接口完整响应
type SearchResponse struct {
	Query   string         `json:"query"`
	Total   int            `json:"total"`   // 总命中数
	Page    int            `json:"page"`
	Size    int            `json:"size"`
	Pages   int            `json:"pages"`   // 总页数
	Results []SearchResult `json:"results"`
	Cost    string         `json:"cost"`    // 耗时（如 "3.2ms"）
	Cached  bool           `json:"cached"`  // 是否来自缓存
}
