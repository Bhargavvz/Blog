<script setup lang="ts">
import PageTransition from '../components/PageTransition.vue';
</script>

<template>
  <PageTransition>
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-8">About</h1>
      <div class="prose dark:prose-invert max-w-none">
        <p class="text-lg text-gray-600 dark:text-gray-300 mb-6">
          Welcome to Modern Blog, where we explore the latest trends and best practices in web development.
          Our mission is to provide high-quality, practical content that helps developers build better applications.
        </p>
        <h2 class="text-2xl font-semibold text-gray-900 dark:text-white mt-8 mb-4">Our Mission</h2>
        <p class="text-gray-600 dark:text-gray-300 mb-6">
          We believe in sharing knowledge and helping developers stay up-to-date with the rapidly evolving
          web development landscape. Through tutorials, articles, and practical examples, we aim to make
          complex concepts accessible to developers of all skill levels.
        </p>
      </div>
    </div>
  </PageTransition>
</template>