<template>
  <nav class="pagination" aria-label="分页">
    <!-- 上一页 -->
    <button
      class="page-btn arrow"
      :disabled="current <= 1"
      @click="go(current - 1)"
      title="上一页"
    >
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
           stroke-linecap="round" stroke-linejoin="round">
        <polyline points="15 18 9 12 15 6" />
      </svg>
    </button>

    <!-- 页码 -->
    <template v-for="p in pages" :key="p">
      <!-- 省略号 -->
      <span v-if="p === '...'" class="page-ellipsis">…</span>
      <!-- 普通页码 -->
      <button
        v-else
        class="page-btn"
        :class="{ active: p === current }"
        @click="go(p)"
      >{{ p }}</button>
    </template>

    <!-- 下一页 -->
    <button
      class="page-btn arrow"
      :disabled="current >= total"
      @click="go(current + 1)"
      title="下一页"
    >
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
           stroke-linecap="round" stroke-linejoin="round">
        <polyline points="9 18 15 12 9 6" />
      </svg>
    </button>
  </nav>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  current: { type: Number, required: true },
  total:   { type: Number, required: true },
})

const emit = defineEmits(['change'])

function go(page) {
  if (page < 1 || page > props.total || page === props.current) return
  emit('change', page)
}

// 生成智能页码序列：当总页数多时使用省略号
// 例：1 2 3 ... 8 9 10  或  1 ... 4 5 6 ... 10
const pages = computed(() => {
  const total = props.total
  const cur = props.current
  if (total <= 7) {
    return Array.from({ length: total }, (_, i) => i + 1)
  }
  const result = []
  const addPage = (p) => result.push(p)
  const addDot = () => {
    if (result[result.length - 1] !== '...') result.push('...')
  }

  addPage(1)
  if (cur > 3) addDot()
  for (let p = Math.max(2, cur - 1); p <= Math.min(total - 1, cur + 1); p++) {
    addPage(p)
  }
  if (cur < total - 2) addDot()
  addPage(total)

  return result
})
</script>

<style scoped>
.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  padding: 32px 0 16px;
}

.page-btn {
  min-width: 36px;
  height: 36px;
  border-radius: var(--radius-sm);
  background: transparent;
  color: var(--primary);
  font-size: 14px;
  font-weight: 500;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: var(--transition);
  padding: 0 6px;
}
.page-btn:hover:not(:disabled):not(.active) {
  background: #e8f0fe;
}
.page-btn.active {
  background: var(--primary);
  color: #fff;
  cursor: default;
  border-radius: 50%;
  min-width: 36px; width: 36px;
}
.page-btn:disabled {
  color: var(--border);
  cursor: not-allowed;
}
.page-btn.arrow svg {
  width: 16px; height: 16px;
}

.page-ellipsis {
  color: var(--text-muted);
  padding: 0 4px;
  font-size: 16px;
  line-height: 36px;
}

/* =============================================
   响应式断点
   ============================================= */

/* 手机 ≤ 480px：缩小按钮，确保不溢出 */
@media (max-width: 480px) {
  .pagination {
    gap: 2px;
    padding: 20px 0 12px;
    flex-wrap: wrap;
    justify-content: center;
  }
  .page-btn {
    min-width: 34px;
    height: 34px;
    font-size: 13px;
    padding: 0 4px;
  }
  .page-btn.active {
    min-width: 34px;
    width: 34px;
  }
  .page-ellipsis {
    font-size: 14px;
    line-height: 34px;
    padding: 0 2px;
  }
}

/* 超小屏 ≤ 360px：进一步收紧 */
@media (max-width: 360px) {
  .page-btn {
    min-width: 30px;
    height: 30px;
    font-size: 12px;
  }
  .page-btn.active { min-width: 30px; width: 30px; }
}

/* 大屏 ≥ 1280px：更舒适的间距 */
@media (min-width: 1280px) {
  .pagination { gap: 6px; padding: 40px 0 20px; }
  .page-btn   { min-width: 40px; height: 40px; font-size: 15px; }
  .page-btn.active { min-width: 40px; width: 40px; }
}
</style>
