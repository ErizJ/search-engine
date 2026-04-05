<template>
  <article class="result-item" @click="openLink">
    <!-- 来源路径 -->
    <div class="result-meta">
      <div class="favicon">
        <svg viewBox="0 0 24 24" fill="currentColor">
          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12
                   2zm-2 14.5v-9l6 4.5-6 4.5z"/>
        </svg>
      </div>
      <div class="meta-text">
        <span class="result-site">{{ siteName }}</span>
        <span class="result-breadcrumb">{{ displayUrl }}</span>
      </div>
      <!-- 更多操作（装饰性） -->
      <button class="btn-more" @click.stop title="更多">
        <svg viewBox="0 0 24 24" fill="currentColor">
          <circle cx="12" cy="5" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="12" cy="19" r="1.5"/>
        </svg>
      </button>
    </div>

    <!-- 标题 -->
    <h3 class="result-title">
      <!-- eslint-disable-next-line vue/no-v-html -->
      <span v-html="item.title" />
    </h3>

    <!-- 摘要 -->
    <!-- eslint-disable-next-line vue/no-v-html -->
    <p class="result-snippet" v-html="item.snippet" />

    <!-- 评分（调试用，默认隐藏） -->
    <div v-if="showScore" class="result-score">
      TF-IDF: {{ item.score.toFixed(4) }}
    </div>
  </article>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  item: { type: Object, required: true },
  query: { type: String, default: '' },
  showScore: { type: Boolean, default: false },
})

// 从 URL 提取站点名
const siteName = computed(() => {
  try {
    const url = new URL(props.item.url || 'https://example.com')
    return url.hostname.replace('www.', '')
  } catch {
    return 'example.com'
  }
})

// 美化 URL 显示（截断过长路径）
const displayUrl = computed(() => {
  const url = props.item.url || ''
  if (!url) return ''
  const clean = url.replace(/^https?:\/\/(www\.)?/, '')
  return clean.length > 60 ? clean.slice(0, 60) + '…' : clean
})

function openLink() {
  if (props.item.url) {
    window.open(props.item.url, '_blank', 'noopener,noreferrer')
  }
}
</script>

<style scoped>
.result-item {
  padding: 18px 0 18px;
  border-bottom: 1px solid #f1f3f4;
  cursor: pointer;
  transition: background var(--transition);
  border-radius: var(--radius-md);
  position: relative;
}
.result-item:hover {
  background: #fafafa;
  padding-left: 12px;
  padding-right: 12px;
  margin-left: -12px;
  margin-right: -12px;
}
.result-item:last-child { border-bottom: none; }

/* Meta 行 */
.result-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 6px;
}
.favicon {
  width: 20px; height: 20px;
  border-radius: 50%;
  background: #e8f0fe;
  display: flex; align-items: center; justify-content: center;
  flex-shrink: 0;
  color: var(--primary);
}
.favicon svg { width: 12px; height: 12px; }
.meta-text {
  display: flex;
  flex-direction: column;
  min-width: 0;
}
.result-site {
  font-size: 13px;
  font-weight: 500;
  color: var(--text);
  line-height: 1.4;
}
.result-breadcrumb {
  font-size: 12px;
  color: var(--text-muted);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.btn-more {
  margin-left: auto;
  width: 28px; height: 28px;
  border-radius: 50%;
  background: transparent;
  color: var(--text-muted);
  display: flex; align-items: center; justify-content: center;
  opacity: 0;
  transition: var(--transition);
}
.btn-more svg { width: 16px; height: 16px; }
.result-item:hover .btn-more { opacity: 1; }
.btn-more:hover { background: #e8eaed; }

/* 标题 */
.result-title {
  font-size: 20px;
  font-weight: 400;
  color: #1a0dab;
  line-height: 1.35;
  margin-bottom: 6px;
  transition: var(--transition);
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.result-item:hover .result-title { color: #1558d6; text-decoration: underline; }

/* 标题高亮 */
:deep(.result-title em) {
  font-style: normal;
  color: #1558d6;
  font-weight: 600;
}

/* 摘要 */
.result-snippet {
  font-size: 14px;
  line-height: 1.65;
  color: #4d5156;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* 摘要高亮 */
:deep(.result-snippet em) {
  font-style: normal;
  font-weight: 700;
  color: #333;
}

/* 评分 */
.result-score {
  margin-top: 6px;
  font-size: 11px;
  color: #aaa;
}

/* =============================================
   响应式断点
   ============================================= */

/* 触摸设备：禁用 hover 偏移效果（负 margin 在移动端会造成布局抖动），
   改用 :active 提供触摸反馈 */
@media (hover: none) {
  .result-item:hover {
    background: transparent;
    padding-left: 0;
    padding-right: 0;
    margin-left: 0;
    margin-right: 0;
  }
  .result-item:hover .result-title {
    text-decoration: none;
    color: #1a0dab;
  }
  .result-item:hover .btn-more { opacity: 0; }
  .result-item:active { background: #f5f7fa; }
}

/* 手机 ≤ 640px */
@media (max-width: 640px) {
  .result-item { padding: 14px 0; }
  .result-title {
    font-size: 17px;
    -webkit-line-clamp: 2;
  }
  .result-snippet {
    font-size: 13px;
    line-height: 1.6;
    -webkit-line-clamp: 4; /* 手机上显示多一行 */
  }
  .result-breadcrumb { max-width: 56vw; }
  /* 隐藏更多按钮（触摸设备上没有 hover，无意义） */
  .btn-more { display: none; }
}

/* 超小屏 ≤ 360px */
@media (max-width: 360px) {
  .result-title   { font-size: 15px; }
  .result-snippet { font-size: 12px; }
  .result-site    { font-size: 12px; }
}

/* 大屏 ≥ 1280px：标题字体放大 */
@media (min-width: 1280px) {
  .result-title   { font-size: 22px; }
  .result-snippet { font-size: 15px; }
}
</style>
