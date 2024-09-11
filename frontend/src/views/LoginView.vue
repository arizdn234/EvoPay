<template>
    <div class="login">
        <h1>Login</h1>
        <form @submit.prevent="handleSubmit">
            <label for="email">Email:</label>
            <input type="email" v-model="email" required />

            <label for="password">Password:</label>
            <input type="password" v-model="password" required />

            <button type="submit">Login</button>
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
.login {
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
</style>