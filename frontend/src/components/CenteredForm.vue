<template>
    <div class="flex items-center justify-center min-h-screen bg-background-soft">
        <div class="bg-background rounded-lg shadow-lg p-8 max-w-md w-full">
            <h1 class="text-2xl font-semibold text-heading text-center mb-6">{{ title }}</h1>
            <p class="text-center text-text-secondary-light">{{ description }}</p>

            <!-- Form Placeholder -->
            <div class="space-y-4">
                <!-- Placeholder for form inputs -->
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'CenteredForm',
    props: {
        title: {
            type: String,
            default: 'Form Title',
        },
        description: {
            type: String,
            default: 'Form Description',
        }
    },
};
</script>

<style scoped></style>