<template>
    <div class="user-profile max-w-4xl mx-auto p-5 mt-16">
        <h1 class="text-2xl font-bold mb-5">Your Profile</h1>

        <!-- Profile Information -->
        <div v-if="userData" class="profile-info bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 class="text-lg font-semibold mb-4">Profile Information</h2>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Name: </strong> <span
                    class="text-grey-light dark:text-white">{{ userData.name }}</span>
            </div>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Email: </strong> <span
                    class="text-grey-light dark:text-white">{{ userData.email }}</span>
            </div>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Account Created: </strong> <span
                    class="text-grey-light dark:text-white">{{ diffForHuman(userData.created_at) }}</span> | <span
                    class="text-grey2-light dark:text-grey2-dark">{{ userData.created_at }}</span>
            </div>
            <div class="mb-3">
                <strong class="text-blue-light dark:text-blue-dark">Email Verified: </strong> <span
                    v-if="userData.email_verified" class="text-green-600 dark:text-green-400">Verified</span>
                <span v-else class="text-red-light dark:text-red-dark">Not Verified</span>
            </div>
        </div>

        <!-- Edit Profile Button -->
        <div v-if="userData" class="mt-6">
            <button @click="editProfile"
                class="bg-blue-light dark:bg-blue-dark hover:filter hover:brightness-125 text-white font-bold py-2 px-4 rounded-lg">
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
                            <span class="text-grey-light dark:text-white">{{
                                diffForHuman(formatDate(transaction.created_at)) }}</span> | <span
                                class="text-grey2-light dark:text-grey2-dark">{{ formatDate(transaction.created_at) }}
                            </span>
                        </div>
                        <p class="font-bold"
                            :class="transaction.amount > 0 ? 'text-green-light dark:text-green-dark' : 'text-red-light dark:text-red-dark'">
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
                        class="bg-red-light dark:bg-red-dark hover:filter hover:brightness-125 text-white font-regular py-2 px-4 rounded-lg">
                        Change Password
                    </button>
                </li>
                <li class="mb-4">
                    <button @click="deleteAccount"
                        class="bg-red-light dark:bg-red-dark hover:filter hover:brightness-125 text-white font-regular py-2 px-4 rounded-lg">
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
import { diffForHuman, formatDate } from '@/dateUtils';

// Refs
const userData = ref(null);
const transactions = ref([]);
const router = useRouter();

onMounted(async () => {
    try {
        // Make API call to fetch user data
        const response = await axios.get('/users/me', { withCredentials: true });
        userData.value = response.data;
        userData.value.created_at = formatDate(userData.value.created_at)

        // Fetch user transactions
        const transactionsResponse = await axios.get(`/transactions/user/${userData.value.id}`, { withCredentials: true });
        transactions.value = transactionsResponse.data;
    } catch (error) {
        console.error('Error fetching user data or transactions:', error);
        if (error.response && error.response.status === 401) {
            router.push('/users/login');
        }
    }
});

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