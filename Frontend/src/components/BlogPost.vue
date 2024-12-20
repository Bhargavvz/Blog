<script setup lang="ts">
defineProps<{
  title: string;
  excerpt: string;
  image: string;
  date: string;
  slug: string;
}>();
</script>

<template>
  <article class="group">
    <router-link :to="`/blog/${slug}`" class="block">
      <div class="aspect-w-16 aspect-h-9 mb-4 overflow-hidden rounded-lg">
        <img
          :src="image"
          :alt="title"
          class="object-cover w-full h-full transform transition-transform duration-300 group-hover:scale-105"
        />
      </div>
      <div class="space-y-2">
        <time class="text-sm text-gray-500 dark:text-gray-400">{{ date }}</time>
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white group-hover:text-primary-600 dark:group-hover:text-primary-400">
          {{ title }}
        </h3>
        <p class="text-gray-600 dark:text-gray-300">
          {{ excerpt }}
        </p>
      </div>
    </router-link>
  </article>
</template>