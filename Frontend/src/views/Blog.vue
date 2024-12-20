<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getPosts } from '../services/api'

const router = useRouter()
const posts = ref([])
const loading = ref(true)
const error = ref('')

onMounted(async () => {
  try {
    posts.value = await getPosts()
  } catch (err) {
    error.value = 'Error loading posts'
  } finally {
    loading.value = false
  }
})

const goToPost = (slug: string) => {
  router.push(`/blog/${slug}`)
}
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-8">
      Blog Posts
    </h1>

    <div v-if="loading" class="text-center">
      Loading...
    </div>

    <div v-else-if="error" class="text-red-500 text-center">
      {{ error }}
    </div>

    <div v-else class="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
      <article
        v-for="post in posts"
        :key="post.id"
        class="bg-white dark:bg-gray-800 rounded-lg shadow-md overflow-hidden cursor-pointer"
        @click="goToPost(post.slug)"
      >
        <img
          v-if="post.image"
          :src="`data:${post.image.mimeType};base64,${post.image.data}`"
          :alt="post.title"
          class="w-full h-48 object-cover"
        >
        
        <div class="p-6">
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-2">
            {{ post.title }}
          </h2>
          
          <p class="text-gray-600 dark:text-gray-400 mb-4">
            {{ post.summary }}
          </p>
          
          <div class="flex flex-wrap gap-2 mb-4">
            <span
              v-for="tag in post.tags"
              :key="tag"
              class="px-2 py-1 bg-gray-200 dark:bg-gray-700 rounded-full text-sm"
            >
              {{ tag }}
            </span>
          </div>
          
          <div class="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
            <span>{{ post.author }}</span>
            <span>{{ post.readTime }} min read</span>
          </div>
        </div>
      </article>
    </div>
  </div>
</template>