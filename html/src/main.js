import Vue from 'vue'
import routes from './routes'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  data () {
    return {
      currentRoute: window.location.pathname,
      name: '',
      clazz: '',
      clazzNo: '',
      mode: 'normal',
      level: 1,
      noOfQuestions: 0,
      stat: {}
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
