import Vue from 'vue'
import Router from './mixins/router'

Vue.filter('capitalize', str => str.charAt(0).toUpperCase() + str.slice(1))
Vue.filter('toPercentage', n => Math.round(n * 100))
/* eslint-disable no-new */
new Vue({
  el: '#app',
  mixins: [Router],
  render (h) {
    return h(Router.ViewComponent)
  }
})
