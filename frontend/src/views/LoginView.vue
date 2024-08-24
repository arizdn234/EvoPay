<template>
    <div style="min-height: 100vh; display: grid; place-items: center;">
        <div class="login bg-white dark:bg-gray-900 shadow-lg rounded-lg p-8 mx-auto">
            <h1 class="text-3xl font-semibold mb-2 text-center text-gray-800 dark:text-gray-200">Login</h1>
            <p class="mb-4 px-16 text-center text-grey3-dark dark:text-grey3-light">Welcome to <span
                    class="font-bold">EvoPay</span>, login to
                continue</p>

            <div v-if="error"
                class="error-message rounded-lg bg-red-light dark:text-red-dark text-white px-8 py-4 text-center mx-auto my-4 max-w-80">
                {{ error }}
            </div>

            <form @submit.prevent="handleSubmit" class="space-y-6">
                <div class="relative space-y-2">
                    <label for="email" class="block text-gray-700 dark:text-gray-300 text-md font-regular">Email</label>
                    <div class="relative">
                        <i class='bx bx-user bx-sm absolute pt-3 pl-3 text-blue-light dark:text-blue-dark'></i>
                        <input type="email" name="email" id="email" v-model="email" autocomplete="email" required
                            class="w-full px-5 py-3 rounded-lg pl-10 font-thin bg-input placeholder-gray-700 dark:placeholder-gray-400"
                            placeholder="Enter your email" />

                    </div>
                </div>

                <div class="relative space-y-2 mb-12">
                    <label for="password"
                        class="block text-gray-700 dark:text-gray-300 text-md font-regular">Password</label>
                    <div class="relative">
                        <i class='bx bx-lock-alt bx-sm absolute pt-3 pl-3 text-blue-light dark:text-blue-dark'></i>

                        <input type="password" name="password" id="password" v-model="password" required
                            class="w-full px-5 py-3 rounded-lg pl-10 font-thin bg-input placeholder-gray-700 dark:placeholder-gray-400"
                            placeholder="Enter your password" />
                    </div>
                </div>
                
                <a href="/users/reset-password"
                    class="mt-4 text-blue-light dark:text-blue-dark hover:filter hover:brightness-125 font-regular rounded-lg transition duration-150 ease-in-out text-right">
                    Forgot Password?
                </a>

                <div class="flex items-center justify-center">
                    <button type="submit"
                        class="bg-blue-light dark:bg-blue-dark hover:filter hover:brightness-125 text-white text-center font-semibold py-2 px-4 rounded-lg transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:ring-opacity-50 flex w-full flex items-center justify-center">
                        Login
                    </button>
                </div>

                <div class="flex items-center justify-between">
                    <span>Don't have an account?<a href="/users/register"
                            class="text-blue-light dark:text-blue-dark hover:filter hover:brightness-125 font-semibold rounded-lg transition duration-150 ease-in-out">
                            Register here
                        </a></span>
                </div>
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
const error = ref('')

const handleSubmit = async () => {
    error.value = ''
    try {
        const response = await axios.post('/users/login', {
            email: email.value,
            password: password.value
        }, { withCredentials: true })

        console.log('Login successful', response.data)
        setTimeout(() => {
            router.push('/users/homepage')
            setTimeout(() => {
                window.location.reload();
            }, 100);
        }, 500);

    } catch (err) {

        if (err.response) {
            error.value = err.response.data.error || 'An error occurred during login.'
            console.error('Error response data', err.response.data)
        } else if (err.request) {
            error.value = 'No response received from the server.'
            console.error('Error request data:', err.request)
        } else {
            error.value = 'An error occurred while setting up the request.'
            console.error('Error message:', err.message)
        }
        console.error('Error logging in:', err)
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
