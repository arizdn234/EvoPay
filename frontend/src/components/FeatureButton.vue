<template>
    <button
      @click="onClick"
      class="bg-blue-500 text-white font-sfpro font-bold py-2 px-4 rounded hover:bg-blue-700 hover:py-4 transition-all duration-300 ease-in-out"
    >
      {{ text }}
    </button>
  </template>
  
  <script setup>
  const props = defineProps({
    text: String,
    onClick: Function
  });
  </script>
  