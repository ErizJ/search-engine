/**
 * Mock 搜索引擎 —— 在浏览器端模拟倒排索引 + TF-IDF 评分
 * 无需后端即可预览完整 UI 效果
 */
import { mockDocuments } from './data.js'

// ── 工具函数 ────────────────────────────────────────────────

/** 模拟网络延迟 */
const delay = (ms) => new Promise((r) => setTimeout(r, ms))

/** 简易中英文分词 */
function tokenize(text) {
  if (!text) return []
  const lower = text.toLowerCase()
  const tokens = []

  // 英文单词
  const engWords = lower.match(/[a-z0-9]+/g) || []
  tokens.push(...engWords.filter((w) => w.length > 1))

  // 中文 bigram
  const chineseChars = lower.match(/[\u4e00-\u9fff]+/g) || []
  chineseChars.forEach((seg) => {
    for (let i = 0; i < seg.length; i++) {
      tokens.push(seg[i]) // 单字
      if (i + 1 < seg.length) tokens.push(seg.slice(i, i + 2)) // bigram
    }
  })

  // 去重
  return [...new Set(tokens)]
}

/** 统计 token 在文本中出现次数 */
function countOccurrences(text, token) {
  const re = new RegExp(token.replace(/[.*+?^${}()|[\]\\]/g, '\\$&'), 'gi')
  return (text.match(re) || []).length
}

/** 生成带高亮的摘要片段 */
function generateSnippet(content, tokens, maxLen = 180) {
  if (!content) return ''
  const lower = content.toLowerCase()

  // 找最早命中位置
  let bestPos = -1
  for (const token of tokens) {
    const idx = lower.indexOf(token)
    if (idx !== -1 && (bestPos === -1 || idx < bestPos)) bestPos = idx
  }

  // 截取上下文
  let start = Math.max(0, (bestPos === -1 ? 0 : bestPos) - 40)
  let snippet = content.slice(start, start + maxLen)
  if (start > 0) snippet = '...' + snippet
  if (start + maxLen < content.length) snippet = snippet + '...'

  // 高亮关键词
  tokens.forEach((token) => {
    const re = new RegExp(`(${token.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')})`, 'gi')
    snippet = snippet.replace(re, '<em>$1</em>')
  })
  return snippet
}

/** 高亮标题 */
function highlightTitle(title, tokens) {
  let result = title
  tokens.forEach((token) => {
    const re = new RegExp(`(${token.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')})`, 'gi')
    result = result.replace(re, '<em>$1</em>')
  })
  return result
}

// ── Mock API ────────────────────────────────────────────────

/** 模拟文档列表（内存副本，支持增删）*/
let docs = mockDocuments.map((d) => ({ ...d }))
let nextId = docs.length + 1

/** 搜索 */
export async function mockSearch(query, page = 1, size = 10) {
  await delay(280 + Math.random() * 220) // 模拟 280~500ms 网络延迟

  const t0 = performance.now()
  const tokens = tokenize(query)

  if (!tokens.length) {
    return { query, total: 0, page, size, pages: 0, results: [], cost: '0ms', cached: false }
  }

  // 计算每篇文档的 TF-IDF 近似得分
  const N = docs.length
  const scored = docs
    .map((doc) => {
      const fullText = doc.title + ' ' + doc.content
      const totalTerms = fullText.split(/\s+/).length

      let score = 0
      for (const token of tokens) {
        const tf = countOccurrences(fullText, token) / totalTerms
        // IDF: log(N / (df + 1)) + 1，df 用所有文档中含该词的文档数
        const df = docs.filter((d) =>
          (d.title + ' ' + d.content).toLowerCase().includes(token),
        ).length
        const idf = Math.log(N / (df + 1)) + 1
        score += tf * idf
      }
      return { doc, score }
    })
    .filter(({ score }) => score > 0)
    .sort((a, b) => b.score - a.score)

  const total = scored.length
  const pages = Math.ceil(total / size)
  const offset = (page - 1) * size

  const results = scored.slice(offset, offset + size).map(({ doc, score }) => ({
    doc_id: doc.id,
    title: highlightTitle(doc.title, tokens),
    url: doc.url,
    snippet: generateSnippet(doc.content, tokens),
    score: parseFloat(score.toFixed(4)),
  }))

  const cost = `${(performance.now() - t0).toFixed(2)}ms`
  return { query, total, page, size, pages, results, cost, cached: false }
}

/** 文档列表 */
export async function mockListDocs() {
  await delay(150)
  return [...docs]
}

/** 新增文档 */
export async function mockCreateDoc(data) {
  await delay(200)
  const doc = { id: nextId++, ...data, created_at: new Date().toISOString(), updated_at: new Date().toISOString() }
  docs.push(doc)
  return doc
}

/** 删除文档 */
export async function mockDeleteDoc(id) {
  await delay(150)
  docs = docs.filter((d) => d.id !== id)
  return null
}

/** 重建索引（Mock 中无实际操作，直接返回） */
export async function mockRebuildIndex() {
  await delay(500)
  return { indexed: docs.length }
}
