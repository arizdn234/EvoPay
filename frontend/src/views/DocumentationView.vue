<template>
	<div class="documentation">
		<h1>{{ info.service }} API Documentation</h1>
		<p>Version: {{ info.version }}</p>
		<div class="table-container">
			<table>
				<thead>
					<tr>
						<th>Method</th>
						<th>Path</th>
						<th>Description</th>
						<th>Note</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(route, index) in info.routes" :key="index">
						<td>{{ route.method }}</td>
						<td>{{ route.path }}</td>
						<td>{{ route.description }}</td>
						<td v-if="route.note">{{ route.note }}</td>
						<td v-else>-</td>
					</tr>
				</tbody>
			</table>
		</div>
		<p>*{{ info.note }}</p>
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

<style scoped>
.documentation {
	padding: 20px;
	max-width: 1000px;
	margin: 0 auto;
}

.documentation h1 {
	font-size: 2em;
	margin-bottom: 20px;
}

.table-container {
	overflow-x: auto;
	/* Enable horizontal scrolling */
	width: 100%;
	/* Ensure container takes full width */
}

table {
	width: 100%;
	border-collapse: collapse;
	margin-top: 20px;
}

th,
td {
	border: 1px solid #dddddd;
	padding: 8px;
	text-align: left;
}

th {
	background-color: #1a282e;
	font-weight: bold;
	color: #ffffff;
	/* Text color for better contrast */
}

tr {
	transition: all .5s;
}

tr:nth-child(even) {
	background-color: #303030;
}

tr:hover {
	background-color: #444444;
}

td {
	white-space: nowrap;
	/* Prevent text from wrapping */
}

@media (max-width: 600px) {
	.documentation h1 {
		font-size: 1.5em;
	}
}
</style>