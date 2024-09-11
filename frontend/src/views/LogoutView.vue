<!-- src/views/LogoutView.vue -->
<template>
    <div class="logout">
        <h1>Logout</h1>
        <button @click="handleLogout">Logout</button>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const handleLogout = async () => {
    try {
        await axios.get('/users/logout')
        console.log('Logged out successfully')
    } catch (error) {
        console.error('Error logging out:', error)
    }
}
</script>

<style scoped>
.logout {
    padding: 20px;
    max-width: 400px;
    margin: 0 auto;
}

button {
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    cursor: pointer;
}
</style>