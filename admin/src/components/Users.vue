<template>
  <div>
    <h1>Users</h1>
    <ul>
      <li v-for="user in users" :key="user.ID">
        <router-link :to="{ path: 'users/' + user.ID }">
          {{ user.ID }} {{ user.name }}
        </router-link>
      </li>
    </ul>
    <b-button @click="createUser">New</b-button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Users',
  data () {
    return {
      users: []
    }
  },
  methods: {
    createUser: function () {
      axios.post('http://localhost:8080/users')
        .then(response => axios.get('http://localhost:8080/users'))
        .then(response => (this.users = response.data))
    }
  },
  mounted: function () {
    axios.get('http://localhost:8080/users')
      .then(response => (this.users = response.data))
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
