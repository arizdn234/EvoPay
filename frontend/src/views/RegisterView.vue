<template>
    <div class="register">
        <h1>Register</h1>
        <form @submit.prevent="handleSubmit">
            <label for="email">Email:</label>
            <input type="email" v-model="email" required />

            <label for="password">Password:</label>
            <input type="password" v-model="password" required />

            <label for="name">Name:</label>
            <input type="text" v-model="name" required />

            <button type="submit">Register</button>

            <div v-if="error" class="error-message">
                {{ error }}
            </div>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from '../axios'

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
        console.log('Registration successful:', response.data)
        // Redirect to login page
        router.push('/login')
    } catch (err) {
        if (err.response) {
            error.value = err.response.data.message || 'An error occurred during registration.'
            console.error('Error response data:', err.response.data)
            console.error('Error response status:', err.response.status)
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
/* Existing styles */
.register {
    padding: 20px;
    max-width: 400px;
    margin: 0 auto;
}

h1 {
    font-size: 2em;
    margin-bottom: 20px;
}

form {
    display: flex;
    flex-direction: column;
}

label {
    margin-top: 10px;
}

input {
    padding: 8px;
    margin-top: 5px;
}

button {
    margin-top: 20px;
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    cursor: pointer;
}

/* New style for error message */
.error-message {
    margin-top: 20px;
    padding: 10px;
    background-color: #f8d7da;
    color: #721c24;
    border: 1px solid #f5c6cb;
    border-radius: 4px;
}
</style>