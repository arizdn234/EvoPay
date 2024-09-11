<!-- src/views/ResetPasswordView.vue -->
<template>
    <div class="reset-password">
        <h1>Reset Password</h1>
        <form @submit.prevent="handleSubmit">
            <label for="new-password">New Password:</label>
            <input type="password" v-model="newPassword" required />

            <label for="confirm-password">Confirm Password:</label>
            <input type="password" v-model="confirmPassword" required />

            <button type="submit">Reset Password</button>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

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
.reset-password {
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