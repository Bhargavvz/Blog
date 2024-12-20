<script setup lang="ts">
import { useRouter } from 'vue-router';
import { logout } from '../services/api';

const router = useRouter();

const handleLogout = () => {
  logout();
  router.push('/login');
};
</script>

<template>
  <nav class="bg-white shadow dark:bg-gray-800">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex">
          <div class="flex-shrink-0 flex items-center">
            <router-link to="/" class="text-xl font-bold text-gray-800 dark:text-white">
              Blog
            </router-link>
          </div>
          <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
            <router-link
              to="/"
              class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900 dark:text-gray-100"
            >
              Home
            </router-link>
            <router-link
              to="/admin"
              class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900 dark:text-gray-100"
            >
              Admin
            </router-link>
            <router-link
              to="/contact"
              class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900 dark:text-gray-100"
            >
              Contact
            </router-link>
          </div>
        </div>
        <div class="flex items-center">
          <button
            @click="handleLogout"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  </nav>
</template> 