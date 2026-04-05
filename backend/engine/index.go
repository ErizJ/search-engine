package engine

import "sync"

// -------------------------------------------------------
// 核心数据结构
// -------------------------------------------------------

// Posting 某词在某文档中的出现记录
type Posting struct {
	DocID     uint64  // 文档 ID
	TF        float64 // 归一化词频 = count(term,doc) / totalTerms(doc)
	Positions []int   // 出现位置（token 序号），用于生成摘要
}

// PostingList 某个词的倒排列表
type PostingList struct {
	DocFreq  int        // 包含该词的文档数，用于 IDF 计算
	Postings []*Posting // 倒排记录列表
}

// DocMeta 存储在内存索引中的文档元信息
type DocMeta struct {
	ID        uint64
	Title     string
	Content   string
	URL       string
	TermCount int // 文档总 token 数（含重复），用于 TF 分母
}

// InvertedIndex 线程安全的倒排索引主结构（单例）
//
//	内存布局：
//	  index:    map[token] -> PostingList
//	  docStore: map[docID] -> DocMeta
//	  docCount: 总文档数（用于 IDF 的 N）
type InvertedIndex struct {
	mu       sync.RWMutex
	index    map[string]*PostingList
	docStore map[uint64]*DocMeta
	docCount int
}

// NewInvertedIndex 创建一个空的倒排索引
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		index:    make(map[string]*PostingList),
		docStore: make(map[uint64]*DocMeta),
	}
}

// -------------------------------------------------------
// 写操作
// -------------------------------------------------------

// AddDocument 将文档加入倒排索引
//
// 流程：
//  1. 对 title+content 分词（含重复 token，用于 TF 计数）
//  2. 统计每个 token 的出现次数和位置
//  3. 计算 TF，写入对应词的 PostingList
func (idx *InvertedIndex) AddDocument(doc *DocMeta) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	// 若文档已存在，先移除旧索引避免重复
	if _, exists := idx.docStore[doc.ID]; exists {
		idx.removeDocLocked(doc.ID)
	}

	fullText := doc.Title + " " + doc.Content
	allTokens := Tokenize(fullText) // 保留重复项

	doc.TermCount = len(allTokens)
	idx.docStore[doc.ID] = doc
	idx.docCount++

	// 统计词频和位置
	termFreq := make(map[string]int)
	termPos := make(map[string][]int)
	for pos, token := range allTokens {
		termFreq[token]++
		termPos[token] = append(termPos[token], pos)
	}

	// 写入倒排列表
	totalTerms := float64(len(allTokens))
	for term, count := range termFreq {
		tf := float64(count) / totalTerms

		posting := &Posting{
			DocID:     doc.ID,
			TF:        tf,
			Positions: termPos[term],
		}

		pl, ok := idx.index[term]
		if !ok {
			pl = &PostingList{}
			idx.index[term] = pl
		}
		pl.Postings = append(pl.Postings, posting)
		pl.DocFreq++
	}
}

// RemoveDocument 从倒排索引中删除文档
func (idx *InvertedIndex) RemoveDocument(docID uint64) {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	idx.removeDocLocked(docID)
}

// removeDocLocked 内部删除逻辑（调用方须持有写锁）
func (idx *InvertedIndex) removeDocLocked(docID uint64) {
	if _, exists := idx.docStore[docID]; !exists {
		return
	}
	delete(idx.docStore, docID)
	idx.docCount--

	for term, pl := range idx.index {
		newPostings := pl.Postings[:0]
		removed := false
		for _, p := range pl.Postings {
			if p.DocID != docID {
				newPostings = append(newPostings, p)
			} else {
				removed = true
			}
		}
		pl.Postings = newPostings
		if removed {
			pl.DocFreq--
		}
		if pl.DocFreq == 0 {
			delete(idx.index, term)
		}
	}
}

// -------------------------------------------------------
// 读操作
// -------------------------------------------------------

// GetDoc 获取文档元信息（无锁外部调用请持有 RLock）
func (idx *InvertedIndex) GetDoc(docID uint64) *DocMeta {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	return idx.docStore[docID]
}

// GetPostings 获取某词的倒排列表（内部查询使用，Search 方法自行加锁）
func (idx *InvertedIndex) getPostings(term string) *PostingList {
	return idx.index[term]
}

// DocCount 返回当前已索引文档总数
func (idx *InvertedIndex) DocCount() int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	return idx.docCount
}
