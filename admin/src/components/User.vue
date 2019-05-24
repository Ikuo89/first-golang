<template>
  <div>
    <h1>User</h1>
    <b-form-input v-model="user.name"></b-form-input>
    <b-button @click="updateUser">Update</b-button>
    <router-link :to="{ path: 'users/' + user.ID + '/ws' }">ws</router-link>
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
