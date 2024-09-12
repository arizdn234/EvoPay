<template>
	<div :class="{ 'dark': isDarkMode }">
		<header class="p-4 bg-gray-800 dark:bg-gray-900">
			<div class="container mx-auto flex justify-between items-center">
				<nav class="flex space-x-4">
					<RouterLink to="/"
						class="text-gray-300 dark:text-gray-200 hover:text-white dark:hover:text-gray-100"
						active-class="text-indigo-400 dark:text-indigo-300" exact-active-class="font-bold">
						API Docs
					</RouterLink>
					<RouterLink to="/users/login"
						class="text-gray-300 dark:text-gray-200 hover:text-white dark:hover:text-gray-100"
						active-class="text-indigo-400 dark:text-indigo-300" exact-active-class="font-bold">
						Login
					</RouterLink>
					<RouterLink to="/users/register"
						class="text-gray-300 dark:text-gray-200 hover:text-white dark:hover:text-gray-100"
						active-class="text-indigo-400 dark:text-indigo-300" exact-active-class="font-bold">
						Register
					</RouterLink>
				</nav>

				<button @click="toggleDarkMode"
					class="text-white dark:text-gray-200 p-2 rounded-lg hover:bg-gray-700 dark:hover:bg-gray-600">
					<i :class="isDarkMode ? 'fas fa-sun' : 'fas fa-moon'"></i>
				</button>
			</div>
		</header>
		<main class="p-4 bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100">
			<!-- Your content here -->
			<RouterView />
		</main>
	</div>
</template>

<script setup>
import { ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'

const isDarkMode = ref(false)

const toggleDarkMode = () => {
	isDarkMode.value = !isDarkMode.value
	document.documentElement.classList.toggle('dark', isDarkMode.value)
}
</script>
