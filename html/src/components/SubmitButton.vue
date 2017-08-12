<template lang="html">
  <div class="">
    {{validateMessages}}
      <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button class="btn btn-success" @click.prevent="onSubmit">Start</button>
      </div>
    </div>
    <warning-message :show='show' :messages='validateMessages'/>
  </div>
</template>

<script>
import WarningMessage from '../components/WarningMessage.vue'
import {mapGetters, mapMutations} from 'vuex'

export default {
  components: {
    WarningMessage
  },
  data () {
    return {
      messages: [],
      firstRun: true
    }
  },
  computed: {
    ...mapGetters(['validateMessages']),
    show () {
      return this.validateMessages.length !== 0 && !this.firstRun
    }
  },
  methods: {
    ...mapMutations('router', ['goto']),
    onSubmit: function (event) {
      // TODO: only user and assessment on form submit,
      // and seperate validation from store, validate values before submit.
      const {goto, validateMessages} = this
      const isValid = validateMessages.length === 0
      this.firstRun = false
      if (isValid) {
        goto('/quiz')
      }
    }
  }
}
</script>

<style lang="css">
</style>
