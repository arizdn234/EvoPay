<template>
	<div class="user-homepage p-5 max-w-4xl mx-auto">
		<h1 class="text-2xl font-bold mb-5">Hello, <span class="font-bold transition-all duration-300 ease-in-out cursor-pointer text-blue-700 hover:text-blue-900">{{ userName }}</span></h1>
		<p class="mt-5 text-sm italic">Welcome to your homepage!</p>

		<p v-if="balance !== null"
			class="text-2xl text-blue-500 font-bold pt-2 rounded transition-all duration-300 ease-in-out hover:bg-blue-400 hover:px-2 hover:text-white">
			<strong class="text-gray-900 dark:text-white">Balance:</strong>
			{{ `${balance.currency} ${balance.current_balance.toFixed(2)}` }}
		</p>


		<!-- Conditionally display a message if no token is found -->
		<p v-if="cookieNotFound" class="text-red-500 mt-4">Cookie not found. Please log in.</p>

		<!-- Display additional user information if available -->
		<div v-if="userData" class="transition-all duration-300 ease-in-out py-2 px-4 border rounded-lg bg-gray-300 dark:bg-gray-900 mt-5 hover:py-4">
			<h2 class="text-lg font-semibold underline mb-2">Your Account information</h2>
			<p><strong>Email:</strong> <span class="text-green-500">{{ userData.email }}</span></p>
			<p><strong>Registered At:</strong> <span class="text-blue-400">{{ formatDate(userData.created_at) }}</span></p>
			<p>
				<strong>Email Verified Status: </strong>
				<!-- Conditionally render verified or not verified -->
				<span v-if="userData.email_verified" class="text-green-400">Email Verified</span>
				<span v-else class="text-yellow-400">Email Not Verified</span>
			</p>
		</div>

		<!-- Feature buttons section -->
		<div class="feature-buttons grid grid-cols-2 sm:grid-cols-3 gap-4 mt-10 transition-all duration-300 ease-in-out hover:px-2">
			<!-- Button 1 -->
			<button @click="handleFeature1"
				class="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-700 hover:py-4 transition-all duration-300 ease-in-out">
				Deposit
			</button>

			<!-- Button 2 -->
			<button @click="handleFeature2"
				class="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-700 hover:py-4 transition-all duration-300 ease-in-out">
				Withdraw
			</button>

			<!-- Button 3 -->
			<button @click="handleFeature3"
				class="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-700 hover:py-4 transition-all duration-300 ease-in-out">
				Your History
			</button>

			<!-- Button 4 -->
			<button @click="handleFeature4"
				class="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-700 hover:py-4 transition-all duration-300 ease-in-out">
				Pay Amazon
			</button>

			<!-- Button 5 -->
			<button @click="handleFeature5"
				class="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-700 hover:py-4 transition-all duration-300 ease-in-out">
				Bank Transfer
			</button>

			<!-- Button 6 -->
			<button @click="handleFeature6"
				class="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-700 hover:py-4 transition-all duration-300 ease-in-out">
				QevoPay
			</button>
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
const balance = ref(null);
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

		// Fetch balance separately
		const balanceResponse = await axios.get(`/balances/${userData.value.id}`, { withCredentials: true });
		balance.value = {
			current_balance: balanceResponse.data.current_balance,
			currency: balanceResponse.data.currency
		};
	} catch (error) {
		console.error('Error fetching user data or balance:', error);
	}
});

// Function to format the date (if needed)
const formatDate = (dateString) => {
	const options = { year: 'numeric', month: 'long', day: 'numeric' };
	return new Date(dateString).toLocaleDateString(undefined, options);
};

// Placeholder functions for feature buttons
const handleFeature1 = () => {
	console.log('Feature 1 clicked');
};

const handleFeature2 = () => {
	console.log('Feature 2 clicked');
};

const handleFeature3 = () => {
	console.log('Feature 3 clicked');
};

const handleFeature4 = () => {
	console.log('Feature 4 clicked');
};

const handleFeature5 = () => {
	console.log('Feature 5 clicked');
};

const handleFeature6 = () => {
	console.log('Feature 6 clicked');
};
</script>
