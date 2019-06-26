<template>
  <div class="container">
    <ChatOutput class="chat-output" />
    <ChatInput class="chat-input" />
  </div>
</template>

<script>
import ChatInput from './ChatInput'
import ChatOutput from './ChatOutput'

export default {
  name: 'chat',
  components: {
    ChatInput,
    ChatOutput
  }
}
</script>

<style lang="scss" scoped>
  .container {
    padding: 0 30px;
  }
  .chat-input {
    margin-top: 40px;
  }
</style>
