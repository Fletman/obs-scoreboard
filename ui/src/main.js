import { createApp } from 'vue/dist/vue.esm-bundler.js';
import { createRouter, createWebHashHistory } from 'vue-router';
import ScoreGridPage from './pages/ScoreGridPage.vue'
import ScoreFeaturedPage from './pages/ScoreFeaturedPage';
import ScoreViewerPage from './pages/ScoreViewerPage.vue';
import ScoreEditorPage from './pages/ScoreEditorPage.vue';
import BracketListPage from './pages/BracketListPage.vue';

const routes = [
    { path: '/scores', component: ScoreGridPage },
    { path: '/scores/featured', component: ScoreFeaturedPage },
    { path: '/scores/:id/view', component: ScoreViewerPage },
    { path: '/scores/:id/edit', component: ScoreEditorPage },
    { path: '/brackets', component: BracketListPage },
    { path: '/brackets/:id', component: {} }
];
const router = createRouter({
  history: createWebHashHistory(),
  routes
});

const app = createApp({});
app.use(router);
app.mount("#app");