import Vue from 'vue'
import routes from './routes'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  data () {
    return {
      currentRoute: '/',
      name: '',
      clazz: '',
      clazzNo: '',
      mode: 'normal',
      correctPercentage: 0,
      assessmentName: '',
      assessmentRecords: [],
      assessmentSize: 0
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
