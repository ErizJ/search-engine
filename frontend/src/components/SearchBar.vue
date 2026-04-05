<template>
  <div class="search-bar-wrap" :class="size">
    <!-- 搜索图标 -->
    <svg class="icon-search" viewBox="0 0 24 24" fill="none" stroke="currentColor"
         stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="11" cy="11" r="8" />
      <line x1="21" y1="21" x2="16.65" y2="16.65" />
    </svg>

    <!-- 输入框 -->
    <input
      ref="inputRef"
      v-model="innerValue"
      type="text"
      :placeholder="placeholder"
      class="search-input"
      @keyup.enter="onSearch"
      @input="onInput"
      autocomplete="off"
      spellcheck="false"
    />

    <!-- 清除按钮 -->
    <button v-if="innerValue" class="btn-clear" @click="clear" title="清除">
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor"
           stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <line x1="18" y1="6" x2="6" y2="18" />
        <line x1="6" y1="6" x2="18" y2="18" />
      </svg>
    </button>

    <!-- 分割线 -->
    <div class="divider" />

    <!-- 搜索按钮 -->
    <button class="btn-go" @click="onSearch" title="搜索">
      <svg viewBox="0 0 24 24" fill="currentColor">
        <path d="M15.5 14h-.79l-.28-.27A6.471 6.471 0 0 0 16 9.5 6.5 6.5 0 1 0 9.5 16c1.61
                 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99
                 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/>
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: { type: String, default: '' },
  size: { type: String, default: 'medium' },       // 'large' | 'medium'
  placeholder: { type: String, default: '搜索...' },
})

const emit = defineEmits(['update:modelValue', 'search'])

const inputRef = ref(null)
const innerValue = ref(props.modelValue)

watch(() => props.modelValue, (v) => { innerValue.value = v })

function onInput() {
  emit('update:modelValue', innerValue.value)
}

function onSearch() {
  emit('update:modelValue', innerValue.value)
  emit('search')
}

function clear() {
  innerValue.value = ''
  emit('update:modelValue', '')
  inputRef.value?.focus()
}

// 暴露 focus 方法供父组件调用
defineExpose({ focus: () => inputRef.value?.focus() })
</script>

<style scoped>
.search-bar-wrap {
  display: flex;
  align-items: center;
  width: 100%;
  background: #fff;
  border: 1.5px solid var(--border);
  border-radius: var(--radius-full);
  transition: var(--transition);
  position: relative;
}
.search-bar-wrap:hover,
.search-bar-wrap:focus-within {
  border-color: transparent;
  box-shadow: var(--shadow-md);
}

/* 尺寸变体 */
.search-bar-wrap.large {
  padding: 14px 8px 14px 20px;
  gap: 10px;
}
.search-bar-wrap.medium {
  padding: 9px 6px 9px 16px;
  gap: 8px;
}

/* 搜索图标 */
.icon-search {
  flex-shrink: 0;
  color: var(--text-muted);
}
.large  .icon-search { width: 20px; height: 20px; }
.medium .icon-search { width: 16px; height: 16px; }

/* 输入框 */
.search-input {
  flex: 1;
  background: transparent;
  font-family: inherit;
  color: var(--text);
  min-width: 0;
}
.large  .search-input { font-size: 18px; }
.medium .search-input { font-size: 15px; }

.search-input::placeholder { color: var(--text-light); }

/* 清除按钮 */
.btn-clear {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px; height: 28px;
  border-radius: 50%;
  background: transparent;
  color: var(--text-muted);
  transition: var(--transition);
}
.btn-clear svg { width: 16px; height: 16px; }
.btn-clear:hover { background: #f1f3f4; color: var(--text); }

/* 分割线 */
.divider {
  width: 1px;
  background: var(--border);
  flex-shrink: 0;
}
.large  .divider { height: 24px; }
.medium .divider { height: 20px; }

/* 搜索按钮 */
.btn-go {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  color: var(--primary);
  transition: var(--transition);
  border-radius: 50%;
}
.large  .btn-go { width: 44px; height: 44px; }
.medium .btn-go { width: 34px; height: 34px; }
.large  .btn-go svg { width: 22px; height: 22px; }
.medium .btn-go svg { width: 18px; height: 18px; }
.btn-go:hover { background: #e8f0fe; }

/* =============================================
   响应式断点
   ============================================= */

/* 手机 ≤ 640px
   重要：iOS Safari 对 font-size < 16px 的 input 会自动缩放页面
   必须强制设为 16px 防止这个问题 */
@media (max-width: 640px) {
  .large .search-input,
  .medium .search-input {
    font-size: 16px;
  }
  /* large 在移动端略微收窄，贴合小屏 */
  .search-bar-wrap.large {
    padding: 11px 6px 11px 14px;
    gap: 8px;
  }
  .large .icon-search { width: 18px; height: 18px; }
  .large .btn-go      { width: 38px; height: 38px; }
  .large .btn-go svg  { width: 19px; height: 19px; }
  .large .divider     { height: 20px; }
}

/* 超小屏 ≤ 360px */
@media (max-width: 360px) {
  .search-bar-wrap.large {
    padding: 10px 4px 10px 12px;
  }
  /* 隐藏搜索图标节省空间 */
  .large .icon-search { display: none; }
}

/* 大屏 ≥ 1440px：输入框字体可以稍大 */
@media (min-width: 1440px) {
  .large  .search-input { font-size: 20px; }
  .search-bar-wrap.large {
    padding: 16px 10px 16px 24px;
  }
  .large .btn-go     { width: 48px; height: 48px; }
  .large .btn-go svg { width: 24px; height: 24px; }
}
</style>
