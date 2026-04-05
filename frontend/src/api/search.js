import axios from 'axios'
import {
  mockSearch,
  mockListDocs,
  mockCreateDoc,
  mockDeleteDoc,
  mockRebuildIndex,
} from '@/mock/search.js'

// VITE_MOCK=true → Mock 模式（npm run dev）
// VITE_MOCK=false → 真实后端（npm run dev:real）
export const isMock = import.meta.env.VITE_MOCK === 'true'

// ── 真实 HTTP 客户端 ────────────────────────────────────────
const http = axios.create({
  baseURL: '/api/v1',
  timeout: 8000,
})

// 响应拦截器：统一处理后端 R 包装 + 友好错误提示
http.interceptors.response.use(
  (res) => {
    // 业务成功：{ code:0, data:... }
    if (res.data.code === 0) return res.data.data
    // 业务失败：{ code:1, message:"..." }
    return Promise.reject(new Error(res.data.message || '请求失败'))
  },
  (err) => {
    // 请求超时
    if (err.code === 'ECONNABORTED' || err.message?.includes('timeout')) {
      return Promise.reject(new Error('请求超时，请稍后重试'))
    }
    // 无法连接到后端（后端未启动 / 网络断开）
    if (err.code === 'ERR_NETWORK' || !err.response) {
      return Promise.reject(
        new Error('无法连接后端服务，请确认后端已启动（go run main.go，端口 8080）'),
      )
    }
    // HTTP 5xx 服务器错误
    if (err.response.status >= 500) {
      return Promise.reject(new Error(`服务器内部错误（${err.response.status}），请稍后重试`))
    }
    // 其他（4xx 等）：尽量读取后端返回的 message
    const msg = err.response.data?.message || err.message || '请求失败'
    return Promise.reject(new Error(msg))
  },
)

// ── 导出接口（Mock / Real 统一入口）────────────────────────

/** 搜索 */
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
