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
                import ('./views/Login.vue')
        },
        {
            path: '/account',
            name: 'account',
            component: () =>
                import ('./views/Account.vue')
        },
        {
            path: '/addGood',
            name: 'addGood',
            component: () =>
                import ('./views/AddGood.vue')
        },
        {
            path: '/category',
            name: 'category',
            component: () =>
                import ('./views/Category.vue')
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
                import ('./views/Guest.vue')
        },
        {
            path: '/index',
            name: 'index',
            component: () =>
                import ('./views/Index.vue')
        },
        {
            path: '/indexConfig',
            name: 'indexConfig',
            component: () =>
                import ('./views/IndexConfig.vue')
        },
        {
            path: '/order',
            name: 'order',
            component: () =>
                import ('./views/Order.vue')
        },
        {
            path: '/orderDetail',
            name: 'orderDetail',
            component: () =>
                import ('./views/OrderDetail.vue')
        },
        {
            path: '/goods',
            name: 'goods',
            component: () =>
                import ('./views/goods/index.vue')
        }
    ]
})

export default router