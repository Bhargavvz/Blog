<script setup lang="ts">
import { ref } from 'vue';
import PageTransition from '../../components/PageTransition.vue';
import { useAuth } from '../../composables/useAuth';
import type { Post } from '../../types';

const { isAdmin } = useAuth();
const posts = ref<Post[]>([]);

// Add sample data
posts.value = [
  {
    id: '1',
    title: 'Sample Post',
    excerpt: 'Sample excerpt',
    content: 'Sample content',
    image: 'https://picsum.photos/800/600',
    date: new Date().toISOString(),
    slug: 'sample-post',
    author: 'Admin',
    tags: ['vue', 'javascript'],
    seo: {
      title: 'Sample Post',
      description: 'Sample description',
      keywords: ['vue', 'javascript']
    }
  }
];
</script>

<template>
  <PageTransition>
    <div v-if="isAdmin" class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="flex justify-between items-center mb-8">
        <h1 class="text-3xl font-bold">Admin Dashboard</h1>
        <router-link
          to="/admin/posts/new"
          class="bg-primary-600 text-white px-4 py-2 rounded-md hover:bg-primary-700 transition-colors"
        >
          New Post
        </router-link>
      </div>
      
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md">
        <div class="p-6">
          <h2 class="text-xl font-semibold mb-4">Recent Posts</h2>
          <div class="space-y-4">
            <div v-for="post in posts" :key="post.id" class="flex justify-between items-center p-4 bg-gray-50 dark:bg-gray-700 rounded-md">
              <div>
                <h3 class="font-medium">{{ post.title }}</h3>
                <p class="text-sm text-gray-500 dark:text-gray-400">{{ new Date(post.date).toLocaleDateString() }}</p>
              </div>
              <div class="space-x-2">
                <router-link
                  :to="`/admin/posts/${post.id}/edit`"
                  class="text-primary-600 hover:text-primary-700"
                >
                  Edit
                </router-link>
                <button
                  class="text-red-600 hover:text-red-700"
                  @click="() => {}"
                >
                  Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="text-center py-16">
      <p>Access denied. Please log in as an administrator.</p>
    </div>
  </PageTransition>
</template>