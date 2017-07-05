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
      level: 1,
      questionIDs: [],
      stat: {}
    }
  },
  computed: {
    ViewComponent () {
      return routes[this.currentRoute]
    },
    noOfQuestions () {
      const vm = this
      return vm.questionIDs.length
    }
  },
  render (h) {
    return h(this.ViewComponent)
  }
})
