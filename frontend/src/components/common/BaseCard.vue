<template>
  <el-card class="base-card" :shadow="shadow">
    <slot />
  </el-card>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{ shadow?: 'never' | 'always' | 'hover' }>();
</script>

<style scoped>
.base-card {
  border-radius: var(--radius-l);
  box-shadow: var(--shadow-m);
}
</style>