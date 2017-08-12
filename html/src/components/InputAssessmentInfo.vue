<template lang="html">
  <div>
    <div class="form-group">
      <label class="col-sm-2 control-label">Assessment</label>
      <div class=" col-sm-10">
        <select
          id='name'
          class="form-control"
          @change='onUpdateName'
        >
          <option disabled selected>Select the test</option>
          <option v-for="n in names">
            {{n}}
          </option>
        </select>
      </div>
    </div>

    <div class="form-group">
      <label class="col-sm-2 control-label">Mode</label>
      <div class="col-sm-10">
        <label class="radio-inline" v-for="m in ['easy', 'normal', 'hard']">
          <input
            name="modeOption"
            type="radio"
            @change='onUpdateMode'
            :value='m'
          />
          {{ m | capitalize }}
        </label>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import _ from 'lodash'

import {mapState, mapMutations} from 'vuex'

export default {
  data () {
    return {
      names: []
    }
  },
  computed: {
    ...mapState('assessment', [
      'name',
      'mode'
    ])
  },
  created () {
    const vm = this
    axios.get('./api/assessment')
      .then(response => {
        vm.names = _.orderBy(response.data)
      })
      .catch(console.error)
  },
  methods: {
    ...mapMutations('assessment', [
      'updateName',
      'updateMode'
    ]),
    onUpdateName (e) {
      this.updateName(e.target.value)
    },
    onUpdateMode (e) {
      this.updateMode(e.target.value)
    }
  }
}
</script>

<style lang="css">
option[value=""][disabled] {
  display: none;
}
</style>
