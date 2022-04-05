import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(), // hash模式：createWebHashHistory，history模式：createWebHistory
    routes: [{
            path: '/',
            redirect: '/product'
        },
        {
            path: '/login',
            name: 'login',
            component: () =>
                import ('./components/Login.vue')
        },
        {
            path: '/account',
            name: 'account',
            component: () =>
                import ('./components/Account.vue')
        },
        {
            path: '/addGood',
            name: 'addGood',
            component: () =>
                import ('./components/AddGood.vue')
        },
        {
            path: '/category',
            name: 'category',
            component: () =>
                import ('./components/Category.vue')
        },
        {
            path: '/addCategory',
            name: 'addCategory',
            component: () =>
                import ('./components/DialogAddCategory.vue')
        },
        {
            path: '/addSwiper',
            name: 'addSwiper',
            component: () =>
                import ('./components/DialogAddSwiper.vue')
        },
        {
            path: '/guest',
            name: 'guest',
            component: () =>
                import ('./components/Guest.vue')
        },
        {
            path: '/index',
            name: 'index',
            component: () =>
                import ('./components/Index.vue')
        },
        {
            path: '/indexConfig',
            name: 'indexConfig',
            component: () =>
                import ('./components/IndexConfig.vue')
        },
        {
            path: '/order',
            name: 'order',
            component: () =>
                import ('./components/Order.vue')
        },
        {
            path: '/orderDetail',
            name: 'orderDetail',
            component: () =>
                import ('./components/OrderDetail.vue')
        },
        {
            path: '/swiper',
            name: 'swiper',
            component: () =>
                import ('./components/Swiper.vue')
        }
    ]
})

export default router