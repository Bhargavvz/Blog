<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const posts = ref([])
const auth = useAuthStore()
const router = useRouter()

onMounted(async () => {
  // Mock data - replace with actual API call
  posts.value = [
    { id: 1, title: 'First Post', status: 'published' },
    { id: 2, title: 'Second Post', status: 'draft' },
  ]
})

const handleNewPost = () => {
  router.push('/admin/edit')
}

const handleEditPost = (id: number) => {
  router.push(`/admin/edit/${id}`)
}

const handleDeletePost = async (id: number) => {
  // Add delete logic here
  posts.value = posts.value.filter(post => post.id !== id)
}
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
        Admin Dashboard
      </h1>
      <button
        @click="handleNewPost"
        class="bg-primary-600 text-white px-4 py-2 rounded-md hover:bg-primary-700"
      >
        New Post
      </button>
    </div>

    <div class="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-md">
      <ul class="divide-y divide-gray-200 dark:divide-gray-700">
        <li v-for="post in posts" :key="post.id" class="px-6 py-4">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">
                {{ post.title }}
              </h3>
              <p class="text-sm text-gray-500 dark:text-gray-400">
                Status: {{ post.status }}
              </p>
            </div>
            <div class="flex space-x-3">
              <button
                @click="handleEditPost(post.id)"
                class="text-primary-600 hover:text-primary-700"
              >
                Edit
              </button>
              <button
                @click="handleDeletePost(post.id)"
                class="text-red-600 hover:text-red-700"
              >
                Delete
              </button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template> 