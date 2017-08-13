<template lang="html">
  <div class="">
      <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button class="btn btn-success" @click.prevent="onSubmit">Start</button>
      </div>
    </div>
    <warning-message />
  </div>
</template>

<script>
import WarningMessage from '../components/WarningMessage.vue'
import {mapGetters, mapMutations} from 'vuex'

export default {
  components: {
    WarningMessage
  },
  computed: {
    ...mapGetters(['isValid'])
  },
  methods: {
    ...mapMutations(['setFirstSubmitToFalse']),
    ...mapMutations('router', ['goto']),
    onSubmit: function (event) {
      const {goto, isValid, setFirstSubmitToFalse} = this
      setFirstSubmitToFalse()
      if (isValid) {
        goto('/quiz')
      }
    }
  }
}
</script>

<style lang="css">
</style>
