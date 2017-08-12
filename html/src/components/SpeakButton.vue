<template lang="html">
  <button class="btn btn-default" onclick="this.blur()">
    <audio :src="audioSource" autoplay ref='player'></audio>
    <span
      class="glyphicon glyphicon-volume-up"
      aria-hidden="true"
      @click='play'></span>
  </button>
</template>

<script>
import {mapState} from 'vuex'
export default {
  computed: {
    ...mapState('vocab', [
      'audioSource'
    ])
  },
  methods: {
    play () {
      this.$refs.player.play()
    }
  }
}
</script>

<style lang="css">
</style>
