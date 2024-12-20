<script setup lang="ts">
// Wrapper component for consistent page transitions
</script>

<template>
  <transition
    name="page"
    mode="out-in"
    appear
  >
    <slot />
  </transition>
</template>

<style scoped>
.page-enter-active,
.page-leave-active {
  transition: all 0.3s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(20px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}
</style>