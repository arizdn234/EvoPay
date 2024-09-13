<template>
	<div class="documentation p-5 max-w-4xl mx-auto">
		<h1 class="text-2xl font-bold mb-5">{{ info.service }} API Documentation</h1>
		<p class="text-lg mb-5">Version: {{ info.version }}</p>
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
					<tr v-for="(route, index) in info.routes" :key="index"
						class="even:dark:bg-gray-800 odd:dark:bg-gray-700 even:bg-gray-400 hover:bg-gray-700 hover:text-white transition-all duration-300">
						<td class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.method }}</td>
						<td class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.path }}</td>
						<td class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.description }}</td>
						<td v-if="route.note" class="border border-gray-400 px-4 py-2 whitespace-nowrap">{{ route.note
							}}</td>
						<td v-else class="border border-gray-400 px-4 py-2 whitespace-nowrap">-</td>
					</tr>
				</tbody>
			</table>
		</div>
		<p class="mt-5 text-sm italic">*{{ info.note }}</p>
	</div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from '../axios';

const info = ref({});

const fetchApiInfo = async () => {
	try {
		const response = await axios.get('/');

		info.value = response.data;
	} catch (error) {
		console.error('Error fetching documentation:', error);
	}
};
onMounted(fetchApiInfo);
</script>