<template lang="html">
  <div class="">
    <audio :src="src" autoplay ref='player'></audio>
  </div>
</template>

<script>
// import {mapState} from 'vuex'

export default {
  data () {
    return {
      src: ''
    }
  },
  methods: {
    play () {
      this.src = this.$store.state.assessment.vocab.src
      this.$refs.player.load()
    }
  }
}
</script>

<style lang="css">
</style>
