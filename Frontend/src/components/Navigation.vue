<script setup lang="ts">
import { ref } from 'vue';
import ThemeToggle from './ThemeToggle.vue';

const isMenuOpen = ref(false);
// Temporarily hardcode isAdmin for testing
const isAdmin = ref(true); // Set this to false to hide the admin link
</script>

<template>
  <nav class="fixed w-full bg-white/80 dark:bg-gray-900/80 backdrop-blur-sm z-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex">
          <div class="flex-shrink-0 flex items-center">
            <router-link to="/" class="text-xl font-bold text-primary-600">
              Modern Blog
            </router-link>
          </div>
        </div>
        
        <div class="hidden sm:ml-6 sm:flex sm:items-center sm:space-x-8">
          <router-link
            to="/"
            class="text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            active-class="text-primary-600 dark:text-primary-400"
          >
            Home
          </router-link>
          <router-link
            to="/blog"
            class="text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            active-class="text-primary-600 dark:text-primary-400"
          >
            Blog
          </router-link>
          <router-link
            to="/about"
            class="text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            active-class="text-primary-600 dark:text-primary-400"
          >
            About
          </router-link>
          <router-link
            to="/contact"
            class="text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            active-class="text-primary-600 dark:text-primary-400"
          >
            Contact
          </router-link>
          <router-link
            v-if="isAdmin"
            to="/admin"
            class="text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
            active-class="text-primary-600 dark:text-primary-400"
          >
            Admin
          </router-link>
          <ThemeToggle />
        </div>

        <div class="flex items-center sm:hidden">
          <button
            @click="isMenuOpen = !isMenuOpen"
            class="inline-flex items-center justify-center p-2 rounded-md text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
          >
            <span class="sr-only">Open main menu</span>
            <svg
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                v-if="!isMenuOpen"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 6h16M4 12h16M4 18h16"
              />
              <path
                v-else
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <div v-show="isMenuOpen" class="sm:hidden">
      <div class="pt-2 pb-3 space-y-1">
        <router-link
          to="/"
          class="block px-3 py-2 text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
          active-class="text-primary-600 dark:text-primary-400"
          @click="isMenuOpen = false"
        >
          Home
        </router-link>
        <router-link
          to="/blog"
          class="block px-3 py-2 text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
          active-class="text-primary-600 dark:text-primary-400"
          @click="isMenuOpen = false"
        >
          Blog
        </router-link>
        <router-link
          to="/about"
          class="block px-3 py-2 text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
          active-class="text-primary-600 dark:text-primary-400"
          @click="isMenuOpen = false"
        >
          About
        </router-link>
        <router-link
          to="/contact"
          class="block px-3 py-2 text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
          active-class="text-primary-600 dark:text-primary-400"
          @click="isMenuOpen = false"
        >
          Contact
        </router-link>
        <router-link
          v-if="isAdmin"
          to="/admin"
          class="block px-3 py-2 text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400"
          active-class="text-primary-600 dark:text-primary-400"
          @click="isMenuOpen = false"
        >
          Admin
        </router-link>
        <div class="px-3 py-2">
          <ThemeToggle />
        </div>
      </div>
    </div>
  </nav>
</template>