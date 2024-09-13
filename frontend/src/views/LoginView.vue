<template>
    <div class="login bg-white dark:bg-gray-900 shadow-lg rounded-lg p-8 max-w-md mx-auto mt-28">
        <h1 class="text-3xl font-semibold mb-8 text-center text-gray-800 dark:text-gray-200">Login</h1>
        <form @submit.prevent="handleSubmit" class="space-y-6">
            <div class="relative space-y-2">
                <label for="email" class="block text-gray-700 dark:text-gray-300 text-sm font-medium">Email</label>
                <div class="relative">
                    <input type="email" v-model="email" required
                        class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg dark:bg-gray-800 shadow-sm focus:border-blue-500 dark:focus:border-blue-400 focus:ring-1 focus:ring-blue-500 dark:focus:ring-blue-400 text-gray-900 dark:text-gray-100 placeholder-gray-500 dark:placeholder-gray-400 pl-10"
                        placeholder="Enter your email" />
                    <i
                        class="fas fa-envelope absolute top-1/2 left-3 transform -translate-y-1/2 text-gray-500 dark:text-gray-400"></i>
                </div>
            </div>

            <div class="relative space-y-2">
                <label for="password"
                    class="block text-gray-700 dark:text-gray-300 text-sm font-medium">Password</label>
                <div class="relative">
                    <input type="password" v-model="password" required
                        class="mb-2 w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg dark:bg-gray-800 shadow-sm focus:border-blue-500 dark:focus:border-blue-400 focus:ring-1 focus:ring-blue-500 dark:focus:ring-blue-400 text-gray-900 dark:text-gray-100 placeholder-gray-500 dark:placeholder-gray-400 pl-10"
                        placeholder="Enter your password" />
                    <i
                        class="fa-solid fa-lock absolute top-1/2 left-3 transform -translate-y-1/2 text-gray-500 dark:text-gray-400"></i>
                </div>
            </div>
            <a href="/users/reset-password"
                class="mt-4 text-indigo-700 dark:text-indigo-600 hover:text-indigo-400 dark:hover:text-indigo-400 font-semibold rounded-lg transition duration-150 ease-in-out text-right">
                Forgot Password?
            </a>

            <div class="flex items-center justify-center">
                <button type="submit"
                    class="bg-indigo-700 dark:bg-indigo-600 hover:bg-indigo-400 dark:hover:bg-indigo-400 text-white text-center font-semibold py-2 px-4 rounded-lg transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:ring-opacity-50 flex w-full flex items-center justify-center">
                    Login
                </button>
            </div>

            <div class="flex items-center justify-between">
                <span>Don't have an account?<a href="/users/register"
                        class="text-indigo-700 dark:text-indigo-600 hover:text-indigo-400 dark:hover:text-indigo-400 font-semibold rounded-lg transition duration-150 ease-in-out">
                        Register here
                    </a></span>
            </div>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from '../axios'

const email = ref('')
const password = ref('')

const handleSubmit = async () => {
    try {
        const response = await axios.post('/users/login', {
            email: email.value,
            password: password.value
        })
        console.log('Login successful:', response.data)
    } catch (error) {
        console.error('Error logging in:', error)
    }
}
</script>

<style scoped>
.relative {
    position: relative;
}

input {
    transition: border-color 0.2s ease, background-color 0.2s ease;
}
</style>
