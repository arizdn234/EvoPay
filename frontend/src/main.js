import './assets/main.css'

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import { diffForHuman, formatDate } from './dateUtils';

const app = createApp(App);

app.mixin({
    methods: {
        diffForHuman,
        formatDate
    }
});

app.use(router);

app.mount('#app');
