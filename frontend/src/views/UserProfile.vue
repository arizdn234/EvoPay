<template>
    <div class="user-profile max-w-4xl mx-auto p-5 mt-16">
        <h1 class="text-2xl font-bold mb-5">Your Profile</h1>

        <!-- Profile Information -->
        <div v-if="userData" class="profile-info bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 class="text-lg font-semibold mb-4">Profile Information</h2>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Name: </strong> <span class="text-grey-light dark:text-white">{{ userData.name }}</span>
            </div>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Email: </strong> <span class="text-grey-light dark:text-white">{{ userData.email }}</span>
            </div>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Account Created: </strong> <span class="text-grey-light dark:text-white">{{ formatDate(userData.created_at) }}</span>
            </div>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Email Verified: </strong> <span v-if="userData.email_verified" class="text-green-600 dark:text-green-400">Verified</span>
                <span v-else class="text-red-light dark:text-red-dark">Not Verified</span>
            </div>
        </div>

        <!-- Edit Profile Button -->
        <div v-if="userData" class="mt-6">
            <button @click="editProfile" class="bg-blue-light hover:bg-blue-dark text-white font-bold py-2 px-4 rounded">
                Edit Profile
            </button>
        </div>

        <!-- Transaction History Section -->
        <div v-if="transactions.length"
            class="transaction-history mt-10 bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 class="text-lg font-semibold mb-4">Transaction History</h2>
            <ul>
                <li v-for="transaction in transactions" :key="transaction.id" class="mb-4">
                    <div class="flex justify-between">
                        <div>
                            <strong>{{ transaction.type }}</strong>
                            <p class="text-gray-600 dark:text-gray-400">{{ formatDate(transaction.date) }}</p>
                        </div>
                        <p class="font-bold" :class="transaction.amount > 0 ? 'text-green-600' : 'text-red-600'">
                            {{ transaction.amount > 0 ? '+' : '' }}{{ transaction.currency }} {{
                                transaction.amount.toFixed(2) }}
                        </p>
                    </div>
                </li>
            </ul>
        </div>

        <!-- Account Settings Section -->
        <div v-if="userData" class="account-settings mt-10 bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 class="text-lg font-semibold mb-4">Account Settings</h2>
            <ul>
                <li class="mb-4">
                    <button @click="changePassword"
                        class="bg-red-light hover:bg-red-dark text-white font-regular py-2 px-4 rounded">
                        Change Password
                    </button>
                </li>
                <li class="mb-4">
                    <button @click="deleteAccount"
                        class="bg-red-light hover:bg-red-dark text-white font-regular py-2 px-4 rounded">
                        Delete Account
                    </button>
                </li>
            </ul>
        </div>

        <!-- Show loading if data is not yet available -->
        <div v-else class="text-center mt-10">
            <p>Loading user information...</p>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from '../axios';

// Refs
const userData = ref(null);
const transactions = ref([]);
const router = useRouter();

onMounted(async () => {
    try {
        // Make API call to fetch user data
        const response = await axios.get('/users/me', { withCredentials: true });
        userData.value = response.data;

        // // Fetch user transactions
        // const transactionsResponse = await axios.get(`/transactions/user/${userData.value.id}`);
        // transactions.value = transactionsResponse.data;
    } catch (error) {
        console.error('Error fetching user data or transactions:', error);
        if (error.response && error.response.status === 401) {
            router.push('/users/login'); // Redirect to login if not authorized
        }
    }
});

// Function to format date
const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return new Date(dateString).toLocaleDateString(undefined, options);
};

// Placeholder functions for actions
const editProfile = () => {
    alert('Edit profile functionality is not implemented yet.');
};

const changePassword = () => {
    alert('Change password functionality is not implemented yet.');
};

const deleteAccount = () => {
    alert('Delete account functionality is not implemented yet.');
};
</script>