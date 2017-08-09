<template lang="html">
  <div class="assessment">
    <div class="form-group">
      <label class="col-sm-2 control-label">Assessment</label>
      <div class=" col-sm-10">
        <select
          class="form-control"
          v-model="name"
          @change='updateName'
        >
          <option value="" disabled selected>Select the test</option>
          <option v-for="name in names">
            {{name}}
          </option>
        </select>
      </div>
    </div>

    <div class="form-group">
      <label class="col-sm-2 control-label">Mode</label>
      <div class="col-sm-10">
        <label class="radio-inline" v-for="m in ['easy', 'normal', 'hard']">
          <input
            type="radio"
            @change='updateMode'
            v-model='mode'
            :value="m"
          /> {{ m | capitalize }}
        </label>
      </div>
    </div>
  </div>
</template>

<script>
import Assessment from '../mixins/Assessment'
import axios from 'axios'
import _ from 'lodash'

export default {
  mixins: [Assessment],
  data () {
    return {
      mode: 'normal',
      name: '',
      names: []
    }
  },
  created () {
    const vm = this

    axios.get('./api/assessment')
      .then(response => {
        vm.names = _.orderBy(response.data)
      })
      .catch(err => console.log(err))
  },
  methods: {
    // TODO: Validate if user didn't select the assessment name
    updateMode () {
      Assessment.$emit('update:assessment-mode', this.mode)
    },
    updateName () {
      Assessment.$emit('update:assessment-name', this.name)
    }
  }
}
</script>

<style lang="css">
option[value=""][disabled] {
  display: none;
}
</style>
