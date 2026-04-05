<template>
  <div class="home">
    <!-- 背景装饰圆 -->
    <div class="bg-circle c1" />
    <div class="bg-circle c2" />
    <div class="bg-circle c3" />

    <main class="center">
      <!-- Logo -->
      <div class="logo-wrap">
        <h1 class="logo">
          <span class="s">S</span><span class="e1">e</span><span class="a">a</span><span class="r">r</span><span class="c">c</span><span class="h">h</span><span class="go">Go</span>
        </h1>
        <p class="tagline">探索知识，触手可及</p>
      </div>

      <!-- 搜索框 -->
      <div class="search-area">
        <SearchBar v-model="query" @search="handleSearch" size="large" />

        <div class="btn-row">
          <button class="btn-search" @click="handleSearch">SearchGo 搜索</button>
          <button class="btn-lucky" @click="handleLucky">手气不错</button>
        </div>
      </div>

      <!-- 热门搜索 -->
      <div class="hot-tags">
        <span class="hot-label">热门：</span>
        <span
          v-for="tag in hotTags"
          :key="tag"
          class="hot-tag"
          @click="quickSearch(tag)"
        >{{ tag }}</span>
      </div>
    </main>

    <!-- 底部 -->
    <footer class="footer">
      <span>© 2025 SearchGo</span>
      <span>·</span>
      <span>Powered by Go + Vue3</span>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import SearchBar from '@/components/SearchBar.vue'

const router = useRouter()
const query = ref('')

const hotTags = ['Golang', 'Vue3', '搜索引擎', 'MySQL', 'Redis', '微服务', 'Docker']

function handleSearch() {
  const q = query.value.trim()
  if (!q) return
  router.push({ name: 'search', query: { q, page: 1 } })
}

function handleLucky() {
  const tag = hotTags[Math.floor(Math.random() * hotTags.length)]
  router.push({ name: 'search', query: { q: tag, page: 1 } })
}

function quickSearch(tag) {
  query.value = tag
  handleSearch()
}
</script>

<style scoped>
.home {
  min-height: 100vh;
  background: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  overflow: hidden;
}

/* 背景装饰 */
.bg-circle {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.12;
  pointer-events: none;
  animation: float 8s ease-in-out infinite;
}
.c1 { width: 520px; height: 520px; background: #4285f4; top: -120px; left: -160px; animation-delay: 0s; }
.c2 { width: 420px; height: 420px; background: #ea4335; bottom: -100px; right: -100px; animation-delay: 3s; }
.c3 { width: 300px; height: 300px; background: #34a853; top: 40%; left: 60%; animation-delay: 6s; }

@keyframes float {
  0%, 100% { transform: translateY(0px) scale(1); }
  50%       { transform: translateY(-30px) scale(1.05); }
}

/* 中央内容 */
.center {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 24px;
  gap: 36px;
  width: 100%;
  max-width: 720px;
}

/* Logo */
.logo-wrap {
  text-align: center;
}
.logo {
  font-size: clamp(3rem, 10vw, 5.5rem);
  font-weight: 800;
  letter-spacing: -2px;
  line-height: 1;
  display: inline-flex;
  gap: 0;
}
/* 逐字着色，模仿 Google 配色 */
.s  { color: #4285f4; }
.e1 { color: #ea4335; }
.a  { color: #fbbc05; }
.r  { color: #4285f4; }
.c  { color: #34a853; }
.h  { color: #ea4335; }
.go { color: #1a73e8; }

.tagline {
  margin-top: 10px;
  color: #9aa0a6;
  font-size: 15px;
  letter-spacing: 1px;
}

/* 搜索区域 */
.search-area {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

/* 按钮行 */
.btn-row {
  display: flex;
  gap: 12px;
}
.btn-search,
.btn-lucky {
  padding: 10px 24px;
  border-radius: var(--radius-sm);
  font-size: 14px;
  font-weight: 500;
  background: #f8f9fa;
  color: var(--text);
  border: 1px solid var(--border);
  transition: var(--transition);
}
.btn-search:hover,
.btn-lucky:hover {
  background: #e8eaed;
  border-color: #c0c2c4;
  box-shadow: var(--shadow-sm);
}

/* 热门标签 */
.hot-tags {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  justify-content: center;
}
.hot-label {
  color: var(--text-muted);
  font-size: 13px;
}
.hot-tag {
  padding: 5px 14px;
  border-radius: var(--radius-full);
  border: 1px solid var(--border);
  font-size: 13px;
  color: var(--text-muted);
  cursor: pointer;
  transition: var(--transition);
}
.hot-tag:hover {
  border-color: var(--primary);
  color: var(--primary);
  background: #e8f0fe;
}

/* 底部 */
.footer {
  width: 100%;
  padding: 16px;
  display: flex;
  justify-content: center;
  gap: 12px;
  color: var(--text-light);
  font-size: 13px;
  border-top: 1px solid #f1f3f4;
}

/* =============================================
   响应式断点
   ============================================= */

/* 平板 ≤ 768px */
@media (max-width: 768px) {
  .center {
    padding: 32px 20px;
    gap: 28px;
  }
  .c1 { width: 360px; height: 360px; }
  .c2 { width: 300px; height: 300px; }
  .c3 { width: 220px; height: 220px; }
}

/* 手机 ≤ 480px */
@media (max-width: 480px) {
  .center {
    padding: 24px 16px;
    gap: 22px;
  }
  .tagline {
    font-size: 12px;
    letter-spacing: 0.5px;
  }
  /* 按钮在手机上垂直排列，宽度撑满 */
  .btn-row {
    flex-direction: column;
    width: 100%;
    align-items: center;
  }
  .btn-search,
  .btn-lucky {
    width: min(280px, 100%);
    text-align: center;
    justify-content: center;
    padding: 11px 16px;
  }
  /* 背景装饰在小屏缩小，避免撑破布局 */
  .c1 { width: 240px; height: 240px; top: -60px; left: -80px; }
  .c2 { width: 200px; height: 200px; bottom: -60px; right: -60px; }
  .c3 { width: 160px; height: 160px; }
  /* 热门标签间距收紧 */
  .hot-tags { gap: 6px; }
  .hot-tag  { padding: 4px 10px; font-size: 12px; }
  .footer   { font-size: 12px; gap: 8px; padding: 12px; }
}

/* 超小屏 ≤ 360px */
@media (max-width: 360px) {
  .center { padding: 16px 12px; gap: 18px; }
  .c1, .c2, .c3 { display: none; } /* 极小屏隐藏装饰圆，减少渲染压力 */
}

/* 大屏 ≥ 1440px：内容区可以更宽 */
@media (min-width: 1440px) {
  .center { max-width: 860px; }
}

/* 超宽屏 ≥ 1920px */
@media (min-width: 1920px) {
  .center { max-width: 960px; }
  .c1 { width: 700px; height: 700px; }
  .c2 { width: 560px; height: 560px; }
}
</style>
