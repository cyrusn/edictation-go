<template lang="html">
  <input
    id='answer'
    placeholder="Please type your answer!"
    type="text"
    class="form-control"
    @input="onInput"
    :value="answer"
    @keypress.enter.prevernt='onNext()'
    />
</template>

<script>
import {mapState, mapMutations, mapActions} from 'vuex'
export default {
  mounted () {
    document.getElementById('answer').focus()
  },
  computed: {
    ...mapState('vocab', ['answer'])
  },
  methods: {
    ...mapMutations('vocab', [
      'updateAnswer'
    ]),
    ...mapActions(['onNext']),
    onInput (e) {
      const {updateAnswer} = this
      updateAnswer(e.target.value)
    }
  }
}
</script>

<style lang="css">
</style>
