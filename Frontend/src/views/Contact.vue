<script setup lang="ts">
import { ref } from 'vue'
import { sendContactMessage } from '../services/api'

const name = ref('')
const email = ref('')
const subject = ref('')
const message = ref('')
const error = ref('')
const success = ref('')
const loading = ref(false)

const handleSubmit = async () => {
  try {
    loading.value = true
    error.value = ''
    success.value = ''

    const contactData = {
      name: name.value,
      email: email.value,
      subject: subject.value,
      message: message.value
    }

    console.log('Sending contact data:', contactData)

    const result = await sendContactMessage(contactData)

    console.log('Contact response:', result)
    success.value = result.message || 'Thank you for your message. We\'ll get back to you soon!'
    
    // Clear form
    name.value = ''
    email.value = ''
    subject.value = ''
    message.value = ''
  } catch (err) {
    console.error('Form submission error:', err)
    error.value = err.response?.data?.error || 'Error sending message. Please try again.'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-8">
      Contact Us
    </h1>

    <div class="bg-white dark:bg-gray-800 shadow sm:rounded-lg p-6">
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div v-if="error" class="text-red-500 text-sm">
          {{ error }}
        </div>
        <div v-if="success" class="text-green-500 text-sm">
          {{ success }}
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Name
          </label>
          <input
            v-model="name"
            type="text"
            required
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
          >
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Email
          </label>
          <input
            v-model="email"
            type="email"
            required
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
          >
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Subject
          </label>
          <input
            v-model="subject"
            type="text"
            required
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
          >
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Message
          </label>
          <textarea
            v-model="message"
            rows="6"
            required
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
          />
        </div>

        <div>
          <button
            type="submit"
            :disabled="loading"
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 disabled:opacity-50"
          >
            {{ loading ? 'Sending...' : 'Send Message' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>