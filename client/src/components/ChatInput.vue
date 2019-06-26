<template>
  <div class="field">
    <div class="control is-medium">
      <div class="columns is-mobile">
        <div class="column is-10">
          <input type="text" class="input is-medium" placeholder="e.g. Hello world"
            v-model="message" />
          <!-- <textarea class="textarea" placeholder="e.g. Hello world"
            v-on:keyup="handleUserTyping"
            v-if="isLongMessage"
            v-model="message"></textarea> -->
        </div>
        <div class="column send">
          <i class="material-icons">send</i>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'chat-input',
  data () {
    return {
      message: '',
      // isLongMessage: false
    }
  },
  methods: {
    // handleUserTyping (e) {
    //   console.log(e.target.value)
    //   var isTextArea = e.target.tagname === 'TEXTAREA' ? true : false
    //   this.isLongMessage = isTextArea
    //     ? e.target.scrollHeight > e.target.offsetHeight
    //     : e.target.scrollWidth > e.target.offsetWidth
    // }

  }
}
</script>

<style lang="scss" scoped>
  $buttonColor: #425983;
  i {
    font-size: 2.2rem;
    color: $buttonColor;
    cursor: pointer;
    &:hover {
      color: darken($buttonColor, 30%)
    }
  }

  .send {
    display: flex;
    align-items: center;
  }

</style>
