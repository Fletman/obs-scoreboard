import { createApp } from 'vue/dist/vue.esm-bundler.js';
import { createRouter, createWebHashHistory } from 'vue-router';
import ScoreGridPage from './pages/ScoreGridPage.vue'
import ScoreFeaturedPage from './pages/ScoreFeaturedPage';
import ScoreViewerPage from './pages/ScoreViewerPage.vue';
import ScoreEditorPage from './pages/ScoreEditorPage.vue';

const routes = [
    { path: '/', component: ScoreGridPage },
    { path: '/featured', component: ScoreFeaturedPage },
    { path: '/view/:id', component: ScoreViewerPage },
    { path: '/edit/:id', component: ScoreEditorPage }
];
const router = createRouter({
  history: createWebHashHistory(),
  routes
});

const app = createApp({});
app.use(router);
app.mount("#app");