<template>
	<div class="user-homepage p-5 max-w-4xl mx-auto">
		<h1 class="text-2xl font-bold mb-5">Hello, {{ userName }}</h1>
		<p class="mt-5 text-sm italic">Welcome to your homepage!</p>

		<!-- Conditionally display a message if no token is found -->
		<p v-if="cookieNotFound" class="text-red-500 mt-4">Cookie not found. Please log in.</p>
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
});
</script>

