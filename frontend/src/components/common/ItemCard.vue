<template>
  <div class="item-card" @click="$emit('click')">
    <div class="left">
      <div v-if="time" class="time">{{ time }}</div>
      <div v-for="(line, idx) in lines" :key="idx" class="row">
        <span class="label" v-if="line.label">{{ line.label }}：</span>
        <span :class="{ highlight: idx === highlightIndex }">{{ line.value }}</span>
      </div>
    </div>
    <div class="right">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{ time?: string; lines: Array<{ label?: string; value: string }>; highlightIndex?: number }>();
</script>

<style scoped>
.item-card {
  background: #fff;
  border-radius: var(--radius-l);
  padding: var(--space-16);
  border: 1px solid rgba(0,0,0,0.03);
  box-shadow: var(--shadow-s);
  display: flex;
  flex-direction: column;
  transition: transform 160ms ease, box-shadow 160ms ease;
}

.item-card:hover { transform: translateY(-2px); box-shadow: var(--shadow-m); }

@media (min-width: 640px) { .item-card { flex-direction: row; align-items: center; justify-content: space-between; } }

.left { flex: 1; padding-right: var(--space-16); }
.right { flex-shrink: 0; display: flex; gap: var(--space-12); }

.time {
  font-size: var(--font-size-s);
  color: var(--color-muted);
  background: #f4f4f5;
  padding: 4px 8px;
  border-radius: 4px;
  display: inline-block;
  margin-bottom: var(--space-12);
}

.row { font-size: 15px; color: var(--color-text); line-height: 1.6; }
.label { color: #606266; font-weight: 500; }
.highlight { color: var(--el-color-primary); }
</style>