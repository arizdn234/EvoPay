<template>
	<div class="user-homepage p-5 max-w-4xl mx-auto">
		<h1 class="text-2xl font-bold mb-5">Hello, {{ userName }}</h1>
		<p class="mt-5 text-sm italic">Welcome to your homepage!</p>

		<!-- Conditionally display a message if no token is found -->
		<p v-if="cookieNotFound" class="text-red-500 mt-4">Cookie not found. Please log in.</p>

		<!-- Display additional user information if available -->
		<div v-if="userData" class="user-info mt-5">
			<h2 class="text-lg font-semibold">Your Information:</h2>
			<p><strong>Email:</strong> {{ userData.email }}</p>
			<p><strong>Registered At:</strong> {{ formatDate(userData.createdAt) }}</p>
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
const userData = ref(null); // Store user data
const router = useRouter();

onMounted(async () => {
	const token = Cookies.get('auth_token');

	if (!token) {
		cookieNotFound.value = true;
		setTimeout(() => {
			router.push('/users/login');
		}, 3000);
		return;
	}

	try {
		// Make API call to fetch user data
		const response = await axios.get('/users/me', {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});
		userData.value = response.data; // Store user data
		userName.value = userData.value.name; // Assuming the API returns the user's name
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

<style scoped>
.user-info {
	padding: 1rem;
	background-color: #f9fafb; /* Light gray background */
	border-radius: 0.5rem; /* Rounded corners */
}
</style>
