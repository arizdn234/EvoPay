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

					<button @click="toggleDarkMode"
						class="text-white dark:text-gray-200 p-2 rounded-lg hover:bg-gray-700 dark:hover:bg-gray-600">
						<i :class="isDarkMode ? 'fas fa-sun' : 'fas fa-moon'"></i>
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
import { ref, onMounted } from 'vue'
import { RouterLink, RouterView } from 'vue-router'

const isDarkMode = ref(false)

const toggleDarkMode = () => {
	isDarkMode.value = !isDarkMode.value
	document.documentElement.classList.toggle('dark', isDarkMode.value)
	localStorage.setItem('theme', isDarkMode.value ? 'dark' : 'light')
}

onMounted(() => {
	const savedTheme = localStorage.getItem('theme')
	if (savedTheme) {
		isDarkMode.value = savedTheme === 'dark'
		document.documentElement.classList.toggle('dark', isDarkMode.value)
	}
})
</script>
