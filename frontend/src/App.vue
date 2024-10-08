<template>
	<div :class="{ 'dark': isDarkMode }">
		<header style="background: rgba(5, 16, 37, 0.12); backdrop-filter: blur(12px);"
			class="fixed top-0 left-0 right-0 z-50 px-4 py-5 shadow-md">
			<div class="container mx-auto flex justify-between items-center">
				<!-- Logo -->
				<div class="flex space-x-4">
					<p
						class="text-2xl cursor-pointer font-sans font-bold text-blue-light dark:text-blue-dark hover:filter hover:brightness-125">
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
						class="text-grey6-dark dark:text-gray-200 hover:filter hover:brightness-125"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold"
						v-if="isLoggedIn">
						Home
					</RouterLink>

					<RouterLink to="/" class="text-grey6-dark dark:text-gray-200 hover:filter hover:brightness-125"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold">
						API Docs
					</RouterLink>

					<RouterLink to="/users/me"
						class="text-blue-light dark:text-blue-dark hover:filter hover:brightness-125" v-if="isLoggedIn">
						<i class='bx bx-user text-2xl'></i>
					</RouterLink>

					<!-- Login/Register links (shown when not logged in) -->
					<RouterLink to="/users/login"
						class="text-grey6-dark dark:text-gray-200 hover:filter hover:brightness-125"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold"
						v-if="!isLoggedIn">
						Login
					</RouterLink>
					<RouterLink to="/users/register"
						class="text-grey6-dark dark:text-gray-200 hover:filter hover:brightness-125"
						active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold"
						v-if="!isLoggedIn">
						Register
					</RouterLink>

					<!-- Dark Mode Toggle -->
					<button @click="toggleDarkMode" class="text-blue-light dark:text-blue-dark p-2">
						<i :class="isDarkMode ? 'bx bxs-sun' : 'bx bxs-moon'" class="icon-size"></i>
					</button>

					<!-- Logout (shown when logged in) -->
					<button @click="handleLogout"
						class="text-white py-2 px-4 rounded-lg bg-red-light hover:bg-red-dark dark:hover:bg-rose-500"
						v-if="isLoggedIn">
						Logout
					</button>
				</nav>
			</div>

			<!-- Navigation Links (mobile) -->
			<div class="md:hidden flex flex-col items-center space-y-4 py-4 text-black-dark dark:text-tertiary-dark"
				v-if="isMenuOpen">

				<RouterLink to="/users/homepage" class="text-grey6-dark dark:text-gray-100"
					active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold" v-if="isLoggedIn">
					Home
				</RouterLink>

				<RouterLink to="/" class="text-grey6-dark dark:text-gray-100"
					active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold">
					API Docs
				</RouterLink>

				<!-- Login/Register links (for mobile) -->
				<RouterLink to="/users/login" class="w-full text-center text-grey6-dark dark:text-gray-100"
					active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold"
					v-if="!isLoggedIn">
					Login
				</RouterLink>
				<RouterLink to="/users/register" class="w-full text-center text-grey6-dark dark:text-gray-100"
					active-class="text-blue-light dark:text-blue-dark" exact-active-class="font-bold"
					v-if="!isLoggedIn">
					Register
				</RouterLink>

				<!-- User Profile link (for mobile) -->
				<RouterLink to="/users/me" class="text-blue-light dark:text-blue-dark hover:filter hover:brightness-125"
					v-if="isLoggedIn">
					<i class='bx bx-user text-2xl'></i>
				</RouterLink>

				<!-- Dark Mode Toggle (for mobile) -->
				<button @click="toggleDarkMode" class="text-blue-light dark:text-blue-dark p-2 rounded-lg">
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

		<!-- Main Contents -->
		<main class="font-sfpro p-4 bg-gray-200 dark:bg-gray-800 text-gray-900 dark:text-gray-100 min-h-screen">
			<RouterView />
		</main>

		<!-- Footer -->
		<footer class="bg-gray-900 dark:bg-gray-800 text-white py-6 mt-10">
			<div class="container mx-auto px-4">
				<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
					<!-- Column 1: Links -->
					<div>
						<h3 class="font-semibold mb-4">Quick Links</h3>
						<ul>
							<li class="mb-2">
								<a href="#" class="hover:underline text-gray-400">Home</a>
							</li>
							<li class="mb-2">
								<a href="#" class="hover:underline text-gray-400">FAQ</a>
							</li>
							<li class="mb-2">
								<a href="#" class="hover:underline text-gray-400">Help & Support</a>
							</li>
							<li>
								<a href="#" class="hover:underline text-gray-400">Privacy Policy</a>
							</li>
						</ul>
					</div>

					<!-- Column 2: Contact Info -->
					<div>
						<h3 class="font-semibold mb-4">Contact Us</h3>
						<p class="text-gray-400">
							Email: <a href="mailto:support@evopay.com"
								class="text-blue-400 hover:underline">support@evopay.com</a>
						</p>
						<p class="text-gray-400">Phone: +62 815 2947 4830</p>
						<p class="text-gray-400">24/7 Customer Support</p>
					</div>

					<!-- Column 3: Social Media -->
					<div>
						<h3 class="font-semibold mb-4">Follow Us</h3>
						<div class="flex space-x-4">
							<a href="#" class="text-gray-400 hover:text-white">
								<i class='bx bxl-facebook text-2xl'></i>
							</a>
							<a href="#" class="text-gray-400 hover:text-white">
								<i class='bx bxl-twitter text-2xl'></i>
							</a>
							<a href="#" class="text-gray-400 hover:text-white">
								<i class='bx bxl-instagram text-2xl'></i>
							</a>
							<a href="#" class="text-gray-400 hover:text-white">
								<i class='bx bxl-linkedin text-2xl'></i>
							</a>
						</div>
					</div>
				</div>

				<!-- Footer Bottom -->
				<div class="mt-10 border-t border-gray-700 pt-4 text-center text-gray-400">
					<p>&copy; 2024 EvoPay. All rights reserved.</p>
				</div>
			</div>
		</footer>
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
	width: 1.4rem;
	height: 1.4rem;
	font-size: 1.4rem;
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
