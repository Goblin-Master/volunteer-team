<template>
  <el-button
    :type="type"
    :size="size"
    :plain="plain"
    :round="round"
    :loading="loading"
    class="base-button"
  >
    <slot />
  </el-button>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{ type?: string; size?: 'small' | 'default' | 'large'; plain?: boolean; round?: boolean; loading?: boolean }>();
</script>

<style scoped>
.base-button {
  border-radius: var(--radius-m);
  transition: transform 160ms ease, box-shadow 160ms ease;
}
.base-button:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-s);
}
</style>