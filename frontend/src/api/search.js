import axios from 'axios'
import {
  mockSearch,
  mockListDocs,
  mockCreateDoc,
  mockDeleteDoc,
  mockRebuildIndex,
} from '@/mock/search.js'

// VITE_MOCK=true 时走 Mock，无需启动后端
const isMock = import.meta.env.VITE_MOCK === 'true'

// ── 真实 HTTP 客户端 ────────────────────────────────────────
const http = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
})

http.interceptors.response.use(
  (res) => {
    if (res.data.code === 0) return res.data.data
    return Promise.reject(new Error(res.data.message || '请求失败'))
  },
  (err) => Promise.reject(err),
)

// ── 导出接口（Mock / Real 统一入口）────────────────────────

/**
 * 搜索
 * @param {string} query
 * @param {number} page
 * @param {number} size
 */
export function search(query, page = 1, size = 10) {
  if (isMock) return mockSearch(query, page, size)
  return http.get('/search', { params: { q: query, page, size } })
}

/** 文档列表 */
export function listDocs() {
  if (isMock) return mockListDocs()
  return http.get('/docs')
}

/** 新增文档 */
export function createDoc(data) {
  if (isMock) return mockCreateDoc(data)
  return http.post('/docs', data)
}

/** 删除文档 */
export function deleteDoc(id) {
  if (isMock) return mockDeleteDoc(id)
  return http.delete(`/docs/${id}`)
}

/** 重建索引 */
export function rebuildIndex() {
  if (isMock) return mockRebuildIndex()
  return http.post('/index/rebuild')
}
