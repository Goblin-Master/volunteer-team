<template>
  <h2 class="section-title">
    <slot />
  </h2>
</template>

<script setup lang="ts"></script>

<style scoped>
.section-title {
  margin: 0 0 var(--space-16);
  font-size: 20px;
  font-weight: 700;
  color: var(--color-text);
  letter-spacing: -0.2px;
}
</style>