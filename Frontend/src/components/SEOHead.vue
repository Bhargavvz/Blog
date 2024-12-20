<script setup lang="ts">
import { computed } from 'vue';
import { useHead } from '@vueuse/head';

const props = defineProps<{
  title: string;
  description: string;
  keywords?: string[];
  image?: string;
}>();

const metaTags = computed(() => [
  {
    name: 'description',
    content: props.description
  },
  {
    name: 'keywords',
    content: props.keywords?.join(', ') || ''
  },
  {
    property: 'og:title',
    content: props.title
  },
  {
    property: 'og:description',
    content: props.description
  },
  {
    property: 'og:image',
    content: props.image
  }
]);

useHead({
  title: computed(() => `${props.title} | Modern Blog`),
  meta: metaTags
});
</script>

<template>
  <!-- Component is renderless -->
</template>