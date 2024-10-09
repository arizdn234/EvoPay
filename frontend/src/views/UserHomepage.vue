<template>
	<div class="user-homepage p-5 max-w-4xl mx-auto mt-16">
		<h1 class="text-2xl font-bold mb-5">Hello, <span
				class="font-bold cursor-pointer text-blue-light hover:text-blue-dark">{{
					userName }}</span></h1>
		<p class="mt-5 text-sm italic">Welcome to your homepage!</p>

		<p v-if="balance !== null"
			class="text-2xl text-blue-light hover:text-blue-dark font-semibold pt-2 rounded hover:bg-blue-light hover:px-2 hover:dark:text-white hover:text-black">
			<strong class="text-gray-900 dark:text-white">Balance:</strong>
			{{ `${balance.currency} ${balance.current_balance.toFixed(2)}` }}
		</p>


		<!-- Conditionally display a message if no token is found -->
		<p v-if="cookieNotFound" class="text-red-light dark:text-red-dark mt-4">Cookie not found. Please log in.</p>

		<!-- Display additional user information if available -->
		<div v-if="userData"
			class="py-2 px-4 border border-secondary-light dark:border-grey4-light rounded-lg bg-transparent mt-5 hover:py-4">
			<h2 class="text-lg font-semibold underline mb-2">Your Account information</h2>
			<p><strong>Email:</strong> <span class="text-mint-light dark:text-mint-dark">{{ userData.email }}</span></p>
			<p><strong>Registered At:</strong> <span class="text-blue-light dark:text-blue-dark">{{
				formatDate(userData.created_at) }}</span></p>
			<p>
				<strong>Email Verified Status: </strong>
				<!-- Conditionally render verified or not verified -->
				<span v-if="userData.email_verified" class="text-green-light dark:text-green-dark">Email Verified</span>
				<span v-else class="text-red-light dark:text-red-dark">Email Not Verified</span>
			</p>
		</div>

		<!-- Promotional Banner -->
		<div
			class="promotional-banner mt-8 bg-blue-light dark:bg-blue-dark text-black dark:text-white rounded-lg p-5 shadow-lg flex justify-between items-center">
			<!-- Banner text -->
			<div>
				<h2 class="text-xl font-bold">Get 20% Cashback on Your Next Transaction!</h2>
				<p class="text-sm mt-2">Use code <span class="font-semibold">CASHBACK20</span> at checkout to enjoy this
					limited time offer.</p>
				<p class="text-xs mt-1 italic">*Offer valid until {{ formatDate(new Date()) }}</p>
			</div>
			<!-- Banner image -->
			<div>
				<img src="https://via.placeholder.com/150x100" alt="Cashback Offer" class="rounded-lg">
			</div>
		</div>

		<!-- Feature buttons section -->
		<div class="feature-buttons font-sfprotext grid grid-cols-3 sm:grid-cols-4 gap-4 mt-10">
			<!-- Button 1 -->
			<button @click="handleFeature1"
				class="flex flex-col items-center text-white font-bold py-4 px-4 scale-100 hover:scale-105">
				<div class="bg-blue-light dark:bg-blue-dark rounded-full p-2 mb-2">
					<i class='bx bx-wallet text-2xl icon-size text-white'></i>
				</div>
				<span class="text-blue-light dark:text-blue-dark">Deposit</span>
			</button>

			<!-- Button 2 -->
			<button @click="handleFeature2"
				class="flex flex-col items-center text-white font-bold py-4 px-4 scale-100 hover:scale-105">
				<div class="bg-blue-light dark:bg-blue-dark rounded-full p-2 mb-2">
					<i class='bx bx-money-withdraw text-2xl icon-size text-white'></i>
				</div>
				<span class="text-blue-light dark:text-blue-dark">Withdraw</span>
			</button>

			<!-- Button 3 -->
			<button @click="handleFeature3"
				class="flex flex-col items-center text-white font-bold py-4 px-4 scale-100 hover:scale-105">
				<div class="bg-blue-light dark:bg-blue-dark rounded-full p-2 mb-2">
					<i class='bx bx-history text-2xl icon-size text-white'></i>
				</div>
				<span class="text-blue-light dark:text-blue-dark">Your History</span>
			</button>

			<!-- Button 4 -->
			<button @click="handleFeature4"
				class="flex flex-col items-center text-white font-bold py-4 px-4 scale-100 hover:scale-105">
				<div class="bg-blue-light dark:bg-blue-dark rounded-full p-2 mb-2">
					<i class='bx bx-cart text-2xl icon-size text-white'></i>
				</div>
				<span class="text-blue-light dark:text-blue-dark">Pay Amazon</span>
			</button>

			<!-- Button 5 -->
			<button @click="handleFeature5"
				class="flex flex-col items-center text-white font-bold py-4 px-4 scale-100 hover:scale-105">
				<div class="bg-blue-light dark:bg-blue-dark rounded-full p-2 mb-2">
					<i class='bx bx-transfer-alt text-2xl icon-size text-white'></i>
				</div>
				<span class="text-blue-light dark:text-blue-dark">Bank Transfer</span>
			</button>

			<!-- Button 6 -->
			<button @click="handleFeature6"
				class="flex flex-col items-center text-white font-bold py-4 px-4 scale-100 hover:scale-105">
				<div class="bg-blue-light dark:bg-blue-dark rounded-full p-2 mb-2">
					<i class='bx bx-credit-card text-2xl icon-size text-white'></i>
				</div>
				<span class="text-blue-light dark:text-blue-dark">QevoPay</span>
			</button>
		</div>

		<!-- Dashboard Graph Placeholder -->
		<div class="dashboard-graph mt-10 p-5 border border-gray-300 rounded-lg bg-transparent shadow-md">
			<h2 class="text-lg font-bold mb-5">Dashboard Overview</h2>
			<!-- Placeholder for Graph -->
			<div class="h-64 bg-grey-light dark:bg-grey-dark flex justify-center items-center rounded-lg">
				<p class="text-gray-500">[Graph displaying user transaction statistics]</p>
			</div>
		</div>

		<!-- FAQ Section -->
		<div class="faq-section mt-10">
			<h2 class="text-lg font-bold underline">Frequently Asked Questions (FAQ)</h2>
			<ul class="mt-4">
				<li class="mb-3">
					<strong>1. How do I deposit money?</strong>
					<p class="text-sm text-grey-light dark:text-grey-dark">You can deposit money by clicking the "Deposit" button and
						following the instructions to link your bank account or credit card.</p>
				</li>
				<li class="mb-3">
					<strong>2. Is my information secure?</strong>
					<p class="text-sm text-grey-light dark:text-grey-dark">Yes, we use top-tier encryption methods to protect your data. Your
						information is stored securely and is never shared without your consent.</p>
				</li>
				<li class="mb-3">
					<strong>3. What is the maximum withdrawal limit?</strong>
					<p class="text-sm text-grey-light dark:text-grey-dark">The maximum withdrawal limit depends on your account type. For
						verified users, the limit is higher. You can find more details in your account settings.</p>
				</li>
				<li class="mb-3">
					<strong>4. How do I contact customer support?</strong>
					<p class="text-sm text-grey-light dark:text-grey-dark">You can contact our 24/7 support team by clicking on the "Help &
						Support" button or by emailing us at support@evopay.com.</p>
				</li>
			</ul>
		</div>

		<!-- Help & Support Section -->
		<div class="help-support-section mt-10">
			<h2 class="text-lg font-bold underline">Help & Support</h2>
			<p class="mt-4 text-sm">If you need further assistance, feel free to reach out to our support team.</p>
			<button @click="contactSupport" class="mt-4 py-2 px-4 bg-blue-light text-white rounded-lg hover:bg-blue-dark">
				Contact Support
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

<style scoped>
.icon-size {
	width: 2.4rem;
	height: 2.4rem;
	font-size: 2rem;
	display: flex;
	align-items: center;
	justify-content: center;
}
</style>