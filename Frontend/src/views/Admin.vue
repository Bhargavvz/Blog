<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import { getPosts, deletePost } from '../services/api';
import type { Post } from '../types';

const router = useRouter();
const auth = useAuthStore();
const posts = ref<Post[]>([]);
const error = ref('');
const loading = ref(true);

onMounted(async () => {
  try {
    const data = await getPosts();
    posts.value = data || [];
  } catch (err) {
    error.value = 'Error loading posts';
    console.error(err);
    posts.value = [];
  } finally {
    loading.value = false;
  }
});

const handleEdit = (id: string) => {
  router.push(`/admin/posts/${id}/edit`);
};

const handleDelete = async (id: string) => {
  if (!confirm('Are you sure you want to delete this post?')) return;
  
  try {
    await deletePost(id);
    posts.value = posts.value.filter(post => post.id !== id);
  } catch (err) {
    error.value = 'Error deleting post';
    console.error(err);
  }
};

const handleLogout = () => {
  auth.logout();
  router.push('/admin/login');
};
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
        Admin Dashboard
      </h1>
      <div class="flex space-x-4">
        <router-link
          to="/admin/posts/new"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          New Post
        </router-link>
        <button
          @click="handleLogout"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 dark:bg-red-700 dark:hover:bg-red-600"
        >
          Logout
        </button>
      </div>
    </div>

    <div v-if="error" class="mb-4 text-red-600">
      {{ error }}
    </div>

    <div v-if="loading" class="text-center py-8">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
    </div>

    <div v-else-if="posts" class="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-md">
      <ul v-if="posts.length > 0" class="divide-y divide-gray-200 dark:divide-gray-700">
        <li v-for="post in posts" :key="post.id" class="px-6 py-4">
          <div class="flex items-center justify-between">
            <div class="flex-1 min-w-0">
              <div class="flex items-center">
                <p class="text-sm font-medium text-primary-600 truncate">
                  {{ post.title }}
                </p>
                <span
                  v-if="!post.published"
                  class="ml-2 px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800"
                >
                  Draft
                </span>
              </div>
              <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">
                {{ new Date(post.createdAt).toLocaleDateString() }}
              </p>
            </div>
            <div class="flex space-x-4">
              <button
                @click="handleEdit(post.id)"
                class="text-primary-600 hover:text-primary-900 dark:hover:text-primary-400"
              >
                Edit
              </button>
              <button
                @click="handleDelete(post.id)"
                class="text-red-600 hover:text-red-900 dark:hover:text-red-400"
              >
                Delete
              </button>
            </div>
          </div>
        </li>
      </ul>
      <div v-else class="px-6 py-4 text-center text-gray-500 dark:text-gray-400">
        No posts yet. Create your first post!
      </div>
    </div>
  </div>
</template>

<style scoped>
.dark .bg-red-600 {
  background-color: rgb(220, 38, 38);
}

.dark .hover\:bg-red-700:hover {
  background-color: rgb(185, 28, 28);
}
</style> 