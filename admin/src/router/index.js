import Vue from 'vue'
import Router from 'vue-router'
import WebSocketSample from '@/components/WebSocketSample'
import UserWebSocket from '@/components/UserWebSocket'
import Users from '@/components/Users'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/users/:id',
      name: 'UserWebSocket',
      component: UserWebSocket
    },
    {
      path: '/',
      name: 'WebSocketSample',
      component: WebSocketSample
    },
    {
      path: '/users',
      name: 'Users',
      component: Users
    }
  ]
})
