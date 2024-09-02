<template>
	<div :class="{ 'dark': isDarkMode }">
		<header
			class="font-sans fixed px-28 top-0 left-0 right-0 z-50 p-4 bg-gray-800 dark:bg-gray-900 flex justify-between items-center shadow-md">
			<div class="container mx-auto flex justify-between items-center">
				<div class="flex space-x-4">
					<RouterLink to="/"
						class="text-gray-300 dark:text-gray-200 hover:text-white dark:hover:text-gray-100"
						active-class="text-indigo-400 dark:text-indigo-300" exact-active-class="font-bold">
						API Docs
					</RouterLink>
				</div>

				<div class="flex items-center space-x-4">
					<!-- Show Login and Register only if the user is not logged in -->
					<RouterLink to="/users/login"
						class="w-16 text-gray-300 dark:text-gray-200 hover:text-white dark:hover:text-gray-100"
						active-class="text-indigo-400 dark:text-indigo-300" exact-active-class="font-bold"
						v-if="!isLoggedIn">
						Login
					</RouterLink>
					<RouterLink to="/users/register"
						class="w-16 text-gray-300 dark:text-gray-200 hover:text-white dark:hover:text-gray-100"
						active-class="text-indigo-400 dark:text-indigo-300" exact-active-class="font-bold"
						v-if="!isLoggedIn">
						Register
					</RouterLink>

					<!-- Dark Mode Toggle -->
					<button @click="toggleDarkMode"
						class="text-white dark:text-gray-200 p-2 rounded-lg hover:bg-gray-700 dark:hover:bg-gray-600 focus:border-gray-500 dark:focus:border-gray-400 focus:ring-1 focus:ring-gray-500 dark:focus:ring-gray-400">
						<i :class="isDarkMode ? 'fas fa-sun' : 'fas fa-moon'" class="icon-size"></i>
					</button>

					<!-- Show Logout only if the user is logged in -->
					<button @click="handleLogout"
						class="text-white py-2 px-4 rounded-lg bg-rose-600 hover:bg-rose-500 dark:hover:bg-rose-500 focus:border-rose-500 dark:focus:border-rose-400 focus:ring-1 focus:ring-rose-500 dark:focus:ring-rose-400"
						v-if="isLoggedIn">
						Logout
					</button>
				</div>
			</div>
		</header>
		<main class="font-sans p-4 bg-gray-200 dark:bg-gray-800 text-gray-900 dark:text-gray-100 mt-16 min-h-screen">
			<!-- Your content here -->
			<RouterView />
		</main>
	</div>
</template>

<script setup>
import { ref, watchEffect } from 'vue';
import Cookies from 'js-cookie';
import { RouterLink, RouterView } from 'vue-router';
import { useRouter } from 'vue-router';
import axios from './axios';

// Router instance
const router = useRouter();

// Theme and login state
const isDarkMode = ref(false);
const isLoggedIn = ref(false);

// Dark Mode Toggle
const toggleDarkMode = () => {
	isDarkMode.value = !isDarkMode.value;
	document.documentElement.classList.toggle('dark', isDarkMode.value);
	localStorage.setItem('theme', isDarkMode.value ? 'dark' : 'light');
};

// Handle Logout Function
const handleLogout = async () => {
	try {
		await axios.post('users/logout', {}, { withCredentials: true });
		console.log('Logout successful');
		Cookies.remove('auth_token');
		isLoggedIn.value = false;
		router.push('/users/login');
	} catch (error) {
		console.error('Logout error:', error);
	}
};

// Watch for changes in the auth_token cookie and update the login state
watchEffect(() => {
	const token = Cookies.get('auth_token');
	isLoggedIn.value = token && token !== '';
});

// Set theme based on user preference
const savedTheme = localStorage.getItem('theme');
if (savedTheme) {
	isDarkMode.value = savedTheme === 'dark';
	document.documentElement.classList.toggle('dark', isDarkMode.value);
}
</script>

<style scoped>
.icon-size {
	width: 1rem;
	height: 1rem;
	font-size: 1rem;
	display: flex;
	align-items: center;
	justify-content: center;
}
</style>
