import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
// 引入Element组件及其样式
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(ElementUI)

var axios = require('axios')
axios.defaults.baseURL = '/api'

Vue.config.productionTip = false

Vue.prototype.$axios = axios
new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')