<template>
  <el-select
    v-model="names"
    multiple
    filterable
    allow-create
    default-first-option
    :reserve-keyword="false"
    class="receivers-input w-full"
    :placeholder="placeholder || '输入姓名，回车添加'"
  />
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{ modelValue: string; placeholder?: string }>();
const emit = defineEmits<{ (e: 'update:modelValue', v: string): void }>();

const parse = (v: string) => v ? v.split(/[,，]/).map(s => s.trim()).filter(Boolean) : [];
const join = (arr: string[]) => arr.join(', ');

const names = ref<string[]>(parse(props.modelValue));

watch(() => props.modelValue, v => { names.value = parse(v); });
watch(names, v => { emit('update:modelValue', join(v)); }, { deep: true });
</script>

<style scoped>
.receivers-input :deep(.el-select__tags) { gap: 6px; flex-wrap: wrap; }
.receivers-input :deep(.el-tag) { border-radius: var(--radius-s); }
</style>