import Vue from 'vue'
import Router from './mixins/Router'
import Home from './pages/Home.vue'
import Report from './pages/Report.vue'
import Quiz from './pages/Quiz.vue'

const routes = {
  '/': Home,
  '/report': Report,
  '/quiz': Quiz
}

Vue.filter('capitalize', str => str.charAt(0).toUpperCase() + str.slice(1))
Vue.filter('toPercentage', n => Math.round(n * 100))
/* eslint-disable no-new */
new Vue({
  el: '#app',
  mixins: [Router],
  render (h) {
    return h(routes[Router.currentRoute])
  }
})
