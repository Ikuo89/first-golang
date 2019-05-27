<template>
  <div>
    <h1>User</h1>
    <div class="container">
      <div class="row">
        <div class="col-sm-4">ID</div><div class="col-sm-8">{{ user.id }}</div>
      </div>
      <div class="row">
        <div class="col-sm-4">Name</div><div class="col-sm-8"><b-form-input v-model="user.name"></b-form-input></div>
      </div>
      <div class="row">
        <div class="col-sm-4">Created At</div><div class="col-sm-8">{{ user.created_at }}</div>
      </div>
      <div class="row">
        <div class="col-sm-4">Updated At</div><div class="col-sm-8">{{ user.updated_at }}</div>
      </div>
      <div class="row">
        <div class="col-sm-4">WebSocket Page</div><div class="col-sm-8"><router-link :to="{ path: '/users/' + user.id + '/ws' }">ws</router-link></div>
      </div>
    </div>
    <b-button @click="updateUser">Update</b-button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'User',
  data () {
    return {
      sendText: '',
      userID: null,
      user: {
        name: ''
      }
    }
  },
  methods: {
    updateUser: function () {
      axios.put('http://localhost:8080/users/' + this.userID, this.user)
        .then(response => axios.get('http://localhost:8080/users/' + this.userID))
        .then(response => (this.user = response.data))
    }
  },
  mounted: function () {
    this.userID = this.$route.params.id
    axios.get('http://localhost:8080/users/' + this.userID)
      .then(response => (this.user = response.data))
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
