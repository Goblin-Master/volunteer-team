<template>
  <el-config-provider :size="'default'" :z-index="2000">
    <component :is="layoutComponent">
      <RouterView />
    </component>
    <NoticeCircle />
    <GlobalLogout />
    <GlobalBack />
  </el-config-provider>
</template>

<script setup>
import { RouterView, useRoute } from 'vue-router';
import { computed } from 'vue';
import NoticeCircle from '@/components/NoticeCircle.vue';
import AuthLayout from '@/layouts/AuthLayout.vue';
import GlobalLogout from '@/components/common/GlobalLogout.vue';
import GlobalBack from '@/components/common/GlobalBack.vue';

const route = useRoute();
const layoutComponent = computed(() => (route.meta && route.meta.layout === 'auth' ? AuthLayout : 'div'));
</script>

<style>
::-webkit-scrollbar { width: 8px; height: 8px; }
::-webkit-scrollbar-thumb { background: #c0c4cc; border-radius: 4px; }
::-webkit-scrollbar-track { background: transparent; }
</style>
