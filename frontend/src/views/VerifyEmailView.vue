<!-- src/views/VerifyEmailView.vue -->
<template>
    <div class="verify-email">
        <h1>Verify Email</h1>
        <p>Please check your email for a verification link.</p>
    </div>
</template>

<script setup>
import { onMounted } from 'vue'
import axios from 'axios'

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

<style scoped>
.verify-email {
    padding: 20px;
    max-width: 400px;
    margin: 0 auto;
}

h1 {
    font-size: 2em;
    margin-bottom: 20px;
}
</style>