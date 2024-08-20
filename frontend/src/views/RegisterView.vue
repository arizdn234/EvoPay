<template>
    <div style="min-height: 100vh; display: grid; place-items: center;">
        <div class="register bg-white dark:bg-gray-900 shadow-lg rounded-lg p-8 mx-auto">
            <h1 class="text-3xl font-semibold mb-2 text-center text-gray-800 dark:text-gray-200">Register</h1>
            <p class="mb-4 px-2 mx-auto text-center text-grey3-dark dark:text-grey3-light max-w-96">Welcome to <span
                    class="font-bold">EvoPay</span>! Create your account and join us today.
                Fill in your details below to get started.</p>

            <div v-if="error"
                class="error-message rounded-lg bg-red-light dark:text-red-dark text-white dark:text-white px-8 py-4 text-center mx-auto my-4 max-w-80">
                {{ error }}
            </div>

            <form @submit.prevent="handleSubmit" class="space-y-6">
                <div class="relative space-y-2">
                    <label for="email" class="block text-gray-700 dark:text-gray-300 text-md font-regular">Email</label>
                    <div class="relative">
                        <i class='bx bx-envelope bx-sm absolute pt-3 pl-3 text-blue-light dark:text-blue-dark'></i>
                        <input type="email" v-model="email" name="email" id="email" autocomplete="email" required
                            spellcheck="false"
                            class="w-full py-3 rounded-lg pl-12 pr-4 font-thin bg-input placeholder-gray-700 dark:placeholder-gray-400"
                            placeholder="Enter your email" />
                    </div>
                </div>

                <div class="relative space-y-2">
                    <label for="password"
                        class="block text-gray-700 dark:text-gray-300 text-md font-regular">Password</label>
                    <div class="relative">
                        <i class='bx bx-lock-alt bx-sm absolute pt-3 pl-3 text-blue-light dark:text-blue-dark'></i>
                        <input type="password" v-model="password" name="password" id="password"
                            autocomplete="current-password" required spellcheck="false"
                            class="w-full py-3 rounded-lg pl-12 pr-4 font-thin bg-input placeholder-gray-700 dark:placeholder-gray-400"
                            placeholder="Enter your password" />
                    </div>
                </div>

                <div class="relative space-y-2">
                    <label for="name" class="block text-gray-700 dark:text-gray-300 text-md font-regular">Name</label>
                    <div class="relative">
                        <i class='bx bx-user bx-sm absolute pt-3 pl-3 text-blue-light dark:text-blue-dark'></i>
                        <input type="text" v-model="name" name="name" id="name" autocomplete="username" required
                            spellcheck="false"
                            class="w-full py-3 rounded-lg pl-12 pr-4 font-thin bg-input placeholder-gray-700 dark:placeholder-gray-400"
                            placeholder="Enter your name" />
                    </div>
                </div>

                <div class="flex items-center justify-center">
                    <button type="submit"
                        class="bg-blue-light dark:bg-blue-dark hover:filter hover:brightness-125 text-white text-center font-semibold py-2 px-4 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:ring-opacity-50 flex w-full flex items-center justify-center mb-3">
                        Register
                    </button>
                </div>
                <span>Already have an account?<a href="/users/login"
                        class="text-blue-light dark:text-blue-dark hover:filter hover:brightness-125 font-regular">
                        Login here
                    </a></span>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from '../axios'
import { useRouter } from 'vue-router'

const router = useRouter()
const email = ref('')
const password = ref('')
const name = ref('')
const error = ref('')

const handleSubmit = async () => {
    error.value = ''
    try {
        const response = await axios.post('/users/register', {
            email: email.value,
            password: password.value,
            name: name.value
        })
        console.log('Registration successful')
        setTimeout(() => {
            router.push('/users/verify-email')
        }, 1000);
    } catch (err) {
        if (err.response) {
            error.value = err.response.data.error || 'An error occurred during registration.'
            console.error('Error response data:', err.response.data)
        } else if (err.request) {
            error.value = 'No response received from server.'
            console.error('Error request data:', err.request)
        } else {
            error.value = 'An error occurred while setting up the request.'
            console.error('Error message:', err.message)
        }
        console.error('Error config:', err.config)
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
