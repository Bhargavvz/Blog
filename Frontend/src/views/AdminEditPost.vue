<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { createPost, updatePost, getPost } from '../services/api'

const route = useRoute()
const router = useRouter()
const isNewPost = !route.params.id

const post = ref({
  title: '',
  content: '',
  slug: '',
  published: false
})

const error = ref('')
const loading = ref(false)

onMounted(async () => {
  if (!isNewPost) {
    try {
      loading.value = true
      const data = await getPost(route.params.id as string)
      post.value = data
    } catch (err) {
      error.value = 'Error loading post'
      console.error(err)
    } finally {
      loading.value = false
    }
  }
})

const handleSubmit = async () => {
  try {
    loading.value = true
    error.value = ''

    if (isNewPost) {
      await createPost(post.value)
    } else {
      await updatePost(route.params.id as string, post.value)
    }

    router.push('/admin')
  } catch (err) {
    error.value = 'Error saving post'
    console.error(err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
        {{ isNewPost ? 'Create New Post' : 'Edit Post' }}
      </h1>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">
      <div v-if="error" class="text-red-500">{{ error }}</div>

      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
          Title
        </label>
        <input
          v-model="post.title"
          type="text"
          required
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
        >
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
          Slug
        </label>
        <input
          v-model="post.slug"
          type="text"
          required
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
        >
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
          Content
        </label>
        <textarea
          v-model="post.content"
          rows="10"
          required
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
        />
      </div>

      <div class="flex items-center">
        <input
          v-model="post.published"
          type="checkbox"
          class="h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded"
        >
        <label class="ml-2 block text-sm text-gray-900 dark:text-gray-300">
          Published
        </label>
      </div>

      <div class="flex justify-end space-x-4">
        <button
          type="button"
          @click="router.push('/admin')"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 dark:bg-gray-700 dark:text-gray-300 dark:border-gray-600 dark:hover:bg-gray-600"
        >
          Cancel
        </button>
        <button
          type="submit"
          :disabled="loading"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 disabled:opacity-50"
        >
          {{ loading ? 'Saving...' : 'Save Post' }}
        </button>
      </div>
    </form>
  </div>
</template> 