<template id="hello-world">
<main-layout>
  <h1 class="text-center">
    S.K.H. Li Ping Secondary School<br><small>eDictation System</small>
  </h1>
  <hr>

  <div v-if="!isValid" class="alert alert-danger" role="alert">Invalid user information</div>
  <form class="form-horizontal" @keypress.enter="submit">
    <div class="form-group">
      <label class="col-sm-2 control-label">Name</label>
      <div class="col-sm-10">
        <input id="name" v-model="$root.name" class="form-control">
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">Class</label>
      <div class=" col-sm-10">
        <select class="form-control" v-model="$root.clazz">
          <option v-for="clazz in ['1A','1B','1C','1D','2A','2B','2C','2D','3A','3B','3C','3D','4A','4B','4C','4D','4E','5A','5B','5C','5D','5E','6A','6B','6C','6D','6E']">
            {{clazz}}
          </option>
        </select>
      </div>
    </div>

    <div class="form-group">
      <label class="col-sm-2 control-label">Class No.</label>
      <div class="col-sm-10">
        <input id="name" v-model="$root.clazzNo" class="form-control">
      </div>
    </div>

    <div class="form-group">
      <label class="col-sm-2 control-label">Level</label>
      <div class="col-sm-10">
        <label class="radio-inline" v-for="l in 6">
          <input type="radio" v-model="$root.level" :value="l" /> {{ l }}
        </label>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">Mode</label>
      <div class="col-sm-10">
        <label class="radio-inline" v-for="mo in ['easy', 'normal', 'hard']">
          <input type="radio" v-model="$root.mode" :value="mo" /> {{ mo | capitalize }}
        </label>
      </div>
    </div>
    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button class="btn btn-success" @click.prevent="submit">Next</button>
      </div>
    </div>
  </form>
</main-layout>
</template>

<script type="text/javascript">
import axios from 'axios'
import MainLayout from '../layouts/Main.vue'
import _ from 'lodash'

// TODO: Validation on student, class and classNo, empty entry is not allowed.
export default {
  data () {
    return {
      isValid: true
    }
  },
  mounted () {
    document.getElementById('name').focus()
  },
  methods: {
    submit: function (event) {
      const vm = this
      const root = vm.$root

      // simple Validation: all fields can't be empty
      vm.isValid = !(root.name === '' || root.clazz === '' || root.clazzNo === '')

      if (vm.isValid) {
        vm.$root.currentRoute = '/quiz'
        axios.get('./api/level/' + vm.$root.level)
        .then(function (response) {
          vm.$root.questionIDs = _.shuffle(response.data)
        })
      }
    }
  },
  filters: {
    capitalize: function (str) {
      return str.charAt(0).toUpperCase() + str.slice(1)
    }
  },
  components: {
    MainLayout
  }
}
</script>
