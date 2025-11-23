<template>
  <div class="page-container">
    <slot />
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.page-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--space-16);
}
</style>