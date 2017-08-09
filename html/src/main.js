import Vue from 'vue'
import routes from './routes'

Vue.filter('capitalize', str => str.charAt(0).toUpperCase() + str.slice(1))
Vue.filter('toPercentage', n => Math.round(n * 100))
/* eslint-disable no-new */
new Vue({
  el: '#app',
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
  render (h) {
    return h(this.ViewComponent)
  }
})
