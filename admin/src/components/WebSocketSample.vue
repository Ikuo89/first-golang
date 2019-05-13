<template>
  <div>
    <h1>WebSocket Sample</h1>
    <b-form-input v-model="sendText"></b-form-input>
    <b-button @click="send">Send</b-button>
  </div>
</template>

<script>
export default {
  name: 'WebSocketSample',
  data () {
    return {
      sendText: '',
      sock: null
    }
  },
  methods: {
    send: function (event) {
      this.sock.send(this.sendText)
    }
  },
  mounted: function () {
    this.sock = new WebSocket('ws://localhost:8080/ws')

    this.sock.addEventListener('open', function (e) {
      console.log('Socket Connected')
    })

    this.sock.addEventListener('message', function (e) {
      console.log(e.data)
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
