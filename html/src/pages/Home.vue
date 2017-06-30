<template id="hello-world">
<main-layout>
  <form>
    <table class="table">
      <tr>
        <td>Name</td>
        <td><input id="name" v-model="name"></td>
      </tr>
      <tr>
        <td>Levels</td>
        <td>
          <label class="radio-inline" v-for="l in [1, 2, 3, 4, 5, 6]">
            <input type="radio" v-model="level" :value="l" /> {{ l }} &nbsp;
          </label>
        </td>
      </tr>
      <tr>
        <td>Mode</td>
        <td>
          <label class="radio-inline" v-for="mo in ['easy', 'normal', 'hard']">
            <input type="radio" v-model="mode" :value="mo" /> {{ mo | capitalize }} &nbsp;
          </label>
        </td>
      </tr>
    </table>
    <button class="btn btn-success" @click="onSubmit">Submit</button>
  </form>
  <p>The test is about to begin, please selected the mode and level, then press submit.</p>
  <p>{{ message }}</p>
</main-layout>
</template>

<script type="text/javascript">
import axios from "axios"
import routes from "../routes"
import MainLayout from '../layouts/Main.vue'

module.exports = {
  data: function() {
    const vm = this
    return {
      name: vm.$root.name,
      mode: vm.$root.mode,
      level: vm.$root.level
    }
  },
  mounted() {
    const vm = this
    document.getElementById('name').focus()
    window.addEventListener('keydown', e => {
      if (e.keyCode === 19){
        vm.onSubmit(e)
      }
    })
  },
  methods: {
    onSubmit: function(event) {
      const vm = this
      vm.$root.name = vm.name
      vm.$root.mode = vm.mode
      vm.$root.level = vm.level

      event.preventDefault()
      vm.$root.currentRoute = "/quiz"
      window.history.pushState(
        null,
        routes["/quiz"],
        "/quiz"
      )
      axios.get('/api/'+vm.level)
      .then(function(response) {
        vm.$root.noOfQuestions = response.data.noOfQuestions
      })
    }
  },
  filters: {
    capitalize: function(str) {
      return str.charAt(0).toUpperCase() + str.slice(1)
    }
  },
  components: {
    MainLayout
  }
};
</script>
