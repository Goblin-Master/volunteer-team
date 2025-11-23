<template>
  <div class="stat-card">
    <div class="icon"><slot name="icon" /></div>
    <div class="content">
      <div class="value">{{ value }}</div>
      <div class="label">{{ label }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{ value: number | string; label: string }>();
</script>

<style scoped>
.stat-card {
  display: flex;
  align-items: center;
  gap: var(--space-12);
  background: #fff;
  border-radius: var(--radius-l);
  box-shadow: var(--shadow-s);
  padding: var(--space-16);
}
.icon { width: 40px; height: 40px; border-radius: 10px; display: flex; align-items: center; justify-content: center; font-size: 22px; background: #ecf2ff; color: var(--el-color-primary); }
.content { display: flex; flex-direction: column; }
.value { font-size: 20px; font-weight: 700; color: var(--color-text); }
.label { font-size: 13px; color: var(--color-muted); }
</style>