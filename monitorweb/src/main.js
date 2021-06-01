import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import qs from 'qs'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// import router from '@/router'
import store from '@/store'
Vue.prototype.$axios = axios
Vue.prototype.$qs = qs
Vue.use(ElementUI)

if (process.env.NODE_ENV==='development'){
  // axios.defaults.baseURL = 'http://10.136.158.218:8089/'
  axios.defaults.baseURL = 'http://192.168.50.6:8089/'
}
console.log(process.env)

Vue.config.productionTip = false

new Vue({
  // router,
  store,
  render: h => h(App),
}).$mount('#app')
