<script setup lang="ts">
import Navigation from './components/Navigation.vue';
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <Navigation />
    <main class="pt-16">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
    <footer class="mt-16 py-8 bg-gray-50 dark:bg-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center text-gray-500 dark:text-gray-400">
          <p>&copy; {{ new Date().getFullYear() }} Modern Blog. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>