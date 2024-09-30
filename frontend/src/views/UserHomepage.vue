<template>
	<div class="user-homepage p-5 max-w-4xl mx-auto">
		<h1 class="text-2xl font-bold mb-5">Hello, {{ userName }}</h1>
		<p class="mt-5 text-sm italic">Welcome to your homepage!</p>

		<!-- Conditionally display a message if no token is found -->
		<p v-if="cookieNotFound" class="text-red-500 mt-4">Cookie not found. Please log in.</p>

		<!-- Display additional user information if available -->
		<div v-if="userData" class="py-2 px-4 border rounded-lg bg-gray-300 dark:bg-gray-900 mt-5">
			<h2 class="text-lg font-semibold">Your Information:</h2>
			<p><strong>Email:</strong> {{ userData.email }}</p>
			<p><strong>Registered At:</strong> {{ formatDate(userData.created_at) }}</p>
			<p><strong>Email Verified Status:</strong> {{ userData.email_verified }}</p>
		</div>
	</div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import Cookies from 'js-cookie';
import { useRouter } from 'vue-router';
import axios from '../axios';

// Set up refs
const userName = ref('');
const cookieNotFound = ref(false);
const userData = ref(null);
const router = useRouter();

onMounted(async () => {
	const token = Cookies.get('auth_token');

	// Check if the token exists in cookies
	if (!token) {
		cookieNotFound.value = true;
		setTimeout(() => {
			router.push('/users/login');
		}, 3000);
		return;
	}

	try {
		// Make API call to fetch user data (cookie will be sent automatically)
		const response = await axios.get('/users/me', { withCredentials: true });

		// Store the fetched user data
		userData.value = response.data;
		userName.value = userData.value.name;
	} catch (error) {
		console.error('Error fetching user data:', error);
	}
});

// Function to format the date (if needed)
const formatDate = (dateString) => {
	const options = { year: 'numeric', month: 'long', day: 'numeric' };
	return new Date(dateString).toLocaleDateString(undefined, options);
};
</script>
