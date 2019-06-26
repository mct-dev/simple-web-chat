<template>
  <div>
    <ChatNav />
    <div class="container">
      <Chat />
    </div>
  </div>
</template>

<script>
import Chat from './components/Chat'
import ChatNav from './components/ChatNav'

export default {
  name: 'app',
  components: {
    Chat,
    ChatNav
  },
  data () {
    return {
      webSocket: null,
      newMsg: '', // a new message for the server
      chatMessages: '', // list of chat messages to display

      // for when we have a login system
      email: null,
      username: null,
      joined: false
    }
  }
}
</script>

<style lang="scss">
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #3d3d3d;
}
</style>
