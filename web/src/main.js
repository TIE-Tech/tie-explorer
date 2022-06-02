import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import NProgress from 'nprogress'
NProgress.configure({
    easing: 'ease',  // 动画方式
    speed: 500,
    showSpinner: true,
    trickleSpeed: 200,
    minimum: 0.5
})

router.beforeEach((to, from , next) => {
    NProgress.start();
    next();
});

router.afterEach(() => {
    NProgress.done()
})

import '@/assets/css/tabler.min.css'
import 'nprogress/nprogress.css'

createApp(App).use(router).mount('#app')
