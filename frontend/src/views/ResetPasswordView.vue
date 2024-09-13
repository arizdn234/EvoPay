<!-- src/views/ResetPasswordView.vue -->
<template>
    <div class="reset-password bg-white dark:bg-gray-900 shadow-lg rounded-lg p-8 max-w-md mx-auto mt-28">
        <h1 class="text-3xl font-semibold mb-8 text-center text-gray-800 dark:text-gray-200">Reset Password</h1>
        <form @submit.prevent="handleSubmit" class="space-y-6">
            <div class="relative space-y-2">
                <label for="new-password" class="block text-gray-700 dark:text-gray-300 text-sm font-medium">New
                    Password:</label>
                <div class="relative">
                    <input type="password" v-model="newPassword" required
                        class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg dark:bg-gray-800 shadow-sm focus:border-blue-500 dark:focus:border-blue-400 focus:ring-1 focus:ring-blue-500 dark:focus:ring-blue-400 text-gray-900 dark:text-gray-100 placeholder-gray-500 dark:placeholder-gray-400"
                        placeholder="Enter new password" />
                </div>
            </div>

            <div class="relative space-y-2">
                <label for="confirm-password" class="block text-gray-700 dark:text-gray-300 text-sm font-medium">Confirm
                    Password:</label>
                <div class="relative">
                    <input type="password" v-model="confirmPassword" required
                        class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg dark:bg-gray-800 shadow-sm focus:border-blue-500 dark:focus:border-blue-400 focus:ring-1 focus:ring-blue-500 dark:focus:ring-blue-400 text-gray-900 dark:text-gray-100 placeholder-gray-500 dark:placeholder-gray-400"
                        placeholder="Confirm your password" />
                </div>
            </div>

            <div class="flex items-center justify-center">
                <button type="submit"
                    class="mb-3 w-full flex justify-center bg-indigo-700 dark:bg-indigo-600 hover:bg-indigo-400 dark:hover:bg-indigo-400 text-white font-semibold py-2 px-4 rounded-lg transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:ring-opacity-50 flex items-center">
                    Reset Password
                </button>
            </div>
            <a href="/users/login"
                class="text-indigo-700 dark:text-indigo-600 hover:text-indigo-400 dark:hover:text-indigo-400 font-semibold py-2 px-4 rounded-lg transition duration-150 ease-in-out">
                Back to Login page
            </a>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from '../axios'

const newPassword = ref('')
const confirmPassword = ref('')

const handleSubmit = async () => {
    if (newPassword.value !== confirmPassword.value) {
        console.error('Passwords do not match')
        return
    }

    try {
        const params = new URLSearchParams(window.location.search)
        const token = params.get('token')

        if (token) {
            await axios.post('/users/reset-password', {
                token,
                newPassword: newPassword.value
            })
            console.log('Password reset successfully')
        }
    } catch (error) {
        console.error('Error resetting password:', error)
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