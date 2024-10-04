<template>
	<div class="documentation p-5 max-w-4xl mx-auto mt-16">
		<h1 class="text-2xl font-bold mb-5">{{ info.service }} API Documentation</h1>
		<p class="text-lg mb-5">Version: {{ info.version }}</p>

		<!-- Toggle Buttons for User and Admin Routes -->
		<div class="flex space-x-4 mb-5">
			<button 
				@click="toggleRoutes('user')" 
				:class="{'bg-blue-light dark:bg-blue-dark text-white scale-105': currentView === 'user', 'border border-blue-dark scale-100': currentView !== 'user'}" 
				class="px-4 py-2 w-40 rounded-lg transition duration-150 ease-in-out">
				User Endpoints
			</button>
			<button 
				@click="toggleRoutes('admin')" 
				:class="{'bg-blue-light dark:bg-blue-dark text-white scale-105': currentView === 'admin', 'border border-blue-dark scale-100': currentView !== 'admin'}" 
				class="px-4 py-2 w-40 rounded-lg transition duration-150 ease-in-out">
				Admin Endpoints
			</button>
		</div>

		<!-- Conditional Table Rendering -->
		<div class="table-container overflow-x-auto xl:w-max">
			<table class="min-w-full border-collapse mt-5">
				<thead>
					<tr>
						<th class="border border-gray-400 px-4 py-2 bg-gray-800 text-white text-left">Method</th>
						<th class="border border-gray-400 px-4 py-2 bg-gray-800 text-white text-left">Path</th>
						<th class="border border-gray-400 px-4 py-2 bg-gray-800 text-white text-left">Description</th>
						<th class="border border-gray-400 px-4 py-2 bg-gray-800 text-white text-left">Note</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(route, index) in currentRoutes" :key="index"
						class="even:dark:bg-gray-800 odd:dark:bg-gray-700 even:bg-gray-400 hover:bg-gray-700 hover:text-white transition-all duration-300">
						<td class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.method }}</td>
						<td class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.path }}</td>
						<td class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.description }}</td>
						<td v-if="route.note" class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.note }}</td>
						<td v-else class="border border-gray-400 px-4 py-2 whitespace-nowrap">-</td>
					</tr>
				</tbody>
			</table>
		</div>

		<p class="mt-5 text-sm italic">*{{ info.note }}</p>
	</div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from '../axios';

const info = ref({
    user_routes: [],
    admin_routes: [],
    note: '',
    version: '',
    service: ''
});
const currentView = ref('user');

const fetchApiInfo = async () => {
    try {
        const response = await axios.get('/');

        info.value = response.data; 
        if (response.data.user_routes) {
            info.value.user_routes = response.data.user_routes; // Set user routes directly
        }

        if (response.data.admin_routes) {
            info.value.admin_routes = response.data.admin_routes; // Set admin routes directly
        }

    } catch (error) {
        console.error('Error fetching documentation:', error);
    }
};

// Toggle function to switch between user and admin routes
const toggleRoutes = (view) => {
    currentView.value = view;
};

// Use computed to dynamically determine the routes to display
const currentRoutes = computed(() => {
    return currentView.value === 'admin' ? info.value.admin_routes : info.value.user_routes;
});

onMounted(fetchApiInfo);
</script>
