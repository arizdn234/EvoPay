<template>
	<div :class="{ 'dark': isDarkMode }">
		<header class="fixed top-0 left-0 right-0 z-50 px-4 py-5 bg-white-light dark:bg-black-dark shadow-md">
			<div class="container mx-auto flex justify-between items-center">
				<!-- Logo -->
				<div class="flex space-x-4">
					<p
						class="text-2xl cursor-pointer font-sans font-bold text-blue-light dark:text-blue-dark hover:text-cyan-light dark:hover:text-cyan-dark">
						EvoPay
					</p>
				</div>

				<!-- Menu Icon for Mobile -->
				<div class="md:hidden flex items-center">
					<button @click="isMenuOpen = !isMenuOpen" class="text-blue-light dark:text-blue-light z-1">
						<i :class="isMenuOpen ? 'bx bx-x bx-md' : 'bx bx-menu bx-md'" class="icon-size"></i>
					</button>
				</div>

				<!-- Navigation Links (dekstop) -->
				<nav class="hidden md:flex items-center space-x-4"
					:class="{ hidden: !isMenuOpen, 'block': isMenuOpen }">

					<RouterLink to="/users/homepage"
						class="text-grey6-dark dark:text-gray-200 hover:text-gray-700 dark:hover:text-gray-100"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold" v-if="isLoggedIn">
						Home
					</RouterLink>
					
					<RouterLink to="/"
						class="text-grey6-dark dark:text-gray-200 hover:text-gray-700 dark:hover:text-gray-100"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold">
						API Docs
					</RouterLink>
					
					<!-- Login/Register links (shown when not logged in) -->
					<RouterLink to="/users/login"
						class="text-grey6-dark dark:text-gray-200 hover:text-gray-700 dark:hover:text-gray-100"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold"
						v-if="!isLoggedIn">
						Login
					</RouterLink>
					<RouterLink to="/users/register"
						class="text-grey6-dark dark:text-gray-200 hover:text-gray-700 dark:hover:text-gray-100"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold"
						v-if="!isLoggedIn">
						Register
					</RouterLink>

					<!-- Dark Mode Toggle -->
					<button @click="toggleDarkMode"
						class="text-blue-light dark:text-blue-dark p-2">
						<i :class="isDarkMode ? 'bx bxs-sun' : 'bx bxs-moon'" class="icon-size"></i>
					</button>

					<!-- Logout (shown when logged in) -->
					<button @click="handleLogout"
						class="text-white py-2 px-4 rounded-lg bg-rose-600 hover:bg-rose-500 dark:hover:bg-rose-500"
						v-if="isLoggedIn">
						Logout
					</button>
				</nav>
			</div>

			<!-- Navigation Links (mobile) -->
			<div class="md:hidden flex flex-col items-center space-y-4 py-4 text-black-dark dark:text-tertiary-dark"
				v-if="isMenuOpen">

				<RouterLink to="/users/homepage"
						class="hover:text-white dark:hover:text-gray-100"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold" v-if="isLoggedIn">
						Home
				</RouterLink>

				<RouterLink to="/" class="hover:text-white dark:hover:text-gray-100"
					active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold">
					API Docs
				</RouterLink>
				
				<!-- Login/Register links (for mobile) -->
				<RouterLink to="/users/login"
					class="w-full text-center hover:text-white dark:hover:text-gray-100"
					v-if="!isLoggedIn">
					Login
				</RouterLink>
				<RouterLink to="/users/register"
					class="w-full text-center hover:text-white dark:hover:text-gray-100"
					v-if="!isLoggedIn">
					Register
				</RouterLink>

				<!-- Dark Mode Toggle (for mobile) -->
				<button @click="toggleDarkMode"
					class="text-white dark:text-gray-200 p-2 rounded-lg">
					<i :class="isDarkMode ? 'bx bxs-sun' : 'bx bxs-moon'" class="icon-size"></i>
				</button>

				<!-- Logout (for mobile) -->
				<button @click="handleLogout"
					class="text-white py-2 px-4 rounded-lg bg-rose-600 hover:bg-rose-500 dark:hover:bg-rose-500"
					v-if="isLoggedIn">
					Logout
				</button>
			</div>
		</header>

		<main class="font-sfpro p-4 bg-gray-200 dark:bg-gray-800 text-gray-900 dark:text-gray-100 min-h-screen">
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

// State for the menu and dark mode
const isMenuOpen = ref(false);
const isDarkMode = ref(false);
const isLoggedIn = ref(false);

const router = useRouter();

// Dark Mode Toggle
const toggleDarkMode = () => {
	isDarkMode.value = !isDarkMode.value;
	document.documentElement.classList.toggle('dark', isDarkMode.value);
	localStorage.setItem('theme', isDarkMode.value ? 'dark' : 'light');
};

// Handle Logout
const handleLogout = async () => {
	try {
		await axios.post('users/logout', {}, { withCredentials: true });
		Cookies.remove('auth_token');
		isLoggedIn.value = false;
		router.push('/users/login');
	} catch (error) {
		console.error('Logout error:', error);
	}
};

// Watch for authentication state
watchEffect(() => {
	const token = Cookies.get('auth_token');
	isLoggedIn.value = token && token !== '';
});

// Load the saved theme from localStorage
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

@media (max-width: 768px) {
	.header {
		flex-direction: column;
	}
}
</style>
