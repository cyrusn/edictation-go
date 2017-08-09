import Vue from 'vue'
import Home from '../pages/Home.vue'
import Report from '../pages/Report.vue'
import Quiz from '../pages/Quiz.vue'

const routes = {
  '/': Home,
  '/report': Report,
  '/quiz': Quiz
}

module.exports = new Vue({
  data () {
    return {
      currentRoute: '/'
    }
  },
  computed: {
    ViewComponent () {
      return routes[this.currentRoute]
    }
  },
  mounted () {
    const vm = this
    vm.$on('update:route', function (newRoute) {
      console.log('update:currentRoute')
      vm.currentRoute = newRoute
    })
  }
})
