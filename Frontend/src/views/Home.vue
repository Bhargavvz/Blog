<script setup lang="ts">
import PageTransition from '../components/PageTransition.vue';
import BlogPost from '../components/BlogPost.vue';

// Sample featured posts data
const featuredPosts = [
  {
    title: 'Getting Started with Vue 3 Composition API',
    excerpt: 'Learn how to use Vue 3\'s Composition API to write more maintainable and reusable code.',
    image: 'https://picsum.photos/800/600?random=1',
    date: 'March 15, 2024',
    slug: 'getting-started-vue-3',
  },
  {
    title: 'Building Modern UIs with Tailwind CSS',
    excerpt: 'Discover how to create beautiful user interfaces using Tailwind CSS utility classes.',
    image: 'https://picsum.photos/800/600?random=2',
    date: 'March 10, 2024',
    slug: 'modern-uis-tailwind',
  },
  {
    title: 'TypeScript Best Practices',
    excerpt: 'Essential TypeScript patterns and practices for building robust applications.',
    image: 'https://picsum.photos/800/600?random=3',
    date: 'March 5, 2024',
    slug: 'typescript-best-practices',
  },
];
</script>

<template>
  <PageTransition>
    <div>
      <!-- Hero Section -->
      <section class="relative bg-gray-50 dark:bg-gray-800 py-16 sm:py-24">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div class="text-center">
            <h1 class="text-4xl tracking-tight font-extrabold text-gray-900 dark:text-white sm:text-5xl md:text-6xl">
              Welcome to Modern Blog
            </h1>
            <p class="mt-3 max-w-md mx-auto text-base text-gray-500 dark:text-gray-300 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl">
              Exploring the latest in web development, design patterns, and software engineering best practices.
            </p>
          </div>
        </div>
      </section>

      <!-- Featured Posts -->
      <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <h2 class="text-3xl font-bold text-gray-900 dark:text-white mb-8">Featured Posts</h2>
        <div class="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
          <BlogPost
            v-for="post in featuredPosts"
            :key="post.slug"
            v-bind="post"
          />
        </div>
      </section>
    </div>
  </PageTransition>
</template>