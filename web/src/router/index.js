import Vue from 'vue'
import VueRouter from 'vue-router'
import Points from '../components/Points'
import Ways from '../components/Ways'

Vue.use(VueRouter)

const routes = [
  {
    path: '/points',
    alias: '/',
    name: 'points',
    component: Points
  },
  {
    path: '/ways',
    name: 'ways',
    component: Ways
  }
]

const router = new VueRouter({
  routes
})

export default router
