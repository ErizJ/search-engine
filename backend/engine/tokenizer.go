package engine

import "strings"

// stopWords 英文停用词表（高频但无检索价值的词）
var stopWords = map[string]bool{
	"the": true, "a": true, "an": true, "and": true, "or": true,
	"but": true, "in": true, "on": true, "at": true, "to": true,
	"for": true, "of": true, "is": true, "are": true, "was": true,
	"be": true, "it": true, "that": true, "this": true, "with": true,
	"as": true, "by": true, "from": true,
}

// isChinese 判断字符是否为 CJK 统一汉字
func isChinese(r rune) bool {
	return r >= 0x4E00 && r <= 0x9FFF
}

// isAlphanumeric 字母或数字
func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

// Tokenize 将文本切分为 token 列表（保留重复项，用于 TF 计算）
//
// 策略：
//   - 英文/数字：按非字母分隔，转小写，过滤停用词
//   - 中文：相邻字符组成 Bigram（"机器学习" → "机器","器学","学习"）
//     同时保留单字，以支持单字搜索
func Tokenize(text string) []string {
	text = strings.ToLower(text)

	var tokens []string
	runes := []rune(text)
	n := len(runes)

	var wordBuf strings.Builder
	var chineseBuf []rune

	// flushWord：把累积的英文/数字词入队
	flushWord := func() {
		if wordBuf.Len() == 0 {
			return
		}
		w := wordBuf.String()
		wordBuf.Reset()
		if !stopWords[w] && len(w) > 1 {
			tokens = append(tokens, w)
		}
	}

	// flushChinese：对中文段落做 Bigram + 单字 tokenization
	flushChinese := func() {
		if len(chineseBuf) == 0 {
			return
		}
		// 单字
		for _, c := range chineseBuf {
			tokens = append(tokens, string(c))
		}
		// Bigram
		for i := 0; i+1 < len(chineseBuf); i++ {
			tokens = append(tokens, string(chineseBuf[i:i+2]))
		}
		chineseBuf = chineseBuf[:0]
	}

	for i := 0; i < n; i++ {
		r := runes[i]
		if isChinese(r) {
			flushWord()
			chineseBuf = append(chineseBuf, r)
		} else if isAlphanumeric(r) {
			flushChinese()
			wordBuf.WriteRune(r)
		} else {
			// 分隔符：刷新两个缓冲区
			flushWord()
			flushChinese()
		}
	}
	flushWord()
	flushChinese()

	return tokens
}

// UniqueTokens 返回去重后的 token 集合（用于查询时去重，避免重复计算 IDF）
func UniqueTokens(text string) []string {
	tokens := Tokenize(text)
	seen := make(map[string]struct{}, len(tokens))
	result := tokens[:0]
	for _, t := range tokens {
		if _, ok := seen[t]; !ok {
			seen[t] = struct{}{}
			result = append(result, t)
		}
	}
	return result
}
