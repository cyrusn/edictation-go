<template lang="html">
  <div class="">
      <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button class="btn btn-success" @click.prevent="onSubmit">Start</button>
      </div>
    </div>
    <warning-message :show='show' :messages='messages'/>
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
      show: false
    }
  },
  computed: {
    ...mapGetters(['isValid'])
  },
  methods: {
    ...mapMutations('router', ['goto']),
    onSubmit: function (event) {
      // TODO: only user and assessment on form submit,
      // and seperate validation from store, validate values before submit.
      const {goto, isValid} = this
      if (isValid === null) {
        return goto('/quiz')
      } else {
        this.show = isValid !== null
        this.messages = isValid.details.map(detail => detail.message)
      }
    }
  }
}
</script>

<style lang="css">
</style>
