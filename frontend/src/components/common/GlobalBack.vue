<template>
  <div v-if="show" class="global-back">
    <el-button circle class="back-btn" @click="goBack">
      <el-icon><ArrowLeft /></el-icon>
    </el-button>
  </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import { computed } from 'vue';
import { ArrowLeft } from '@element-plus/icons-vue';

const router = useRouter();
const route = useRoute();

const hideOn = new Set(['Home']);
const show = computed(() => !hideOn.has(route.name as string));

const goBack = () => {
  router.back();
};
</script>

<style scoped>
.global-back {
  position: fixed;
  top: var(--space-16);
  left: var(--space-16);
  z-index: 2200;
  box-shadow: var(--shadow-s);
}
.back-btn {
  background: #f0f2f5;
  border: none;
  color: #606266;
}
.back-btn:hover { filter: brightness(1.03); }
</style>