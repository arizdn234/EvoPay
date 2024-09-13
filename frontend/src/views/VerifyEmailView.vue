<!-- src/views/VerifyEmailView.vue -->
<template>
    <div class="verify-email bg-white dark:bg-gray-900 shadow-lg rounded-lg p-8 max-w-md mx-auto mt-28">
      <h1 class="text-3xl font-semibold mb-4 text-center text-gray-800 dark:text-gray-200">Verify Email</h1>
      <p class="text-center text-gray-700 dark:text-gray-300">Please check your email for a verification link.</p>
    </div>
  </template>
  
  <script setup>
  import { onMounted } from 'vue'
  import axios from '../axios'
  
  onMounted(async () => {
    const params = new URLSearchParams(window.location.search)
    const token = params.get('token')
  
    if (token) {
      try {
        await axios.get(`/users/verify-email?token=${token}`)
        console.log('Email verified successfully')
      } catch (error) {
        console.error('Error verifying email:', error)
      }
    }
  })
  </script>
  