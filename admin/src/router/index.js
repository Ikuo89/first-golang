import Vue from 'vue'
import Router from 'vue-router'
import User from '@/components/User'
import Users from '@/components/Users'
import UserWebSocket from '@/components/UserWebSocket'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/users/:id',
      name: 'User',
      component: User
    },
    {
      path: '/users/:id/ws',
      name: 'UserWebSocket',
      component: UserWebSocket
    },
    {
      path: '/',
      name: 'Users',
      component: Users
    }
  ]
})
