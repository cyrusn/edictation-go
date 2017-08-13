<template lang="html">
  <div v-if="!hidden" class="alert alert-danger">
    <div v-for='message in validateMessages'>
      <strong>Warning!</strong> &nbsp{{message}}
    </div>
  </div>
</template>

<script>
import {mapGetters, mapState} from 'vuex'
export default {
  props: ['messages', 'show'],
  computed: {
    ...mapState(['firstSubmit']),
    ...mapGetters(['isValid', 'validateMessages']),
    hidden () {
      const {isValid, firstSubmit} = this
      return isValid || firstSubmit
    }
  }
}
</script>

<style lang="css">
</style>
