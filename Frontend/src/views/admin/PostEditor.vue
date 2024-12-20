<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import PageTransition from '../../components/PageTransition.vue';
import type { Post } from '../../types';

const route = useRoute();
const router = useRouter();
const isEditing = computed(() => route.params.id !== 'new');

const post = ref<Partial<Post>>({
  title: '',
  excerpt: '',
  content: '',
  image: '',
  tags: [],
  seo: {
    title: '',
    description: '',
    keywords: []
  }
});

const savePost = async () => {
  // Implement save logic
  router.push('/admin');
};
</script>

<template>
  <PageTransition>
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <h1 class="text-3xl font-bold mb-8">
        {{ isEditing ? 'Edit Post' : 'New Post' }}
      </h1>
      
      <form @submit.prevent="savePost" class="space-y-6">
        <!-- Basic Info -->
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1">Title</label>
            <input
              v-model="post.title"
              type="text"
              class="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium mb-1">Excerpt</label>
            <textarea
              v-model="post.excerpt"
              rows="3"
              class="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
              required
            ></textarea>
          </div>
          
          <div>
            <label class="block text-sm font-medium mb-1">Content</label>
            <textarea
              v-model="post.content"
              rows="10"
              class="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
              required
            ></textarea>
          </div>
        </div>

        <!-- SEO Section -->
        <div class="border-t pt-6 space-y-4">
          <h2 class="text-xl font-semibold">SEO Settings</h2>
          <div>
            <label class="block text-sm font-medium mb-1">SEO Title</label>
            <input
              v-model="post.seo.title"
              type="text"
              class="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium mb-1">SEO Description</label>
            <textarea
              v-model="post.seo.description"
              rows="2"
              class="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
            ></textarea>
          </div>
        </div>

        <div class="flex justify-end space-x-4">
          <button
            type="button"
            @click="router.push('/admin')"
            class="px-4 py-2 border border-gray-300 rounded-md hover:bg-gray-50"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700"
          >
            Save Post
          </button>
        </div>
      </form>
    </div>
  </PageTransition>
</template>