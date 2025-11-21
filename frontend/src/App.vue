<template>
  <div id="app">
    <!-- RouterView 用于渲染当前路由匹配到的组件 -->
    <RouterView />

    <!-- 导入并使用 NoticeCircle 组件 -->
    <NoticeCircle />
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router';
import NoticeCircle from '@/components/NoticeCircle.vue';
</script>

<style>
/* global.css 或 style.css */
:root {
  --primary-color: #409eff;
  --bg-gradient: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  --card-bg: rgba(255, 255, 255, 0.9);
  --glass-border: 1px solid rgba(255, 255, 255, 0.6);
  --shadow-sm: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  --shadow-lg: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

body {
  margin: 0;
  background: #f0f2f5; /* 兜底色 */
  font-family: 'Inter', 'Helvetica Neue', Helvetica, 'PingFang SC', 'Microsoft YaHei', sans-serif;
  -webkit-font-smoothing: antialiased;
}

/* 优化 Element Plus 默认圆角 */
.el-input__wrapper, .el-button, .el-card {
  border-radius: 12px !important;
}

/* 滚动条美化 */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
::-webkit-scrollbar-thumb {
  background: #c0c4cc;
  border-radius: 4px;
}
::-webkit-scrollbar-track {
  background: transparent;
}
</style>
