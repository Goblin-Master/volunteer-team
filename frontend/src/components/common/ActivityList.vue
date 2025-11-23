<template>
  <div class="activity">
    <SectionTitle>{{ title }}</SectionTitle>
    <div class="list">
      <ItemCard
        v-for="(it, i) in items"
        :key="i"
        :time="it.time"
        :lines="it.lines"
        :highlightIndex="it.highlightIndex ?? 0"
      >
        <slot name="actions" :item="it" />
      </ItemCard>
    </div>
  </div>
</template>

<script setup lang="ts">
import SectionTitle from '@/components/common/SectionTitle.vue';
import ItemCard from '@/components/common/ItemCard.vue';

defineProps<{ title: string; items: Array<{ time?: string; lines: Array<{ label?: string; value: string }>; highlightIndex?: number }> }>();
</script>

<style scoped>
.list { display: flex; flex-direction: column; gap: var(--space-16); }
</style>