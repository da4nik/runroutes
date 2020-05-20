import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import YmapPlugin from 'vue-yandex-maps'

const ymapSettings = {
  apiKey: '',
  lang: 'ru_RU',
  coordorder: 'latlong',
  version: '2.1',
  coords: [52.638634, 41.443438]
}

Vue.config.productionTip = false
Vue.use(YmapPlugin, ymapSettings)

Vue.prototype.$eventHub = new Vue()

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
