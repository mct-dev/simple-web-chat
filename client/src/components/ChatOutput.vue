<template>
  <div class="container">
    <!-- for each message render component -->
    <ChatItem v-bind:key="index" v-for="(msg, index) in messages" v-bind:msg="msg.content" v-bind:user="msg.user" />
  </div>
</template>

<script>
import ChatItem from './ChatItem'

export default {
  name: 'chat-output',
  components: {
    ChatItem
  },
  data () {
    return {
      messages: [
        { content: 'wow here is message 1', user: 'user1' },
        { content: 'wow here is message 2', user: 'user2' },
        { content: 'wow here is message 3', user: 'user3' },
        { content: 'wow here is message 4', user: 'user4' },
        { content: 'wow here is message 5', user: 'user5' },
        { content: 'wow here is message 6', user: 'user6' },
        { content: 'wow here is message 7', user: 'user7' },
        { content: 'wow here is message 8', user: 'user8' },
        { content: 'wow here is message 9', user: 'user9' },
        { content: 'wow here is message 10', user: 'user10' }
      ]
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
