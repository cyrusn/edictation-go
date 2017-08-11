import Vue from 'vue'
import Home from './pages/Home.vue'
import Report from './pages/Report.vue'
import Quiz from './pages/Quiz.vue'
import _ from 'lodash'

import store from './store'

const routes = {
  '/': Home,
  '/report': Report,
  '/quiz': Quiz
}

Vue.filter('capitalize', str => str.charAt(0).toUpperCase() + str.slice(1))
Vue.filter('toPercentage', n => Math.round(n * 100))
Vue.filter('sortByIndex', array => _.orderBy(array, 'index'))

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  render (h) {
    return h(routes[this.$store.state.currentRoute])
  }
})
