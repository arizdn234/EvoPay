<template>
	<div class="documentation p-5 max-w-4xl mx-auto">
		<h1 class="text-2xl font-bold mb-5">{{ info.service }} API Documentation</h1>
		<p class="text-lg mb-5">Version: {{ info.version }}</p>

		<!-- Toggle Buttons for User and Admin Routes -->
		<div class="flex space-x-4 mb-5">
			<button 
				@click="toggleRoutes('user')" 
				:class="{'bg-indigo-700 text-white scale-105': currentView === 'user', 'border border-indigo-400 scale-100': currentView !== 'user'}" 
				class="px-4 py-2 w-40 rounded-lg transition duration-150 ease-in-out">
				User Endpoints
			</button>
			<button 
				@click="toggleRoutes('admin')" 
				:class="{'bg-indigo-700 text-white scale-105': currentView === 'admin', 'border border-indigo-400 scale-100': currentView !== 'admin'}" 
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
import { ref, onMounted, computed } from 'vue'; // Import computed here
import axios from '../axios';

const info = ref({});
const currentView = ref('user'); // Default to user endpoints

const fetchApiInfo = async () => {
	try {
		const response = await axios.get('/'); // Adjust the URL based on your API endpoint
		info.value = response.data;

		// Initialize user_routes and admin_routes from the fetched data
		info.value.user_routes = response.data.routes.filter(route => !route.note); // Assuming user routes have no 'note'
		info.value.admin_routes = response.data.routes.filter(route => route.note); // Assuming admin routes have a 'note'
	} catch (error) {
		console.error('Error fetching documentation');
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

<style scoped>
.table-container {
	overflow-x: auto;
}
</style>
