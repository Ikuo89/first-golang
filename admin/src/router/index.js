import Vue from 'vue'
import Router from 'vue-router'
import WebSocketSample from '@/components/WebSocketSample'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'WebSocketSample',
      component: WebSocketSample
    }
  ]
})
