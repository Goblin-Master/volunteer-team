<template>
  <div v-if="isLogin && show" class="global-logout">
    <el-button size="small" plain type="danger" @click="onLogout">退出登录</el-button>
  </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import { computed } from 'vue';
import { useUserStore } from '@/stores/user';

const router = useRouter();
const route = useRoute();
const userStore = useUserStore();
const isLogin = computed(() => userStore.isUserLogin);
const show = computed(() => (route.name as string) !== 'Home');
const onLogout = () => {
  userStore.logout();
  router.replace('/login');
};
</script>

<style scoped>
.global-logout {
  position: fixed;
  top: var(--space-16);
  right: var(--space-16);
  z-index: 2200;
  box-shadow: var(--shadow-s);
}
</style>