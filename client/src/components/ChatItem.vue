<template>
  <div class="chat-item">
    <p>{{ user }}</p>
    <p>{{ msg }}</p>
  </div>
</template>

<script>
export default {
  name: 'chat-item',
  props: {
    msg: String,
    user: String
  }
}
</script>
