<template>
    <div class="user-profile max-w-4xl mx-auto p-5 mt-10">
      <h1 class="text-2xl font-bold mb-5">Your Profile</h1>
  
      <!-- Profile Information -->
      <div class="profile-info bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
        <h2 class="text-lg font-semibold mb-4">Profile Information</h2>
        <div class="mb-3">
          <strong class="block text-gray-700 dark:text-gray-300">Name:</strong>
          <p class="text-gray-900 dark:text-white">{{ user.name }}</p>
        </div>
        <div class="mb-3">
          <strong class="block text-gray-700 dark:text-gray-300">Email:</strong>
          <p class="text-gray-900 dark:text-white">{{ user.email }}</p>
        </div>
        <div class="mb-3">
          <strong class="block text-gray-700 dark:text-gray-300">Account Created:</strong>
          <p class="text-gray-900 dark:text-white">{{ formatDate(user.created_at) }}</p>
        </div>
        <div class="mb-3">
          <strong class="block text-gray-700 dark:text-gray-300">Email Verified:</strong>
          <span v-if="user.email_verified" class="text-green-600 dark:text-green-400">Verified</span>
          <span v-else class="text-red-600 dark:text-red-400">Not Verified</span>
        </div>
      </div>
  
      <!-- Edit Profile Button -->
      <div class="mt-6">
        <button @click="editProfile"
          class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded">
          Edit Profile
        </button>
      </div>
  
      <!-- Transaction History Section -->
      <div class="transaction-history mt-10 bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
        <h2 class="text-lg font-semibold mb-4">Transaction History</h2>
        <ul>
          <li v-for="transaction in transactions" :key="transaction.id" class="mb-4">
            <div class="flex justify-between">
              <div>
                <strong>{{ transaction.type }}</strong>
                <p class="text-gray-600 dark:text-gray-400">{{ formatDate(transaction.date) }}</p>
              </div>
              <p class="font-bold" :class="transaction.amount > 0 ? 'text-green-600' : 'text-red-600'">
                {{ transaction.amount > 0 ? '+' : '' }}{{ transaction.currency }} {{ transaction.amount.toFixed(2) }}
              </p>
            </div>
          </li>
        </ul>
      </div>
  
      <!-- Account Settings Section -->
      <div class="account-settings mt-10 bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
        <h2 class="text-lg font-semibold mb-4">Account Settings</h2>
        <ul>
          <li class="mb-4">
            <button @click="changePassword"
              class="bg-yellow-500 hover:bg-yellow-600 text-white font-bold py-2 px-4 rounded">
              Change Password
            </button>
          </li>
          <li class="mb-4">
            <button @click="deleteAccount"
              class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded">
              Delete Account
            </button>
          </li>
        </ul>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        user: {
          name: 'John Doe',
          email: 'john.doe@example.com',
          email_verified: true,
          created_at: '2021-09-15T12:00:00Z'
        },
        transactions: [
          { id: 1, type: 'Deposit', amount: 200, currency: 'USD', date: '2024-09-10' },
          { id: 2, type: 'Withdrawal', amount: -50, currency: 'USD', date: '2024-09-11' },
          { id: 3, type: 'Payment', amount: -30, currency: 'USD', date: '2024-09-12' }
        ]
      }
    },
    methods: {
      formatDate(date) {
        const options = { year: 'numeric', month: 'long', day: 'numeric' }
        return new Date(date).toLocaleDateString(undefined, options)
      },
      editProfile() {
        alert('Edit profile feature coming soon.')
      },
      changePassword() {
        alert('Change password feature coming soon.')
      },
      deleteAccount() {
        alert('Delete account feature coming soon.')
      }
    }
  }
  </script>
  
  <style scoped>
  .user-profile {
    background-color: #f9fafb;
  }
  
  .profile-info,
  .transaction-history,
  .account-settings {
    background-color: white;
  }
  
  .dark .profile-info,
  .dark .transaction-history,
  .dark .account-settings {
    background-color: #2d3748;
  }
  </style>
  