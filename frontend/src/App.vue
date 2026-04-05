<template>
  <router-view />

  <!-- 模式指示徽标：固定在右下角，开发时随时知道当前跑哪个模式 -->
  <div class="mode-badge" :class="isMock ? 'mock' : 'real'">
    <span class="mode-dot" />
    {{ isMock ? 'Mock 模式' : 'Real 后端' }}
  </div>
</template>

<script setup>
import { isMock } from '@/api/search.js'
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body, #app {
  height: 100%;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC',
               'Hiragino Sans GB', 'Microsoft YaHei', sans-serif;
  -webkit-font-smoothing: antialiased;
}
</style>

<style scoped>
.mode-badge {
  position: fixed;
  bottom: 20px;
  right: 20px;
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 5px 12px;
  border-radius: 100px;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.3px;
  pointer-events: none;
  z-index: 9999;
  backdrop-filter: blur(6px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  user-select: none;
}

/* Mock 模式：橙色 */
.mode-badge.mock {
  background: rgba(255, 107, 53, 0.9);
  color: #fff;
}

/* 真实后端：绿色 */
.mode-badge.real {
  background: rgba(52, 168, 83, 0.9);
  color: #fff;
}

.mode-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.8);
  animation: pulse 2s ease-in-out infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1;   transform: scale(1); }
  50%       { opacity: 0.5; transform: scale(0.8); }
}
</style>
