package engine

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

// ScoredDoc 搜索结果中的文档 + TF-IDF 分数
type ScoredDoc struct {
	DocID uint64
	Score float64
}

// -------------------------------------------------------
// TF-IDF 搜索
// -------------------------------------------------------

// Search 执行 TF-IDF 相关性搜索
//
// 公式：
//
//	TF(t,d)  = count(t,d) / |d|                     词频，归一化
//	IDF(t)   = log(N / (df(t) + 1)) + 1             逆文档频率（+1 平滑，防止 df=N 时 IDF=0）
//	Score(d) = Σ TF(t,d) × IDF(t)   for t in query
//
// 返回按分数降序排列的文档列表。
func (idx *InvertedIndex) Search(query string) []ScoredDoc {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if idx.docCount == 0 {
		return nil
	}

	// 查询词去重，避免同一词多次累加分数
	queryTokens := uniqueSlice(Tokenize(query))
	if len(queryTokens) == 0 {
		return nil
	}

	N := float64(idx.docCount)
	scores := make(map[uint64]float64, 64)

	for _, token := range queryTokens {
		pl := idx.getPostings(token)
		if pl == nil || pl.DocFreq == 0 {
			continue
		}

		// IDF：稀有词权重高，常见词权重低
		idf := math.Log(N/float64(pl.DocFreq+1)) + 1.0

		for _, posting := range pl.Postings {
			scores[posting.DocID] += posting.TF * idf
		}
	}

	// 构建结果并按分数降序排列
	result := make([]ScoredDoc, 0, len(scores))
	for docID, score := range scores {
		result = append(result, ScoredDoc{DocID: docID, Score: score})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	return result
}

// -------------------------------------------------------
// 摘要 & 高亮
// -------------------------------------------------------

// GenerateSnippet 从 content 中提取包含查询词的摘要片段，并用 <em> 高亮关键词
//
// 策略：
//  1. 找到 query 中任意 token 在文本中的第一个出现位置
//  2. 以该位置为中心截取 maxLen 个字符
//  3. 用正则将所有 query token 替换为 <em>token</em>
func GenerateSnippet(content, query string, maxLen int) string {
	tokens := UniqueTokens(query)
	if len(tokens) == 0 || len(content) == 0 {
		runes := []rune(content)
		if len(runes) > maxLen {
			return string(runes[:maxLen]) + "..."
		}
		return content
	}

	runes := []rune(content)
	lowerContent := strings.ToLower(string(runes))

	// 找最早出现的 token 位置（rune 级别）
	bestRunePos := -1
	for _, token := range tokens {
		bytePos := strings.Index(lowerContent, token)
		if bytePos == -1 {
			continue
		}
		runePos := len([]rune(string(runes)[:bytePos])) // byte → rune 位置
		// 用 strings.Index 返回的是字节偏移，需转换
		// 直接遍历 lowerContent 的 rune 来找
		runePos = runeIndexOf([]rune(lowerContent), []rune(token))
		if runePos != -1 && (bestRunePos == -1 || runePos < bestRunePos) {
			bestRunePos = runePos
		}
	}

	// 确定截取范围
	start := 0
	if bestRunePos > 40 {
		start = bestRunePos - 40
	}
	end := start + maxLen
	if end > len(runes) {
		end = len(runes)
	}

	snippet := string(runes[start:end])
	prefix, suffix := "", ""
	if start > 0 {
		prefix = "..."
	}
	if end < len(runes) {
		suffix = "..."
	}

	// 高亮：将所有 token 替换为 <em>token</em>（大小写不敏感）
	for _, token := range tokens {
		if token == "" {
			continue
		}
		re, err := regexp.Compile(`(?i)` + regexp.QuoteMeta(token))
		if err != nil {
			continue
		}
		snippet = re.ReplaceAllStringFunc(snippet, func(match string) string {
			return "<em>" + match + "</em>"
		})
	}

	return prefix + snippet + suffix
}

// HighlightTitle 对标题进行关键词高亮
func HighlightTitle(title, query string) string {
	tokens := UniqueTokens(query)
	result := title
	for _, token := range tokens {
		if token == "" {
			continue
		}
		re, err := regexp.Compile(`(?i)` + regexp.QuoteMeta(token))
		if err != nil {
			continue
		}
		result = re.ReplaceAllStringFunc(result, func(match string) string {
			return "<em>" + match + "</em>"
		})
	}
	return result
}

// -------------------------------------------------------
// 内部工具函数
// -------------------------------------------------------

// uniqueSlice 去重保序
func uniqueSlice(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	out := s[:0]
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// runeIndexOf 在 rune slice 中查找子 slice 的位置，不存在返回 -1
func runeIndexOf(text, pattern []rune) int {
	if len(pattern) == 0 {
		return 0
	}
	for i := 0; i+len(pattern) <= len(text); i++ {
		match := true
		for j, p := range pattern {
			if text[i+j] != p {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
