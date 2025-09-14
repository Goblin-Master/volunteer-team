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
/* 全局重置样式 */
body {
  margin: 0;
  padding: 0;
  font-family: sans-serif; /* 添加一个基础字体 */
}
</style>
