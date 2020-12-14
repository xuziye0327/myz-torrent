import Vue from 'vue'
import App from './App.vue'
import Argon from '@/plugins/argon-kit'

Vue.config.productionTip = false
Vue.use(Argon)

new Vue({
  render: h => h(App),
}).$mount('#app')
