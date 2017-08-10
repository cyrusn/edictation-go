<template lang="html">
  <div class="panel panel-success">
    <div class="panel-heading">
      <h1 class="panel-title">{{vocab.definition}} {{vocab.partOfSpeech}}</h1>
    </div>

    <div class="panel-body">
      <div class="input-group" @keypress.enter="onNext">
        <input
          placeholder="Please type your answer!"
          id="answer" type="text" class="form-control" @input="onChange" v-model='answer'>

        <div class="input-group-btn">
          <button class="btn btn-default" @click.prevent="onNext">
            <span class="glyphicon glyphicon-thumbs-up" aria-hidden="true"></span>
          </button>

          <button class="btn btn-default" @click="speak">
            <span class="glyphicon glyphicon-volume-up" aria-hidden="true"></span>
          </button>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['vocab', 'speak', 'next'],
  data () {
    return {
      answer: ''
    }
  },
  watch: {
    vocab () {
      const vm = this
      vm.answer = ''
      document.getElementById('answer').focus()
    }
  },
  methods: {
    onChange () {
      const vm = this
      vm.$emit('updateAnswer', vm.answer)
    },
    onNext () {
      const vm = this
      vm.answer = ''
      document.getElementById('answer').focus()
      vm.next()
    }
  }
}
</script>

<style lang="css">
</style>
