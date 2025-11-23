<template>
  <div class="auth-layout">
    <div class="auth-center">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.auth-layout {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--space-16);
  background: linear-gradient(180deg, #fbfcfe 0%, #f2f5fa 100%);
}
.auth-center {
  width: 100%;
}
</style>