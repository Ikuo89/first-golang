<template>
  <div>
    <h1>WebSocket User</h1>
    <div class="container">
      <div class="row">
        <b-form-input class="col-sm-10" v-model="sendText"></b-form-input>
        <b-button class="col-sm-2" @click="send">Send</b-button>
      </div>
      <div class="row">
        <b-form-textarea readonly v-model="socketResponse" rows="20"></b-form-textarea>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UserWebSocket',
  data () {
    return {
      sendText: '',
      socketResponse: '',
      sock: null,
      userID: null
    }
  },
  methods: {
    send: function (event) {
      this.sock.send(this.sendText)
      this.sendText = ''
    }
  },
  mounted: function () {
    this.userID = this.$route.params.id
    this.sock = new WebSocket('ws://localhost:8080/users/' + this.userID + '/ws')

    this.sock.addEventListener('open', (e) => {
      console.log('Socket Connected: ' + this.userID)
    })

    this.sock.addEventListener('message', (e) => {
      console.log(e.data)
      this.socketResponse += e.data + '\n'
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
