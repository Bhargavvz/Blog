<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { getPost } from '../services/api';

const route = useRoute();
const post = ref(null);
const loading = ref(true);
const error = ref('');

onMounted(async () => {
  try {
    post.value = await getPost(route.params.slug as string);
  } catch (err) {
    error.value = 'Error loading post';
  } finally {
    loading.value = false;
  }
});
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <div v-if="loading" class="text-center">
      Loading...
    </div>
    
    <div v-else-if="error" class="text-red-500 text-center">
      {{ error }}
    </div>
    
    <article v-else class="prose dark:prose-invert lg:prose-lg mx-auto">
      <img
        v-if="post.image"
        :src="`data:${post.image.mimeType};base64,${post.image.data}`"
        :alt="post.title"
        class="w-full h-64 object-cover rounded-lg mb-8"
      >
      
      <h1>{{ post.title }}</h1>
      
      <div class="flex items-center gap-4 text-gray-600 dark:text-gray-400 mb-8">
        <span>{{ post.author }}</span>
        <span>{{ new Date(post.createdAt).toLocaleDateString() }}</span>
        <span>{{ post.readTime }} min read</span>
      </div>
      
      <div class="flex gap-2 mb-8">
        <span
          v-for="tag in post.tags"
          :key="tag"
          class="px-3 py-1 bg-gray-200 dark:bg-gray-700 rounded-full text-sm"
        >
          {{ tag }}
        </span>
      </div>
      
      <div v-html="post.content" />
    </article>
  </div>
</template>