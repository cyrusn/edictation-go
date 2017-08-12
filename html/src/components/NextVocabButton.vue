<template lang="html">
  <button
    class="btn btn-default"
    @click.prevent="onNext"
  >
    <span class="glyphicon glyphicon-thumbs-up" aria-hidden="true"></span>
  </button>
</template>

<script>
import {mapActions} from 'vuex'
export default {
  methods: {
    ...mapActions([
      'onNext'
    ])
  }
}
</script>

<style lang="css">
</style>
