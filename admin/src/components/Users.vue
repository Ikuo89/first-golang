<template>
  <div>
    <h1>Users</h1>
    <b-table striped hover :items="users" :fields="fields">
      <template slot="actions" slot-scope="row">
        <router-link :to="{ path: 'users/' + row.item.id }">
          detail
        </router-link>
        <b-button @click="deleteUser(row.item.id)">Delete</b-button>
      </template>
    </b-table>
    <div class="container">
      <div class="row">
        <b-form-input v-model="user.name" class="col-sm-8"></b-form-input>
        <b-button @click="createUser" class="col-sm-4">New</b-button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Users',
  data () {
    return {
      users: [],
      user: {},
      fields: [
        { key: 'id', lable: 'ID' },
        { key: 'name', lable: 'Name' },
        { key: 'created_at', lable: 'Created At' },
        { key: 'updated_at', lable: 'Updated At' },
        { key: 'actions', lable: 'Actions' }
      ]
    }
  },
  methods: {
    createUser: function () {
      axios.post('http://localhost:8080/users', this.user)
        .then(response => axios.get('http://localhost:8080/users'))
        .then(response => (this.users = response.data))
    },
    deleteUser: function (id) {
      axios.delete('http://localhost:8080/users/' + id)
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
