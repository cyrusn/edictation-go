<template lang="html">
  <audio :src="audioSrc" autoplay ref='player'></audio>
</template>

<script>
export default {
  data () {
    return {
      audioSrc: ''
    }
  },
  methods: {
    play () {
      const vm = this
      vm.$refs.player.play()
    }
  },
  created () {
    const vm = this
    vm.$on('update:tts-audio', function ({name, index}) {
      console.log('update:tts-audio')
      vm.audioSrc = `./api/voice/assessment/${name}/index/${index}`
      vm.$refs.player.load()
    })
  }
}
</script>

<style lang="css">
</style>
