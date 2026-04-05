<template>
  <div class="result-page">
    <!-- ===== 顶部导航栏 ===== -->
    <header class="topbar">
      <RouterLink to="/" class="mini-logo">
        <span class="s">S</span><span class="e1">e</span><span class="a">a</span><span
          class="r">r</span><span class="c">c</span><span class="h">h</span><span
          class="go">Go</span>
      </RouterLink>

      <div class="topbar-search">
        <SearchBar v-model="queryInput" @search="doSearch" size="medium" />
      </div>
    </header>

    <!-- ===== 主内容 ===== -->
    <main class="main">
      <!-- 加载态 -->
      <div v-if="loading" class="skeleton-list">
        <div v-for="i in 5" :key="i" class="skeleton-item">
          <div class="sk sk-url" />
          <div class="sk sk-title" />
          <div class="sk sk-line" />
          <div class="sk sk-line short" />
        </div>
      </div>

      <!-- 错误态 -->
      <div v-else-if="error" class="empty-state">
        <div class="empty-icon">⚠️</div>
        <p>{{ error }}</p>
        <button class="retry-btn" @click="doSearch">重试</button>
      </div>

      <!-- 无结果 -->
      <div v-else-if="resp && resp.total === 0" class="empty-state">
        <div class="empty-icon">🔍</div>
        <p>没有找到与 <strong>"{{ resp.query }}"</strong> 相关的结果</p>
        <p class="empty-tip">建议：检查拼写是否正确，或尝试其他关键词</p>
      </div>

      <!-- 搜索结果 -->
      <template v-else-if="resp">
        <!-- 统计信息栏 -->
        <div class="stats-bar">
          <span>找到约 <strong>{{ resp.total }}</strong> 条结果</span>
          <span class="stats-sep">·</span>
          <span>耗时 {{ resp.cost }}</span>
          <span v-if="resp.cached" class="cache-badge">
            <i class="cache-dot" />来自缓存
          </span>
        </div>

        <!-- 结果列表 -->
        <div class="result-list">
          <ResultItem
            v-for="item in resp.results"
            :key="item.doc_id"
            :item="item"
            :query="resp.query"
          />
        </div>

        <!-- 分页 -->
        <Pagination
          v-if="resp.pages > 1"
          :current="currentPage"
          :total="resp.pages"
          @change="onPageChange"
        />
      </template>
    </main>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { search } from '@/api/search.js'
import SearchBar from '@/components/SearchBar.vue'
import ResultItem from '@/components/ResultItem.vue'
import Pagination from '@/components/Pagination.vue'

const route = useRoute()
const router = useRouter()

const queryInput = ref('')
const resp = ref(null)
const loading = ref(false)
const error = ref('')
const currentPage = ref(1)

async function doSearch() {
  const q = queryInput.value.trim()
  if (!q) return

  // 更新 URL query 参数（保持可分享）
  router.replace({ name: 'search', query: { q, page: currentPage.value } })

  loading.value = true
  error.value = ''
  resp.value = null

  try {
    resp.value = await search(q, currentPage.value, 10)
  } catch (e) {
    error.value = e.message || '搜索服务暂时不可用，请稍后重试'
  } finally {
    loading.value = false
  }
}

function onPageChange(page) {
  currentPage.value = page
  doSearch()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

// 监听 route.query 变化（如浏览器前进/后退）
watch(
  () => route.query,
  (q) => {
    queryInput.value = q.q || ''
    currentPage.value = Number(q.page) || 1
    if (queryInput.value) doSearch()
  },
  { immediate: true },
)
</script>

<style scoped>
.result-page {
  min-height: 100vh;
  background: #fff;
  display: flex;
  flex-direction: column;
}

/* ===== 顶部栏 ===== */
.topbar {
  position: sticky;
  top: 0;
  z-index: 100;
  background: #fff;
  border-bottom: 1px solid #ebebeb;
  padding: 12px 24px;
  display: flex;
  align-items: center;
  gap: 24px;
  box-shadow: 0 1px 6px rgba(0,0,0,.07);
}

.mini-logo {
  font-size: 1.7rem;
  font-weight: 800;
  letter-spacing: -1px;
  flex-shrink: 0;
  display: flex;
}
.s  { color: #4285f4; }
.e1 { color: #ea4335; }
.a  { color: #fbbc05; }
.r  { color: #4285f4; }
.c  { color: #34a853; }
.h  { color: #ea4335; }
.go { color: #1a73e8; }

.topbar-search {
  flex: 1;
  max-width: 620px;
}

/* ===== 主内容 ===== */
.main {
  flex: 1;
  padding: 20px 24px 60px;
  max-width: 860px;
  margin: 0 auto;
  width: 100%;
}

/* ===== 统计栏 ===== */
.stats-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-muted);
  font-size: 13px;
  padding: 4px 0 16px;
  border-bottom: 1px solid #f1f3f4;
  margin-bottom: 8px;
}
.stats-bar strong { color: var(--text); font-weight: 600; }
.stats-sep { color: var(--text-light); }
.cache-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background: #e6f4ea;
  color: #1e7e34;
  border-radius: var(--radius-full);
  padding: 2px 10px;
  font-size: 12px;
}
.cache-dot {
  display: inline-block;
  width: 6px; height: 6px;
  border-radius: 50%;
  background: #34a853;
}

/* ===== 骨架屏 ===== */
.skeleton-list { padding-top: 12px; }
.skeleton-item {
  padding: 16px 0;
  border-bottom: 1px solid #f1f3f4;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.sk {
  background: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
  background-size: 200% 100%;
  animation: shimmer 1.4s infinite;
  border-radius: var(--radius-sm);
}
.sk-url   { width: 180px; height: 13px; }
.sk-title { width: 70%;  height: 20px; }
.sk-line  { width: 100%; height: 14px; }
.sk-line.short { width: 60%; }

@keyframes shimmer {
  0%   { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

/* ===== 空态 ===== */
.empty-state {
  padding: 80px 0;
  text-align: center;
  color: var(--text-muted);
}
.empty-icon { font-size: 3rem; margin-bottom: 16px; }
.empty-state p { font-size: 16px; line-height: 1.8; }
.empty-tip { font-size: 14px; color: var(--text-light); margin-top: 8px; }
.retry-btn {
  margin-top: 20px;
  padding: 8px 24px;
  background: var(--primary);
  color: #fff;
  border-radius: var(--radius-sm);
  font-size: 14px;
  transition: var(--transition);
}
.retry-btn:hover { background: var(--primary-hover); }

/* ===== 结果列表 ===== */
.result-list {
  display: flex;
  flex-direction: column;
}

/* =============================================
   响应式断点
   ============================================= */

/* 平板 ≤ 768px */
@media (max-width: 768px) {
  .topbar {
    padding: 10px 16px;
    gap: 14px;
  }
  .mini-logo { font-size: 1.4rem; }
  .main { padding: 16px 16px 48px; }
  .stats-bar {
    flex-wrap: wrap;
    row-gap: 4px;
  }
}

/* 手机 ≤ 640px */
@media (max-width: 640px) {
  .topbar {
    padding: 8px 12px;
    gap: 10px;
  }
  .mini-logo { font-size: 1.2rem; letter-spacing: -0.5px; }
  .topbar-search { /* 让搜索框尽可能宽 */ }
  .main { padding: 12px 12px 40px; }
  .stats-bar { font-size: 12px; }
  /* 骨架屏标题宽度自适应 */
  .sk-url   { width: 130px; }
  .sk-title { width: 85%; }
  /* 空态 padding 收紧 */
  .empty-state { padding: 48px 0; }
  .empty-icon  { font-size: 2.4rem; }
  .empty-state p { font-size: 14px; }
}

/* 超小屏 ≤ 360px */
@media (max-width: 360px) {
  /* 极小屏隐藏 Logo 文字，只保留搜索框 */
  .mini-logo { display: none; }
  .topbar { padding: 8px 10px; }
  .main { padding: 10px 10px 32px; }
}

/* 大屏 ≥ 1280px：内容区拓宽 */
@media (min-width: 1280px) {
  .main { max-width: 900px; }
  .topbar-search { max-width: 680px; }
}

/* 超宽屏 ≥ 1920px */
@media (min-width: 1920px) {
  .main { max-width: 1100px; }
  .topbar { padding: 14px 40px; }
  .topbar-search { max-width: 800px; }
}
</style>
