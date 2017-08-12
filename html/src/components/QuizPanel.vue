<template lang="html">
  <div class="panel panel-success">
    <div class="panel-heading">
      <h1 class="panel-title">Q{{runningIndex + 1}}. {{definition}} {{partOfSpeech}}</h1>
    </div>

    <div class="panel-body">
      <div class="input-group">
        <input-answer />
        <div class="input-group-btn">
          <next-vocab-button />
          <speak-button />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NextVocabButton from './NextVocabButton.vue'
import SpeakButton from './SpeakButton.vue'
import InputAnswer from './InputAnswer.vue'
import {mapState, mapActions} from 'vuex'

export default {
  components: {
    NextVocabButton,
    SpeakButton,
    InputAnswer
  },
  computed: {
    ...mapState('vocab', [
      'runningIndex',
      'definition',
      'partOfSpeech'
    ])
  },
  methods: {
    ...mapActions([
      'onStart'
    ])
  },
  mounted () {
    this.onStart()
  }
}
</script>

<style lang="css">
</style>
