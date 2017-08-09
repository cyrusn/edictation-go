import Vue from 'vue'

module.exports = new Vue({
  data () {
    return {
      currentRoute: '/'
    }
  },
  created () {
    const vm = this
    vm.$on('update:route', function (route) {
      vm.currentRoute = route
    })
  }
})
